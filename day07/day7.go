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
	Children Children
}

type Directory map[string]any

func (f Dir) GetSize() int {
	size := 0

	for _, file := range f.Children {
		size += file.GetSize()
	}

	return size
}

type File struct {
	Size int
}

func (f File) GetSize() int {
	return f.Size
}

func BuildDirTree(input []string) Directory {
	root := Directory{}
	dir := root

	for _, line := range input {
		if line == "$ cd /" {
			dir = root
			continue
		}

		if line == "$ ls" {
			continue
		}

		if strings.HasPrefix(line, "$ cd") {
			break // only go 1 level deep for now
		}

		parts := strings.Split(line, " ")
		size := parts[0]
		name := parts[1]

		if size == "dir" {
			dir[name] = Directory{}
		} else {
			s, _ := strconv.Atoi(size)
			dir[name] = s
		}
	}

	return root
}
