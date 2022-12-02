package day02

import (
	"adventofcode/helpers"
	"testing"
)

func TestRockPaperScissors(t *testing.T) {
	input := helpers.ReadInput("./example.txt")

	t.Run("part 1 example case", func(t *testing.T) {
		got := Part1(input)
		want := 15

		if want != got {
			t.Errorf("expected total score %v, got %v", want, got)
		}
	})

	t.Run("part 2 example case", func(t *testing.T) {
		got := Part2(input)
		want := 12

		if want != got {
			t.Errorf("expected total score %v, got %v", want, got)
		}
	})

	t.Run("game is draw", func(t *testing.T) {
		game := Game{ROCK, ROCK}

		if !game.IsDraw() {
			t.Error("expected game to be draw")
		}
	})

	t.Run("victory conditions", func(t *testing.T) {
		tests := map[string][]Choice{
			"rock vs paper":     {ROCK, PAPER},
			"paper vs scissors": {PAPER, SCISSORS},
			"scissors vs rock":  {SCISSORS, ROCK},
		}

		for name, tt := range tests {
			game := Game{tt[0], tt[1]}
			if !game.IsWon() {
				t.Errorf("%s: expected game to be won", name)
			}
		}
	})

	t.Run("converts to Game", func(t *testing.T) {
		tests := map[string]Game{
			"A Y": {ROCK, PAPER},
			"B X": {PAPER, ROCK},
			"C Z": {SCISSORS, SCISSORS},
		}

		for s, want := range tests {
			game := CreateGame(s)
			if game != want {
				t.Errorf("expected game %v, got %v", want, game)
			}
		}
	})

	t.Run("converts to RiggedGame", func(t *testing.T) {
		tests := map[string]Game{
			"A Y": {ROCK, ROCK},
			"B X": {PAPER, ROCK},
			"C Z": {SCISSORS, ROCK},
		}

		for s, want := range tests {
			game := CreateRiggedGame(s)
			if game != want {
				t.Errorf("%s: expected game %v, got %v", s, want, game)
			}
		}
	})

	t.Run("calculates game score", func(t *testing.T) {
		tests := map[string]int{
			"A Y": 8,
			"B X": 1,
			"C Z": 6,
		}

		for s, want := range tests {
			game := CreateGame(s)
			if game.Score() != want {
				t.Errorf("%s: expected score of %d, got %d", s, want, game.Score())
			}
		}
	})
}
