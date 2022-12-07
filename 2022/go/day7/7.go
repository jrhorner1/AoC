package day7

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	sizeLimit  = 100000
	fsLimit    = 70000000
	freeNeeded = 30000000
)

type Directory struct {
	size           int
	name           string
	parent         *Directory
	subDirectories map[string]*Directory
	files          map[string]*File
}

type File struct {
	size int
	name string
}

func Puzzle(input *[]byte, part2 bool) int {
	//logrus.SetLevel(logrus.DebugLevel)
	var fs *Directory
	var currentDir *Directory
	for i, command := range strings.Split(strings.TrimSpace(string(*input)), "\n$ ") {
		if i == 0 {
			fs = newDirectory("/", nil)
			command = command[2:] // trim leading $
		}
		cd_re := regexp.MustCompile(`^cd`)
		ls_re := regexp.MustCompile(`^ls`)
		if cd_re.Match([]byte(command)) {
			currentDir = cd(command, currentDir, fs)
		} else if ls_re.Match([]byte(command)) {
			ls(command, currentDir)
			continue
		}
	}
	// return to root so that sizes get populated properly
	for currentDir.parent != nil {
		currentDir = cd("cd ..", currentDir, fs)
	}
	logrus.Debug(fs.name, " ", fs.size)
	if part2 {
		freeSpace := fsLimit - fs.size
		targetSize := freeNeeded - freeSpace
		return smallestDir(fs, targetSize)
	}
	return sumUnderLimit(fs)
}

func cd(line string, dir *Directory, root *Directory) *Directory {
	var argument string
	_, err := fmt.Sscanf(line, "cd %s", &argument)
	if err != nil {
		logrus.Error(err, ": ", line)
	}
	switch argument {
	case "/":
		return root
	case "..":
		logrus.Debugf("Changing directory: %s to %s", dir.name, dir.parent.name)
		dir.parent.size += dir.size
		return dir.parent
	default:
		logrus.Debugf("Changing directory: %s to %s", dir.name, dir.subDirectories[argument].name)
		return dir.subDirectories[argument]
	}
}

func ls(output string, dir *Directory) {
	for i, line := range strings.Split(strings.TrimSpace(output), "\n") {
		logrus.Debugf("Line: '%s'", line)
		if i == 0 {
			continue // skip command line
		}
		if strings.Contains(line, "dir") {
			var name string
			n, err := fmt.Sscanf(line, "dir %s", &name)
			if err != nil {
				logrus.Error(err, ": ", line)
			}
			logrus.Debugf("Scanned %d", n)
			dir.subDirectories[name] = newDirectory(name, dir)
			logrus.Debugf("Added directory: %s %d %s", dir.name, dir.size, name)
		} else {
			file := &File{}
			n, err := fmt.Sscanf(line, "%d %s", &file.size, &file.name)
			logrus.Debugf("Scanned %d", n)
			if err != nil {
				logrus.Error(err, ": ", line)
			}
			dir.files[file.name] = file
			dir.size += file.size
			logrus.Debugf("Added file: %s %d %s %d", dir.name, dir.size, file.name, file.size)
		}
	}
}

func newDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		size:           0,
		name:           name,
		parent:         parent,
		subDirectories: make(map[string]*Directory),
		files:          make(map[string]*File),
	}
}

func sumUnderLimit(dir *Directory) int {
	sum := 0
	if dir.size < sizeLimit {
		logrus.Debugf("Adding %d from %s", dir.size, dir.name)
		sum += dir.size
	}
	for _, subDir := range dir.subDirectories {
		logrus.Debug(subDir.name, " ", subDir.size)
		if len(subDir.subDirectories) > 0 {
			sum += sumUnderLimit(subDir)
		} else if subDir.size < sizeLimit {
			sum += subDir.size
		}
	}
	return sum
}

func smallestDir(dir *Directory, targetSize int) int {
	size := dir.size
	if dir.size > targetSize {
		for _, subDir := range dir.subDirectories {
			if subDir.size > targetSize && subDir.size < size {
				size = subDir.size
			}
			if len(subDir.subDirectories) > 0 {
				subSize := smallestDir(subDir, targetSize)
				if subSize > targetSize && subSize < size {
					size = subSize
				}
			}
		}
	}
	return size
}
