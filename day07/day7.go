package day07

import (
	"sort"
	"strconv"
	"strings"
)

type File int
type FileOrDirectory interface {
	GetSize() int
	IsDir() bool
	AddChild(name string, child FileOrDirectory)
}

type Directory map[string]FileOrDirectory

func (d Directory) GetSize() int {
	size := 0

	for _, file := range d {
		size += file.GetSize()
	}

	return size
}

func (d Directory) IsDir() bool {
	return true
}

func (f File) IsDir() bool {
	return false
}

func (d Directory) AddChild(path string, child FileOrDirectory) {
	d[path] = child
}

func (f File) AddChild(_ string, _ FileOrDirectory) {
	panic("cannot add child to file")
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

func BuildBetterDirTree(input []string) Directory {
	simpleTree := BuildDirTree(input)
	tree := Directory{}

	// sort map keys
	keys := make([]string, 0)
	for k := range simpleTree {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		// take key from simple map
		file, ok := simpleTree[key]
		if !ok {
			// item gone, move on
			continue
		}

		delete(simpleTree, key)
		simpleTree, file = FindChildren(file, simpleTree, key)
		tree.AddChild(key, file)
	}

	return tree
}

func FindChildren(tree FileOrDirectory, simpleTree Directory, path string) (Directory, FileOrDirectory) {
	if !tree.IsDir() {
		return simpleTree, tree
	}

	// sort map keys
	keys := make([]string, 0)
	for k := range simpleTree {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, key := range keys {
		if path == key || !strings.HasPrefix(key, path) {
			// item not child, move on
			continue
		}

		file, ok := simpleTree[key]
		if !ok {
			// item gone, move on
			continue
		}

		delete(simpleTree, key)
		simpleTree, file = FindChildren(file, simpleTree, key)
		tree.AddChild(key, file)
	}
	return simpleTree, tree
}

func Part1(input []string) int {
	tree := BuildBetterDirTree(input)
	sum := 0

	for _, file := range tree {
		if file.GetSize() <= 100000 {
			sum += file.GetSize()
		}
	}

	return sum
}
