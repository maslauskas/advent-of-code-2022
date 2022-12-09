package day09

import (
	"adventofcode/helpers"
	"testing"
)

func TestRopeBridge(t *testing.T) {
	t.Run("will move the head", func(t *testing.T) {

		dataset := map[string]Point{
			"U 1": {0, 1},
			"D 1": {0, -1},
			"R 1": {1, 0},
			"L 1": {-1, 0},
		}

		for set, point := range dataset {
			game := MakeRope(Point{0, 0}, 2)
			game.Move(set)

			tail := game.GetHead()
			if tail != point {
				t.Errorf("set %q: expected possition to be %v, got %v", set, point, tail)
			}
		}
	})

	t.Run("will move the tail when moving straight", func(t *testing.T) {
		dataset := map[string]Point{
			"U 1": {0, 0},
			"U 5": {0, 4},
			"D 5": {0, -4},
			"R 5": {4, 0},
			"L 5": {-4, 0},
		}

		for set, point := range dataset {
			game := MakeRope(Point{0, 0}, 2)
			game.Move(set)

			tail := game.GetTail()
			if tail != point {
				t.Errorf("set %q: expected possition to be %v, got %v", set, point, tail)
			}
		}
	})

	t.Run("will move the tail when moving diagonally", func(t *testing.T) {
		dataset := map[string]Point{
			"U 3": {1, 2},
			"D 3": {1, -2},
			"R 3": {2, 1},
			"L 3": {-2, 1},
		}

		for set, point := range dataset {
			game := MakeRope(Point{1, 0}, 2)
			game.Move(set)

			tail := game.GetTail()
			if tail != point {
				t.Errorf("set %q: expected possition to be %v, got %v", set, point, tail)
			}
		}
	})

	// tail does not move when head moves on top of tail;
	t.Run("tail does not move if head moves on top of tail", func(t *testing.T) {
		game := MakeRope(Point{0, -1}, 2)
		game.MoveOnce("U")

		AssertPosition(t, game.GetHead(), 0, 0)
		AssertPosition(t, game.GetTail(), 0, 0)
	})

	t.Run("tail does not move if head moves on top of tail", func(t *testing.T) {
		game := MakeRope(Point{0, 1}, 2)
		game.MoveOnce("D")

		AssertPosition(t, game.GetHead(), 0, 0)
		AssertPosition(t, game.GetTail(), 0, 0)
	})

	t.Run("tail does not move if head moves on top of tail", func(t *testing.T) {
		game := MakeRope(Point{-1, 0}, 2)
		game.MoveOnce("R")

		AssertPosition(t, game.GetHead(), 0, 0)
		AssertPosition(t, game.GetTail(), 0, 0)
	})

	t.Run("tail does not move if head moves on top of tail", func(t *testing.T) {
		game := MakeRope(Point{1, 0}, 2)
		game.MoveOnce("L")

		AssertPosition(t, game.GetHead(), 0, 0)
		AssertPosition(t, game.GetTail(), 0, 0)
	})

	// tail does not move when head moves diagonally one step;
	t.Run("tail does not move if head moves one step diagonally", func(t *testing.T) {
		game := MakeRope(Point{1, 0}, 2)
		game.MoveOnce("U")

		AssertPosition(t, game.GetHead(), 1, 1)
		AssertPosition(t, game.GetTail(), 0, 0)
	})

	t.Run("tail does not move if head moves one step diagonally", func(t *testing.T) {
		game := MakeRope(Point{1, 0}, 2)
		game.MoveOnce("D")

		AssertPosition(t, game.GetHead(), 1, -1)
		AssertPosition(t, game.GetTail(), 0, 0)
	})

	t.Run("tail does not move if head moves one step diagonally", func(t *testing.T) {
		game := MakeRope(Point{0, 1}, 2)
		game.MoveOnce("R")

		AssertPosition(t, game.GetHead(), 1, 1)
		AssertPosition(t, game.GetTail(), 0, 0)
	})

	t.Run("tail does not move if head moves one step diagonally", func(t *testing.T) {
		game := MakeRope(Point{0, 1}, 2)
		game.MoveOnce("L")

		AssertPosition(t, game.GetHead(), -1, 1)
		AssertPosition(t, game.GetTail(), 0, 0)
	})

	t.Run("will run example instructions", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		game := MakeRope(Point{0, 0}, 2)

		for _, row := range input {
			game.Move(row)
		}

		AssertPosition(t, game.GetHead(), 2, 2)
		AssertPosition(t, game.GetTail(), 1, 2)
	})

	t.Run("part 1 example case", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		got := Part1(input)

		want := 13
		if got != want {
			t.Errorf("expected result to be %d, got %d", want, got)
		}
	})

	t.Run("segment will catch up to its head", func(t *testing.T) {
		dataset := map[string][2]Point{
			"move up by 1":               {Point{0, 1}, Point{0, 0}},
			"move down by 1":             {Point{0, -1}, Point{0, 0}},
			"move right by 1":            {Point{1, 0}, Point{0, 0}},
			"move left by 1":             {Point{-1, 0}, Point{0, 0}},
			"move up by 2":               {Point{0, 2}, Point{0, 1}},
			"move right by 2":            {Point{2, 0}, Point{1, 0}},
			"move down by 2":             {Point{0, -2}, Point{0, -1}},
			"move left by 2":             {Point{-2, 0}, Point{-1, 0}},
			"move diagonally up by 1":    {Point{1, 2}, Point{1, 1}},
			"move diagonally right by 1": {Point{2, -1}, Point{1, -1}},
			"move diagonally down by 1":  {Point{-1, -2}, Point{-1, -1}},
			"move diagonally left by 1":  {Point{-2, 1}, Point{-1, 1}},
		}

		for set, points := range dataset {
			head := points[0]
			want := points[1]
			tail := Point{0, 0}

			tail.CatchUp(head)

			if tail != want {
				t.Errorf("set %q: expected possition to be %v, got %v", set, want, tail)
			}
		}
	})

	t.Run("part 2 example case step by step", func(t *testing.T) {
		game := MakeRope(Point{0, 0}, 10)

		game.Move("R 5")
		AssertPosition(t, game.GetHead(), 5, 0)
		AssertPosition(t, game.GetTail(), 0, 0)

		game.Move("U 8")
		AssertPosition(t, game.GetHead(), 5, 8)
		AssertPosition(t, game.GetTail(), 0, 0)

		game.Move("L 8")
		AssertPosition(t, game.GetHead(), -3, 8)
		AssertPosition(t, game.GetTail(), 1, 3)

		game.Move("D 3")
		AssertPosition(t, game.GetHead(), -3, 5)
		AssertPosition(t, game.GetTail(), 1, 3)

		game.Move("R 17")
		AssertPosition(t, game.GetHead(), 14, 5)
		AssertPosition(t, game.GetTail(), 5, 5)

		game.Move("D 10")
		AssertPosition(t, game.GetHead(), 14, -5)
		AssertPosition(t, game.GetTail(), 10, 0)

		game.Move("L 25")
		AssertPosition(t, game.GetHead(), -11, -5)
		AssertPosition(t, game.GetTail(), -2, -5)

		game.Move("U 20")
		AssertPosition(t, game.GetHead(), -11, 15)
		AssertPosition(t, game.GetTail(), -11, 6)

	})

	t.Run("part 2 example case", func(t *testing.T) {
		input := helpers.ReadInput("./example2.txt")
		got := Part2(input)

		want := 36
		if got != want {
			t.Errorf("expected result to be %d, got %d", want, got)
		}
	})
}

func AssertPosition(t *testing.T, p Point, x int, y int) {
	t.Helper()

	if p.PosX != x || p.PosY != y {
		t.Errorf("expected possition to be {%d %d}, got %v", x, y, p)
	}
}
