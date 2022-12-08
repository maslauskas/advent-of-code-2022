package day07

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Path     string
	FileSize int
	Children []*Node
}

func (n Node) Size() int {
	sum := 0

	for _, child := range n.Children {
		sum += child.Size()
	}

	return sum + n.FileSize
}

func (n Node) Print() {
	fmt.Println(n.Path)
	for _, child := range n.Children {
		child.Print()
	}
}

func BuildTree(input []string) Node {
	path := "/"

	root := Node{Path: path}
	elements := map[string]*Node{
		path: &root,
	}

	var currentDir *Node
	currentDir = elements[path]

	for len(input) > 0 {
		command := input[0]
		input = input[1:]

		parts := strings.Split(command, " ")
		if strings.HasPrefix(command, "$ cd") {
			targetDir := parts[2]
			if targetDir == ".." {
				// go up one level
				segments := strings.Split(path, "/")
				segments = segments[:len(segments)-2]
				path = strings.Join(segments, "/") + "/"
			} else if targetDir != "/" {
				path += targetDir + "/"
			}
			currentDir = elements[path]
			continue
		}

		if command == "$ ls" {
			// skip listing
			continue
		}

		size := parts[0]
		name := path + parts[1]

		var s int
		if size == "dir" {
			s = 0
			name += "/"
		} else {
			s, _ = strconv.Atoi(size)
		}

		file := Node{Path: name, FileSize: s}
		elements[name] = &file

		currentDir.Children = append(currentDir.Children, &file)
	}

	return *elements["/"]
}
