package day07

import (
	"adventofcode/helpers"
	"reflect"
	"testing"
)

// + create folder
// + create file
// + get file size
// + create file inside folder and get its size
// + create folder with subfolders and get its size
// build dir tree

func TestExample(t *testing.T) {
	t.Run("will create folder within root", func(t *testing.T) {
		root := Dir{"/", Children{
			"x": Dir{"x", Children{}},
		}}

		want := 1
		got := len(root.Children)

		if want != got {
			t.Errorf("expected dir count to be %d, got %d", want, got)
		}
	})

	t.Run("will create file within root", func(t *testing.T) {
		root := Dir{"/", Children{
			"a": File{123},
		}}

		want := 1
		got := len(root.Children)

		if want != got {
			t.Errorf("expected file count to be %d, got %d", want, got)
		}
	})

	t.Run("will add multiple files and folders to dir", func(t *testing.T) {
		root := Dir{"/", Children{
			"x": Dir{"x", Children{}},
			"y": Dir{"y", Children{}},
			"a": File{100},
			"b": File{200},
			"c": File{300},
		}}

		wantCount := 5
		gotCount := len(root.Children)

		if wantCount != gotCount {
			t.Errorf("expected dir count to be %d, gotCount %d", wantCount, gotCount)
		}
	})

	t.Run("will get file size", func(t *testing.T) {
		file := File{123}
		want := 123
		got := file.GetSize()

		if want != got {
			t.Errorf("expected file size to be %d, got %d", want, got)
		}
	})

	t.Run("folder size is 0 when folder is empty", func(t *testing.T) {
		dir := Dir{"x", Children{}}

		want := 0
		got := dir.GetSize()

		if want != got {
			t.Errorf("expected dir size to be %d, got %d", want, got)
		}
	})

	t.Run("folder size is sum of file sizes when folder is not empty", func(t *testing.T) {
		dir := Dir{"x", Children{
			"y": File{150},
		}}

		want := 150
		got := dir.GetSize()

		if want != got {
			t.Errorf("expected dir size to be %d, got %d", want, got)
		}
	})

	t.Run("will get size of folder with 2 levels of subfolders", func(t *testing.T) {
		root := Dir{
			"/",
			Children{
				"x": Dir{"x", Children{
					"y": Dir{"y", Children{
						"a": File{250},
					}},
				}},
			},
		}

		want := 250
		got := root.GetSize()

		if want != got {
			t.Errorf("expected dir size to be %d, got %d", want, got)
		}
	})

	t.Run("will build directory tree", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		got := BuildDirTree(input)
		want := Dir{"/", Children{
			"a":     Dir{"a", Children{}},
			"b.txt": File{14848514},
			"c.dat": File{8504156},
			"d":     Dir{"d", Children{}},
		}}

		if len(got.Children) != 4 {
			t.Errorf("expected root to have 4 children, got %d", len(got.Children))
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected different tree, got %v", got)
		}
	})
}
