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
		game.Up(2)

		AssertPosition(t, game.Head, 0, 2)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("will move head down", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Down(3)

		AssertPosition(t, game.Head, 0, -3)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("will move head right", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Right(4)

		AssertPosition(t, game.Head, 4, 0)
		AssertPosition(t, game.Tail, 0, 0)
	})

	t.Run("will move head left", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Left(5)

		AssertPosition(t, game.Head, -5, 0)
		AssertPosition(t, game.Tail, 0, 0)
	})
}

func AssertPosition(t *testing.T, p Point, x int, y int) {
	t.Helper()

	if p.PosX != x || p.PosY != y {
		t.Errorf("expected possition to be %v, got {%d %d}", p, x, y)
	}
}
