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
