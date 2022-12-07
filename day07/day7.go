package day07

import (
	"strconv"
	"strings"
)

type File int
type FileOrDirectory interface {
	GetSize() int
}

type Directory map[string]FileOrDirectory

func (d Directory) GetSize() int {
	size := 0

	for _, file := range d {
		size += file.GetSize()
	}

	return size
}

func (f File) GetSize() int {
	return int(f)
}

func BuildDirTree(input []string) Directory {
	root := Directory{}

	for _, line := range input {
		if line == "$ cd /" {
			continue
		}

		if line == "$ ls" {
			continue
		}

		parts := strings.Split(line, " ")

		if strings.HasPrefix(line, "$ cd") {
			//target := parts[2]
			//continue
			break // only go 1 level deep for now
		}

		size := parts[0]
		name := parts[1]

		if size == "dir" {
			root[name] = Directory{}
		} else {
			s, _ := strconv.Atoi(size)
			root[name] = File(s)
		}
	}

	return root
}
