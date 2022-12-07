package day07

import (
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
	//dir := root

	for _, line := range input {
		if line == "$ cd /" {
			//dir = root
			continue
		}

		if line == "$ ls" {
			continue
		}

		//parts := strings.Split(line, " ")

		if strings.HasPrefix(line, "$ cd") {
			//target := parts[2]
			//dir = dir[target]
			//continue
			break // only go 1 level deep for now
		}

		//size := parts[0]
		//name := parts[1]

		//if size == "dir" {
		//	dir[name] = Directory{}
		//} else {
		//	s, _ := strconv.Atoi(size)
		//	dir[name] =
		//}
	}

	return root
}
