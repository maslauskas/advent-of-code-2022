package day07

type Directory struct {
	Name    string
	Subdirs []Directory
	Files   map[string]int
}

func (d Directory) Size() int {
	sum := 0
	for _, size := range d.Files {
		sum += size
	}

	for _, subdir := range d.Subdirs {
		sum += subdir.Size()
	}

	return sum
}

func (d Directory) AddFile(name string, size int) {
	d.Files[name] = size
}

func (d *Directory) AddDir(subdir Directory) {
	d.Subdirs = append(d.Subdirs, subdir)
}

func MakeDir(name string) Directory {
	return Directory{
		name,
		[]Directory{},
		map[string]int{},
	}
}
