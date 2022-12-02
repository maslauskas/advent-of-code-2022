package day02

import (
	"reflect"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	t.Run("calculates total score", func(t *testing.T) {
		want := 15
		got := CalculateTotalScore("./example.txt")

		if want != got {
			t.Errorf("expected total score %v, got %v", want, got)
		}
	})

	t.Run("calculates each row", func(t *testing.T) {
		want := []int{8, 1, 6}
		got := CalculateRowsScore("./example.txt")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected rows score to be %v, got %v", want, got)
		}
	})

	t.Run("calculate total score based on desired endings", func(t *testing.T) {
		want := 12
		got := PredictEndings("./example.txt")

		if want != got {
			t.Errorf("expected total score %v, got %v", want, got)
		}
	})
}
