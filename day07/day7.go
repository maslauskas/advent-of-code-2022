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
	path := "/"

	for _, line := range input {
		if line == "$ cd /" {
			path = "/"
			continue
		}

		if line == "$ ls" {
			continue
		}

		parts := strings.Split(line, " ")

		if strings.HasPrefix(line, "$ cd") {
			target := parts[2]

			if target == ".." {
				segments := strings.Split(path, "/")
				segments = segments[:len(segments)-2]
				path = strings.Join(segments, "/") + "/"
			} else {
				path = path + target + "/"
			}

			continue
		}

		size := parts[0]
		name := path + parts[1]

		if size == "dir" {
			root[name] = Directory{}
		} else {
			s, _ := strconv.Atoi(size)
			root[name] = File(s)
		}
	}

	return root
}
