package day09

import "testing"

func TestExample(t *testing.T) {
	t.Run("will not move tail if head is on top", func(t *testing.T) {
		game := Rope{Point{0, 0}, Point{0, 0}}
		game.Up()

		if game.Tail.PosX != 0 || game.Tail.PosY != 0 {
			t.Errorf("expected possition to be 0,0, got %d,%d", game.Tail.PosX, game.Tail.PosY)
		}
	})
}
