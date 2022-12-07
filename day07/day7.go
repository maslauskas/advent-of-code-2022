package day07

type Item interface {
	GetSize() int
}

type Folder struct {
	Name  string
	Files []File
}

func (f Folder) GetSize() int {
	size := 0

	for _, file := range f.Files {
		size += file.GetSize()
	}

	return size
}

func (f *Folder) AddFile(name string, size int) {
	f.Files = append(f.Files, File{name, size})
}

type File struct {
	Name string
	Size int
}

func (f File) GetSize() int {
	return f.Size
}

type RootDir struct {
	Children []Item
}

func (r *RootDir) Add(child Item) {
	r.Children = append(r.Children, child)
}
