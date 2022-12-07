package day07

import (
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("directory can have a size", func(t *testing.T) {
		dir := Node{Path: "x"}
		size := dir.Size()

		if size != 0 {
			t.Errorf("expected size to be 0, got %d", size)
		}
	})

	t.Run("directory size includes files size within it", func(t *testing.T) {
		file := Node{
			Path:     "x.txt",
			FileSize: 100,
		}
		dir := Node{
			Path: "x",
			Children: []*Node{
				&file,
			},
		}

		got := dir.Size()
		want := 100

		if got != want {
			t.Errorf("expected got to be %d, got %d", want, got)
		}
	})

	t.Run("directory size includes files and subdirs size within it", func(t *testing.T) {
		dir := Node{
			Path: "x",
			Children: []*Node{
				{
					Path:     "x.txt",
					FileSize: 100,
				},
				{
					Path: "x.txt",
					Children: []*Node{
						{
							Path:     "b.txt",
							FileSize: 200,
						},
					},
				},
			},
		}

		size := dir.Size()
		want := 300

		if size != want {
			t.Errorf("expected size to be %d, got %d", want, size)
		}
	})

	t.Run("will find node by path", func(t *testing.T) {
		t.Errorf("did not implement yet")
	})
}
