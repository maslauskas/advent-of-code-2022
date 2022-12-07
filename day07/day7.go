package day07

type Item interface {
	GetSize() int
	GetName() string
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
	Name string
	Size int
}

func (f File) GetSize() int {
	return f.Size
}

func (f File) GetName() string {
	return f.Name
}

func (r *Dir) Add(child Item) {
	r.Children[child.GetName()] = child
}
