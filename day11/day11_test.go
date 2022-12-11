package day11

import (
	"testing"
)

func TestMonkeyBusiness(t *testing.T) {
	t.Run("will investigate item", func(t *testing.T) {
		monkey := Monkey{}

		dataset := map[string]int{
			"+ 1": 2,
			"+ 2": 3,
			"* 1": 1,
			"* 5": 5,
		}

		for operation, want := range dataset {
			got := monkey.Investigate(1, operation)

			if want != got {
				t.Errorf("expected item value to be %d, got %d", want, got)
			}
		}
	})

	t.Run("will get bored", func(t *testing.T) {
		monkey := Monkey{}

		dataset := [][2]int{
			{1, 0},
			{2, 0},
			{3, 1},
			{4, 1},
			{5, 1},
			{1862, 620},
		}

		for _, data := range dataset {
			value := data[0]
			want := data[1]
			got := monkey.GetBored(value)

			if want != got {
				t.Errorf("expected item value to be %d, got %d", want, got)
			}
		}
	})
}
