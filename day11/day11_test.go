package day11

import (
	"adventofcode/helpers"
	"fmt"
	"reflect"
	"testing"
)

func TestMonkeyBusiness(t *testing.T) {
	t.Run("will investigate item", func(t *testing.T) {
		monkey := Monkey{}

		dataset := map[string]int{
			"+ 1":   2,
			"+ 2":   3,
			"* 1":   1,
			"* 5":   5,
			"* old": 1,
			"+ old": 2,
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

	t.Run("will process all items", func(t *testing.T) {
		monkey1 := Monkey{}
		monkey2 := Monkey{}

		monkey0 := Monkey{
			Items:     []int{79, 60, 97},
			Operation: "* old",
			Test:      13,
			Target1:   &monkey1,
			Target2:   &monkey2,
		}

		monkey0.ProcessAllItems()

		AssertItemCount(t, monkey0, 0)
		AssertItemCount(t, monkey1, 1)
		AssertItemCount(t, monkey2, 2)

		AssertItems(t, monkey1, []int{2080})
		AssertItems(t, monkey2, []int{1200, 3136})
	})

	t.Run("will create monkey squad", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		monkeys := CreateMonkeySquad(input)

		wantCount := 4
		if len(monkeys) != wantCount {
			t.Errorf("expected squad size of %d, got %d", wantCount, len(monkeys))
		}

		want := []Monkey{
			{Items: []int{79, 98}, Operation: "* 19", Test: 23},
			{Items: []int{54, 65, 75, 74}, Operation: "+ 6", Test: 19},
			{Items: []int{79, 60, 97}, Operation: "* old", Test: 13},
			{Items: []int{74}, Operation: "+ 3", Test: 17},
		}

		if !reflect.DeepEqual(want, monkeys) {
			t.Errorf("expected different squad")
			fmt.Printf("got squad %v\n", monkeys)
			fmt.Printf("wanted squad %v\n", want)
		}
	})
}

func AssertItems(t *testing.T, m Monkey, items []int) {
	t.Helper()

	if !reflect.DeepEqual(m.Items, items) {
		t.Fatalf("expected monkey to have %v items, got %v", items, m.Items)
	}
}

func AssertItemCount(t *testing.T, m Monkey, count int) {
	t.Helper()

	if len(m.Items) != count {
		t.Fatalf("expected monkey to have %d items, got %d", count, len(m.Items))
	}
}
