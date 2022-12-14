package day03

import (
	"adventofcode/helpers"
	"reflect"
	"testing"
)

func TestRucksackReorganisation(t *testing.T) {
	t.Run("part 1 example case", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")

		want := 157
		got := Part1(input)

		if got != want {
			t.Errorf("expected score of %d, got %d", want, got)
		}
	})

	t.Run("part 2 example case", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")

		want := 70
		got := Part2(input)

		if got != want {
			t.Errorf("expected score of %d, got %d", want, got)
		}
	})

	t.Run("split items in a rucksack in two", func(t *testing.T) {
		dp := map[string][]string{
			"vJrwpWtwJgWrhcsFMMfFFhFp":         {"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL": {"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
			"PmmdzqPrVvPwwTWBwg":               {"PmmdzqPrV", "vPwwTWBwg"},
		}

		for input, want := range dp {
			got := SplitItems(input)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("%s: expected %v, got %v", input, want, got)
			}
		}
	})

	t.Run("find common item", func(t *testing.T) {
		dataset := map[string]string{
			"vJrwpWtwJgWrhcsFMMfFFhFp":         "p",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL": "L",
			"PmmdzqPrVvPwwTWBwg":               "P",
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn":   "v",
			"ttgJtRGJQctTZtZT":                 "t",
			"CrZsJsPPZsGzwwsLwLmpwMDw":         "s",
		}

		for input, want := range dataset {
			got := FindCommonItem(input)

			if got != want {
				t.Errorf("expected %q, got %q", want, got)
			}
		}
	})

	t.Run("split elves into groups of 3", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")

		want := [][]string{
			{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"},
			{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
		}

		got := SplitGroups(input)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %v, got %v", want, got)
		}
	})

	t.Run("find common badge", func(t *testing.T) {
		sets := map[string][]string{
			"r": {"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"},
			"Z": {"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
		}

		for want, input := range sets {
			got := FindBadge(input)
			if got != want {
				t.Errorf("expected %q, got %q", want, got)
			}
		}
	})
}
