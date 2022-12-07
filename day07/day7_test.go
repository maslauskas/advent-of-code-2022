package day07

import "testing"

// + create folder
// + create file
// + get file size
// create file inside folder and get its size
// create folder with subfolders and get its size
// select folder

func TestExample(t *testing.T) {
	t.Run("will create folder within root", func(t *testing.T) {
		root := RootDir{}
		root.Add(NewDir("x"))

		want := 1
		got := len(root.Children)

		if want != got {
			t.Errorf("expected dir count to be %d, got %d", want, got)
		}
	})

	t.Run("will create file within root", func(t *testing.T) {
		root := RootDir{}
		root.Add(File{"a", 123})

		want := 1
		got := len(root.Children)

		if want != got {
			t.Errorf("expected file count to be %d, got %d", want, got)
		}
	})

	t.Run("will add multiple files and folders to dir", func(t *testing.T) {
		root := RootDir{}
		root.Add(NewDir("x"))
		root.Add(NewDir("y"))

		root.Add(File{"a", 100})
		root.Add(File{"b", 200})
		root.Add(File{"c", 300})

		wantCount := 5
		gotCount := len(root.Children)

		if wantCount != gotCount {
			t.Errorf("expected dir count to be %d, gotCount %d", wantCount, gotCount)
		}
	})

	t.Run("will get file size", func(t *testing.T) {
		file := File{"x", 123}
		want := 123
		got := file.GetSize()

		if want != got {
			t.Errorf("expected file size to be %d, got %d", want, got)
		}
	})

	t.Run("folder size is 0 when folder is empty", func(t *testing.T) {
		dir := Dir{"x", []Item{}}
		want := 0

		got := dir.GetSize()

		if want != got {
			t.Errorf("expected dir size to be %d, got %d", want, got)
		}
	})

	t.Run("folder size is sum of file sizes when folder is not empty", func(t *testing.T) {
		dir := Dir{"x", []Item{}}
		dir.AddFile("y", 150)
		want := 150

		got := dir.GetSize()

		if want != got {
			t.Errorf("expected dir size to be %d, got %d", want, got)
		}
	})
}

func NewDir(name string) Dir {
	return Dir{name, []Item{}}
}
