package day08

import (
	"adventofcode/helpers"
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

		if got == true {
			t.Errorf("expected result to be %v, got %v", false, got)
		}
	})

	t.Run("middle tree blocked", func(t *testing.T) {
		board := Board{
			[]int{1, 1, 1},
			[]int{1, 0, 1},
			[]int{1, 1, 1},
		}

		got := board.CheckBlocked(1, 1)

		if got == false {
			t.Errorf("expected result to be %v, got %v", true, got)
		}
	})

	t.Run("middle tree unblocked from top", func(t *testing.T) {
		board := Board{
			[]int{1, 0, 1},
			[]int{1, 1, 1},
			[]int{1, 1, 1},
		}

		got := board.CheckBlocked(1, 1)

		if got == true {
			t.Errorf("expected result to be %v, got %v", true, got)
		}
	})

	t.Run("middle tree unblocked from right", func(t *testing.T) {
		board := Board{
			[]int{1, 1, 1},
			[]int{1, 1, 0},
			[]int{1, 1, 1},
		}

		got := board.CheckBlocked(1, 1)

		if got == true {
			t.Errorf("expected result to be %v, got %v", true, got)
		}
	})

	t.Run("middle tree unblocked from bottom", func(t *testing.T) {
		board := Board{
			[]int{1, 1, 1},
			[]int{1, 1, 1},
			[]int{1, 0, 1},
		}

		got := board.CheckBlocked(1, 1)

		if got == true {
			t.Errorf("expected result to be %v, got %v", true, got)
		}
	})

	t.Run("middle tree unblocked from left", func(t *testing.T) {
		board := Board{
			[]int{1, 1, 1},
			[]int{0, 1, 1},
			[]int{1, 1, 1},
		}

		got := board.CheckBlocked(1, 1)

		if false != got {
			t.Errorf("expected result to be %v, got %v", false, got)
		}
	})
}

func TestPart1(t *testing.T) {
	input := helpers.ReadInput("./example.txt")
	got := Part1(input)
	want := 21

	if want != got {
		t.Errorf("expected %d, got %d", want, got)
	}
}

func TestScenicScore(t *testing.T) {
	t.Run("example case 1", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		board := BuildBoard(input)

		got := board.ScenicScore(1, 2)
		want := 4

		if want != got {
			t.Errorf("expected %d, got %d", want, got)
		}
	})

	t.Run("example case 2", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		board := BuildBoard(input)

		got := board.ScenicScore(3, 2)
		want := 8

		if want != got {
			t.Errorf("expected %d, got %d", want, got)
		}
	})
}

func TestPart2(t *testing.T) {
	input := helpers.ReadInput("./example.txt")
	got := Part2(input)
	want := 8

	if want != got {
		t.Errorf("expected %d, got %d", want, got)
	}
}
