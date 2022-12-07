package day07

import (
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("directory can have a size", func(t *testing.T) {
		dir := MakeDir("x")
		size := dir.Size()

		if size != 0 {
			t.Errorf("expected size to be 0, got %d", size)
		}
	})

	t.Run("directory size includes files size within it", func(t *testing.T) {
		dir := MakeDir("x")
		dir.AddFile("a.txt", 100)

		size := dir.Size()
		want := 100

		if size != want {
			t.Errorf("expected size to be %d, got %d", want, size)
		}
	})

	t.Run("directory size includes files and subdirs size within it", func(t *testing.T) {
		dir := MakeDir("x")
		dir.AddFile("a.txt", 100)

		subdir := MakeDir("b")
		subdir.AddFile("b.txt", 200)

		dir.AddDir(subdir)

		size := dir.Size()
		want := 300

		if size != want {
			t.Errorf("expected size to be %d, got %d", want, size)
		}
	})
}
