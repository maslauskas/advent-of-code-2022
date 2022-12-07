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
		root := Directory{
			"x": Directory{},
		}

		want := 1
		got := len(root)

		if want != got {
			t.Errorf("expected dir count to be %d, got %d", want, got)
		}
	})

	t.Run("will create file within root", func(t *testing.T) {
		root := Directory{
			"a": File(123),
		}

		want := 1
		got := len(root)

		if want != got {
			t.Errorf("expected file count to be %d, got %d", want, got)
		}
	})

	t.Run("will add multiple files and folders to dir", func(t *testing.T) {
		root := Directory{
			"x": Directory{},
			"y": Directory{},
			"a": File(100),
			"b": File(200),
			"c": File(100),
		}

		wantCount := 5
		gotCount := len(root)

		if wantCount != gotCount {
			t.Errorf("expected dir count to be %d, gotCount %d", wantCount, gotCount)
		}
	})

	t.Run("will get file size", func(t *testing.T) {
		file := File(123)
		want := 123
		got := file.GetSize()

		if want != got {
			t.Errorf("expected file size to be %d, got %d", want, got)
		}
	})

	t.Run("folder size is 0 when folder is empty", func(t *testing.T) {
		dir := Directory{}

		want := 0
		got := dir.GetSize()

		if want != got {
			t.Errorf("expected dir size to be %d, got %d", want, got)
		}
	})

	t.Run("folder size is sum of file sizes when folder is not empty", func(t *testing.T) {
		dir := Directory{
			"y": File(150),
		}

		want := 150
		got := dir.GetSize()

		if want != got {
			t.Errorf("expected dir size to be %d, got %d", want, got)
		}
	})

	t.Run("will get size of folder with 2 levels of subfolders", func(t *testing.T) {
		root := Directory{
			"x": Directory{
				"y": Directory{
					"a": File(250),
				},
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
		want := Directory{
			"/a":       Directory{},
			"/b.txt":   File(14848514),
			"/c.dat":   File(8504156),
			"/d":       Directory{},
			"/a/e":     Directory{},
			"/a/f":     File(29116),
			"/a/g":     File(2557),
			"/a/h.lst": File(62596),
			"/a/e/i":   File(584),
			"/d/j":     File(4060174),
			"/d/d.log": File(8033020),
			"/d/d.ext": File(5626152),
			"/d/k":     File(7214296),
		}

		if len(got) != len(want) {
			t.Errorf("expected root to have %d children, got %d", len(want), len(got))
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected different tree, got %v, wanted %v", got, want)
		}
	})
}
