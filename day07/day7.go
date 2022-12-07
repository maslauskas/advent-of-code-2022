package day07

type Folder struct {
	Name string
}

type File struct {
	Name string
	Size int
}

func (f File) GetSize() int {
	return f.Size
}

type RootDir struct {
	Folders []Folder
	Files   []File
}

func (r *RootDir) addFolder(name string) {
	r.Folders = append(r.Folders, Folder{name})
}

func (r *RootDir) addFile(name string, size int) {
	r.Files = append(r.Files, File{name, size})
}
