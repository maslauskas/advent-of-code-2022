package day07

import (
	"strconv"
	"strings"
)

type Item interface {
	GetSize() int
}

type Children map[string]Item

type Dir struct {
	Name     string
	Children Children
}

func (f Dir) GetSize() int {
	size := 0

	for _, file := range f.Children {
		size += file.GetSize()
	}

	return size
}

func (f Dir) GetName() string {
	return f.Name
}

type File struct {
	Size int
}

func (f File) GetSize() int {
	return f.Size
}

func BuildDirTree(input []string) Dir {
	root := Dir{"/", Children{}}

	for _, command := range input {
		if command == "$ cd /" {
			continue
		}

		if command == "$ ls" {
			continue
		}

		if strings.HasPrefix(command, "$ cd") {
			break // read only top level for now
		}

		parts := strings.Split(command, " ")
		size := parts[0]
		name := parts[1]

		if size == "dir" {
			root.Children[name] = Dir{name, Children{}}
		} else {
			s, _ := strconv.Atoi(size)
			root.Children[name] = File{s}
		}
	}

	return root
}
