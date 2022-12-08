package day08

import (
	"reflect"
	"testing"
)

func TestBuildBoard(t *testing.T) {
	input := []string{"00", "00"}
	board := BuildBoard(input)
	want := Board{
		[]int{0, 0},
		[]int{0, 0},
	}

	if !reflect.DeepEqual(want, board) {
		t.Errorf("got %v, expected %v", board, want)
	}
}

func TestCheckBlockingTrees(t *testing.T) {
	t.Run("corner tree", func(t *testing.T) {
		board := Board{
			[]int{0, 1},
			[]int{1, 0},
		}

		got := board.CheckBlocked(0, 0)
		want := false

		if want != got {
			t.Errorf("expected result to be %v, got %v", want, got)
		}
	})

	t.Run("middle tree blocked", func(t *testing.T) {
		board := Board{
			[]int{1, 1, 1},
			[]int{1, 0, 1},
			[]int{1, 1, 1},
		}

		got := board.CheckBlocked(1, 1)
		want := true

		if want != got {
			t.Errorf("expected result to be %v, got %v", want, got)
		}
	})

	t.Run("middle tree unblocked from top", func(t *testing.T) {
		board := Board{
			[]int{1, 0, 1},
			[]int{1, 0, 1},
			[]int{1, 1, 1},
		}

		got := board.CheckBlocked(1, 1)
		want := false

		if want != got {
			t.Errorf("expected result to be %v, got %v", want, got)
		}
	})

	t.Run("middle tree unblocked from right", func(t *testing.T) {
		board := Board{
			[]int{1, 1, 1},
			[]int{1, 0, 0},
			[]int{1, 1, 1},
		}

		got := board.CheckBlocked(1, 1)
		want := false

		if want != got {
			t.Errorf("expected result to be %v, got %v", want, got)
		}
	})

	t.Run("middle tree unblocked from bottom", func(t *testing.T) {
		board := Board{
			[]int{1, 1, 1},
			[]int{1, 0, 1},
			[]int{1, 0, 1},
		}

		got := board.CheckBlocked(1, 1)
		want := false

		if want != got {
			t.Errorf("expected result to be %v, got %v", want, got)
		}
	})

	t.Run("middle tree unblocked from left", func(t *testing.T) {
		board := Board{
			[]int{1, 1, 1},
			[]int{0, 0, 1},
			[]int{1, 1, 1},
		}

		got := board.CheckBlocked(1, 1)
		want := false

		if want != got {
			t.Errorf("expected result to be %v, got %v", want, got)
		}
	})
}
