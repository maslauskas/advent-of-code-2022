package day07

type Directory struct {
	Name  string
	Files map[string]int
}

func (d Directory) Size() int {
	sum := 0
	for _, size := range d.Files {
		sum += size
	}
	return sum
}

func (d Directory) AddFile(name string, size int) {
	d.Files[name] = size
}

func MakeDir(name string) Directory {
	return Directory{
		name,
		map[string]int{},
	}
}
