package day04

import (
	"adventofcode/helpers"
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("part 1 example case", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		got := Part1(input)
		want := 2

		if got != want {
			t.Errorf("expected score to be %d, got %d", want, got)
		}
	})

	t.Run("part 2 example case", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		got := Part2(input)
		want := 4

		if got != want {
			t.Errorf("expected score to be %d, got %d", want, got)
		}
	})

	t.Run("will convert to sections", func(t *testing.T) {
		want := Section{2, 4}
		got := MakeSection("2-4")

		if got != want {
			t.Errorf("expected section %v, got %v", want, got)
		}
	})

	t.Run("sections will intersect", func(t *testing.T) {
		datasets := map[string]bool{
			"2-4,6-8": false,
			"2-3,4-5": false,
			"5-7,7-9": true,
			"2-8,3-7": true,
			"6-6,4-6": true,
			"2-6,4-8": true,
		}

		for input, want := range datasets {
			s1, s2 := MakeSections(input)

			contains := s1.Intersects(s2)
			if contains != want {
				t.Errorf("expected %v intersect %v to be %v", s1, s2, want)
			}
		}
	})

	t.Run("sections can contain one another", func(t *testing.T) {
		datasets := map[string]bool{
			"2-4,6-8": false,
			"2-3,4-5": false,
			"5-7,7-9": false,
			"2-8,3-7": true,
			"6-6,4-6": true,
			"2-6,4-8": false,
		}

		for input, want := range datasets {
			s1, s2 := MakeSections(input)

			contains := s1.Overlaps(s2)
			if contains != want {
				t.Errorf("expected %v overlap %v to be %v", s1, s2, want)
			}
		}
	})
}
