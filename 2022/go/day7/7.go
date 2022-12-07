package day7

import (
	"fmt"
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
	files          map[string]int
	pwd            *Directory
}

func Puzzle(input *[]byte, part2 bool) int {
	var fs *Directory
	for i, command := range strings.Split(strings.TrimSpace(string(*input)), "\n$ ") {
		if i == 0 {
			fs = newDirectory("/", nil)
			fs.pwd = fs
			command = command[2:] // trim leading $ and space
		}
		if command[:2] == "cd" {
			fs.pwd = fs.cd(command[3:])
		} else if command[:2] == "ls" {
			fs.ls(strings.Split(strings.TrimSpace(command), "\n")[1:])
			continue
		}
	}
	// traverse back up to root so that directory sizes get populated properly
	for fs.pwd.parent != nil {
		fs.pwd = fs.cd("..")
	}
	if part2 {
		freeSpace := fsLimit - fs.size
		targetSize := freeNeeded - freeSpace
		return smallestDir(fs, targetSize)
	}
	return sumUnderLimit(fs)
}

func (fs *Directory) cd(argument string) *Directory {
	switch argument {
	case "/":
		return fs
	case "..":
		fs.pwd.parent.size += fs.pwd.size
		return fs.pwd.parent
	default:
		return fs.pwd.subDirectories[argument]
	}
}

func (fs *Directory) ls(output []string) {
	for _, line := range output {
		var name string
		if line[:3] == "dir" {
			name = line[4:]
			fs.pwd.subDirectories[name] = newDirectory(name, fs.pwd)
		} else {
			var size int
			_, err := fmt.Sscanf(line, "%d %s", &size, &name)
			if err != nil {
				logrus.Error(err, ": ", line)
			}
			fs.pwd.files[name] = size
			fs.pwd.size += size
		}
	}
}

func newDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		size:           0,
		name:           name,
		parent:         parent,
		subDirectories: make(map[string]*Directory),
		files:          make(map[string]int),
		pwd:            nil,
	}
}

func sumUnderLimit(dir *Directory) int {
	sum := 0
	if dir.size < sizeLimit {
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
