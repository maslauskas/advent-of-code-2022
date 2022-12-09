package day09

import (
	"adventofcode/helpers"
	"testing"
)

func TestRopeBridge(t *testing.T) {
	t.Run("will not move tail if head is on top and head moved by 1", func(t *testing.T) {
		game := MakeRope(Point{0, 0}, 2)
		game.MoveOnce("U")

		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("will move head up", func(t *testing.T) {
		game := MakeRope(Point{0, 0}, 2)
		game.MoveOnce("U")

		AssertPosition(t, game.Head, 0, 1)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("will move head down", func(t *testing.T) {
		game := MakeRope(Point{0, 0}, 2)
		game.MoveOnce("D")

		AssertPosition(t, game.Head, 0, -1)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("will move head right", func(t *testing.T) {
		game := MakeRope(Point{0, 0}, 2)
		game.Right(1)

		AssertPosition(t, game.Head, 1, 0)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("will move head left", func(t *testing.T) {
		game := MakeRope(Point{0, 0}, 2)
		game.Left(1)

		AssertPosition(t, game.Head, -1, 0)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("tail will catch up when moved straight up", func(t *testing.T) {
		game := MakeRope(Point{0, 0}, 2)
		game.Move("U 5")

		AssertPosition(t, game.Head, 0, 5)
		AssertPosition(t, game.Tail, 0, 4)
	})

	t.Run("tail will catch up when moved straight down", func(t *testing.T) {
		game := MakeRope(Point{0, 0}, 2)
		game.Move("D 5")

		AssertPosition(t, game.Head, 0, -5)
		AssertPosition(t, game.Tail, 0, -4)
	})

	t.Run("tail will catch up when moved straight right", func(t *testing.T) {
		game := MakeRope(Point{0, 0}, 2)
		game.Right(5)

		AssertPosition(t, game.Head, 5, 0)
		AssertPosition(t, game.Tail, 4, 0)
	})

	t.Run("tail will catch up when moved straight left", func(t *testing.T) {
		game := MakeRope(Point{0, 0}, 2)
		game.Left(5)

		AssertPosition(t, game.Head, -5, 0)
		AssertPosition(t, game.Tail, -4, 0)
	})

	// tail moves diagonally in all directions
	t.Run("tail will catch up when head is positioned diagonally and moved up", func(t *testing.T) {
		game := MakeRope(Point{1, 0}, 2)
		game.Move("U 3")

		AssertPosition(t, game.Head, 1, 3)
		AssertPosition(t, game.Tail, 1, 2)
	})

	t.Run("tail will catch up when head is positioned diagonally and moved down", func(t *testing.T) {
		game := MakeRope(Point{1, 0}, 2)
		game.Move("D 3")

		AssertPosition(t, game.Head, 1, -3)
		AssertPosition(t, game.Tail, 1, -2)
	})

	t.Run("tail will catch up when head is positioned diagonally and moved right", func(t *testing.T) {
		game := MakeRope(Point{0, 1}, 2)
		game.Right(3)

		AssertPosition(t, game.Head, 3, 1)
		AssertPosition(t, game.Tail, 2, 1)
	})

	t.Run("tail will catch up when head is positioned diagonally and moved left", func(t *testing.T) {
		game := MakeRope(Point{0, 1}, 2)
		game.Left(3)

		AssertPosition(t, game.Head, -3, 1)
		AssertPosition(t, game.Tail, -2, 1)
	})

	// tail does not move when head moves on top of tail;
	t.Run("tail does not move if head moves on top of tail", func(t *testing.T) {
		game := MakeRope(Point{0, -1}, 2)
		game.MoveOnce("U")

		AssertPosition(t, game.Head, 0, 0)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("tail does not move if head moves on top of tail", func(t *testing.T) {
		game := MakeRope(Point{0, 1}, 2)
		game.MoveOnce("D")

		AssertPosition(t, game.Head, 0, 0)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("tail does not move if head moves on top of tail", func(t *testing.T) {
		game := MakeRope(Point{-1, 0}, 2)
		game.Right(1)

		AssertPosition(t, game.Head, 0, 0)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("tail does not move if head moves on top of tail", func(t *testing.T) {
		game := MakeRope(Point{1, 0}, 2)
		game.Left(1)

		AssertPosition(t, game.Head, 0, 0)
		AssertPosition(t, game.Tail, 0, 0)
	})

	// tail does not move when head moves diagonally one step;
	t.Run("tail does not move if head moves one step diagonally", func(t *testing.T) {
		game := MakeRope(Point{1, 0}, 2)
		game.MoveOnce("U")

		AssertPosition(t, game.Head, 1, 1)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("tail does not move if head moves one step diagonally", func(t *testing.T) {
		game := MakeRope(Point{1, 0}, 2)
		game.MoveOnce("D")

		AssertPosition(t, game.Head, 1, -1)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("tail does not move if head moves one step diagonally", func(t *testing.T) {
		game := MakeRope(Point{0, 1}, 2)
		game.Right(1)

		AssertPosition(t, game.Head, 1, 1)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("tail does not move if head moves one step diagonally", func(t *testing.T) {
		game := MakeRope(Point{0, 1}, 2)
		game.Left(1)

		AssertPosition(t, game.Head, -1, 1)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("will run example instructions", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		game := MakeRope(Point{0, 0}, 2)

		for _, row := range input {
			game.Move(row)
		}

		AssertPosition(t, game.Head, 2, 2)
		AssertPosition(t, game.Tail, 1, 2)
	})

	t.Run("part 1 example case", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		got := Part1(input)

		want := 13
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
