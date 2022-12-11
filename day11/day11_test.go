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

	t.Run("will throw item to another monkey based on test result", func(t *testing.T) {
		monkey0 := Monkey{}
		monkey1 := Monkey{}
		monkey2 := Monkey{}

		monkey0.Throw(2, &monkey1, &monkey2, 2)

		if len(monkey1.Items) == 0 {
			t.Errorf("expected item to be thrown to target, got %v", monkey1.Items)
		}

		monkey0.Throw(1, &monkey1, &monkey2, 2)

		if len(monkey2.Items) == 0 {
			t.Errorf("expected item to be thrown to target, got %v", monkey1.Items)
		}
	})

	t.Run("will process one item fully", func(t *testing.T) {
		monkey1 := Monkey{}
		monkey2 := Monkey{}

		monkey0 := Monkey{
			Operation: "* 19",
			Test:      23,
		}

		monkey0.ProcessItem(79, &monkey1, &monkey2)

		if len(monkey2.Items) == 0 {
			t.Fatalf("expected monkey2 to have new item")
		}

		want := 500
		got := monkey2.Items[0]

		if got != want {
			t.Errorf("expected item value to be %d, got %d", want, got)
		}
	})
}
