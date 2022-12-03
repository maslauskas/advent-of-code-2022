package day03

import (
	"reflect"
	"testing"
)

func TestRucksackReorganisation(t *testing.T) {
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
}
