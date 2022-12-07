package day07

type Item interface {
	GetSize() int
}

type Dir struct {
	Name     string
	Children []Item
}

func (f Dir) GetSize() int {
	size := 0

	for _, file := range f.Children {
		size += file.GetSize()
	}

	return size
}

func (f *Dir) AddFile(name string, size int) {
	f.Children = append(f.Children, File{name, size})
}

type File struct {
	Name string
	Size int
}

func (f File) GetSize() int {
	return f.Size
}
func (r *Dir) Add(child Item) {
	r.Children = append(r.Children, child)
}
