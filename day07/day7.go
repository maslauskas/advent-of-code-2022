package day07

type Node struct {
	Name     string
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
