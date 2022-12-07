package day07

type Folder struct {
	Name string
}

type RootDir struct {
	Folders []Folder
}

func (r *RootDir) addFolder(name string) {
	r.Folders = append(r.Folders, Folder{name})
}
