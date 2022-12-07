package day07

import "testing"

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
}
