package day02

import (
	"strings"
)

func Part1(input []string) int {
	score := 0
	for _, s := range input {
		game := CreateGame(s)
		score += game.Score()
	}

	return score
}

func Part2(input []string) int {
	score := 0
	for _, s := range input {
		game := CreateRiggedGame(s)
		score += game.Score()
	}

	return score
}

type NewGame struct {
	Player1 Choice
	Player2 Choice
}

func (g NewGame) IsDraw() bool {
	return g.Player1 == g.Player2
}

func (g NewGame) IsWon() bool {
	if g.Player1 == ROCK {
		return g.Player2 == PAPER
	}

	if g.Player1 == PAPER {
		return g.Player2 == SCISSORS
	}

	if g.Player1 == SCISSORS {
		return g.Player2 == ROCK
	}

	return false
}

func (g NewGame) Score() int {
	score := int(g.Player2)

	if g.IsDraw() {
		score += 3
	}

	if g.IsWon() {
		score += 6
	}

	return score
}

type Choice int

const (
	ROCK     Choice = 1
	PAPER           = 2
	SCISSORS        = 3
)

func CreateGame(s string) NewGame {
	mapPlayer1 := map[string]Choice{
		"A": ROCK,
		"B": PAPER,
		"C": SCISSORS,
	}

	mapPlayer2 := map[string]Choice{
		"X": ROCK,
		"Y": PAPER,
		"Z": SCISSORS,
	}

	parts := strings.Split(s, " ")

	return NewGame{mapPlayer1[parts[0]], mapPlayer2[parts[1]]}
}

func CreateRiggedGame(s string) NewGame {
	parts := strings.Split(s, " ")

	mapPlayer1 := map[string]Choice{
		"A": ROCK,
		"B": PAPER,
		"C": SCISSORS,
	}

	p1 := mapPlayer1[parts[0]]

	outcomes := map[string]Outcome{
		"X": LOSS,
		"Y": DRAW,
		"Z": WIN,
	}

	outcome := outcomes[parts[1]]

	var p2 Choice

	switch outcome {
	case DRAW:
		p2 = p1
	case WIN:
		switch p1 {
		case ROCK:
			p2 = PAPER
		case PAPER:
			p2 = SCISSORS
		case SCISSORS:
			p2 = ROCK
		}
	case LOSS:
		switch p1 {
		case ROCK:
			p2 = SCISSORS
		case PAPER:
			p2 = ROCK
		case SCISSORS:
			p2 = PAPER
		}
	}

	return NewGame{p1, p2}
}

type Outcome int

const (
	LOSS Outcome = -1
	DRAW         = 0
	WIN          = 1
)
