package day05

import (
	"adventofcode/helpers"
	"reflect"
	"testing"
)

func TestSupplyStacks(t *testing.T) {
	t.Run("part 1 example case", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		got := Part1(input)
		want := "CMZ"

		if want != got {
			t.Errorf("expected result %q, got %q", want, got)
		}

	})

	t.Run("part 2 example case", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		got := Part2(input)
		want := "MCD"

		if want != got {
			t.Errorf("expected result %q, got %q", want, got)
		}

	})

	t.Run("builds stack collection", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		stacks := map[int][]string{
			1: {"Z", "N"},
			2: {"M", "C", "D"},
			3: {"P"},
		}
		instructions := []string{
			"move 1 from 2 to 1",
			"move 3 from 1 to 3",
			"move 2 from 2 to 1",
			"move 1 from 1 to 2",
		}
		want := StackCollection{stacks, instructions}
		got := CreateStackCollection(input)

		if !reflect.DeepEqual(got.Stacks, want.Stacks) {
			t.Errorf("wanted stacks to be %v, got %v", want.Stacks, got.Stacks)
		}

		if !reflect.DeepEqual(got.Instructions, want.Instructions) {
			t.Errorf("wanted instructions to be %v, got %v", want.Instructions, got.Instructions)
		}
	})

	t.Run("will parse instructions", func(t *testing.T) {
		dataset := map[string][]int{
			"move 1 from 2 to 1":       {1, 2, 1},
			"move 3 from 1 to 3":       {3, 1, 3},
			"move 2 from 2 to 1":       {2, 2, 1},
			"move 1 from 1 to 2":       {1, 1, 2},
			"move 99 from 100 to 2987": {99, 100, 2987},
		}

		for inst, want := range dataset {
			count, from, to := ParseInstructions(inst)
			got := []int{count, from, to}
			if !reflect.DeepEqual(want, got) {
				t.Errorf("dataset %q: expected %v, got %v", inst, want, got)
			}
		}
	})

	t.Run("will execute 1 instruction", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		stacks := CreateStackCollection(input)

		stacks.Move(1)

		want := map[int][]string{
			1: {"Z", "N", "D"},
			2: {"M", "C"},
			3: {"P"},
		}

		if !reflect.DeepEqual(stacks.Stacks, want) {
			t.Errorf("wanted stacks to be %v, got %v", want, stacks.Stacks)
		}
	})

	t.Run("will move retaining the order", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		stacks := CreateStackCollection(input)

		stacks.MoveAllStacks()

		want := map[int][]string{
			1: {"M"},
			2: {"C"},
			3: {"P", "Z", "N", "D"},
		}

		if !reflect.DeepEqual(stacks.Stacks, want) {
			t.Errorf("wanted stacks to be %v, got %v", want, stacks.Stacks)
		}
	})

	t.Run("will execute all instructions", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		stacks := CreateStackCollection(input)

		stacks.MoveAll()

		want := map[int][]string{
			1: {"C"},
			2: {"M"},
			3: {"P", "D", "N", "Z"},
		}

		if !reflect.DeepEqual(stacks.Stacks, want) {
			t.Errorf("wanted stacks to be %v, got %v", want, stacks.Stacks)
		}
	})
}
