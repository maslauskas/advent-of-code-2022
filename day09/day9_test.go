package day09

import "testing"

func TestExample(t *testing.T) {
	t.Run("will not move tail if head is on top and head moved by 1", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Up(1)

		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("will move head up", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Up(1)

		AssertPosition(t, game.Head, 0, 1)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("will move head down", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Down(1)

		AssertPosition(t, game.Head, 0, -1)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("will move head right", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Right(1)

		AssertPosition(t, game.Head, 1, 0)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("will move head left", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Left(1)

		AssertPosition(t, game.Head, -1, 0)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("tail will catch up when moved straight up", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Up(5)

		AssertPosition(t, game.Head, 0, 5)
		AssertPosition(t, game.Tail, 0, 4)
	})

	t.Run("tail will catch up when moved straight down", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Down(5)

		AssertPosition(t, game.Head, 0, -5)
		AssertPosition(t, game.Tail, 0, -4)
	})

	t.Run("tail will catch up when moved straight right", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Right(5)

		AssertPosition(t, game.Head, 5, 0)
		AssertPosition(t, game.Tail, 4, 0)
	})

	t.Run("tail will catch up when moved straight right", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Left(5)

		AssertPosition(t, game.Head, -5, 0)
		AssertPosition(t, game.Tail, -4, 0)
	})
}

func AssertPosition(t *testing.T, p Point, x int, y int) {
	t.Helper()

	if p.PosX != x || p.PosY != y {
		t.Errorf("expected possition to be %v, got {%d %d}", p, x, y)
	}
}
