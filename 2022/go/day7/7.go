package day7

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

func Puzzle(input *[]byte, part2 bool) int {
	fs := newDirectory("/", nil)
	for i, command := range strings.Split(strings.TrimSpace(string(*input)), "\n$ ") {
		if i == 0 {
			command = command[2:] // trim leading $ and space
		}
		switch command[:2] {
		case "cd":
			fs.pwd = fs.cd(command[3:])
		case "ls":
			fs.ls(strings.Split(strings.TrimSpace(command), "\n")[1:])
			continue
		default:
			logrus.Panic("404 Command not found.")
		}
	}
	// traverse back up to root so that directory sizes get populated properly
	for fs.pwd.parent != nil {
		fs.pwd = fs.cd("..")
	}
	if part2 {
		const fsLimit = 70000000
		const freeNeeded = 30000000
		freeSpace := fsLimit - fs.size
		targetSize := freeNeeded - freeSpace
		return fs.smallestSubDirSize(targetSize)
	}
	const sizeLimit = 100000
	return fs.sumUnderLimit(sizeLimit)
}

type Directory struct {
	size           int
	name           string
	parent         *Directory
	subDirectories map[string]*Directory
	files          map[string]int
	pwd            *Directory
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

func (d *Directory) cd(argument string) *Directory {
	switch argument {
	case "/":
		return d
	case "..":
		d.pwd.parent.size += d.pwd.size
		return d.pwd.parent
	default:
		return d.pwd.subDirectories[argument]
	}
}

func (d *Directory) ls(output []string) {
	for _, line := range output {
		var name string
		if line[:3] == "dir" {
			name = line[4:]
			d.pwd.subDirectories[name] = newDirectory(name, d.pwd)
		} else {
			var size int
			_, err := fmt.Sscanf(line, "%d %s", &size, &name)
			if err != nil {
				logrus.Error(err, ": ", line)
			}
			d.pwd.files[name] = size
			d.pwd.size += size
		}
	}
}

func (d *Directory) sumUnderLimit(limit int) int {
	sum := 0
	if d.size < limit {
		sum += d.size
	}
	for _, subDir := range d.subDirectories {
		logrus.Debug(subDir.name, " ", subDir.size)
		if len(subDir.subDirectories) > 0 {
			sum += subDir.sumUnderLimit(limit)
		} else if subDir.size < limit {
			sum += subDir.size
		}
	}
	return sum
}

func (d *Directory) smallestSubDirSize(targetSize int) int {
	size := d.size
	if d.size > targetSize {
		for _, subDir := range d.subDirectories {
			if subDir.size > targetSize && subDir.size < size {
				size = subDir.size
			}
			if len(subDir.subDirectories) > 0 {
				subSize := subDir.smallestSubDirSize(targetSize)
				if subSize > targetSize && subSize < size {
					size = subSize
				}
			}
		}
	}
	return size
}
