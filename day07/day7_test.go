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
		root.addFolder("a")

		want := 1
		got := len(root.Folders)

		if want != got {
			t.Errorf("expected dir count to be %d, got %d", want, got)
		}
	})

	t.Run("will create file within root", func(t *testing.T) {
		root := RootDir{}
		root.addFile("a", 123)

		want := 1
		got := len(root.Files)

		if want != got {
			t.Errorf("expected file count to be %d, got %d", want, got)
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
}
