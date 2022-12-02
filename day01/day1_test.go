package day01

import (
	"adventofcode/helpers"
	"testing"
)

func TestCalorieCounting(t *testing.T) {
	input := helpers.ReadInput("./example.txt")

	t.Run("part 1 example case", func(t *testing.T) {
		want := 24000
		got := Part1(input)

		if want != got {
			t.Errorf("expected %v, got %v", want, got)
		}
	})

	t.Run("part 2 example case", func(t *testing.T) {
		want := 45000
		got := Part2(input)

		if want != got {
			t.Errorf("expected %v, got %v", want, got)
		}
	})
}
