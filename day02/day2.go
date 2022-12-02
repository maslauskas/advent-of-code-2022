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

type Game struct {
	Player1 Choice
	Player2 Choice
}

func (g Game) IsDraw() bool {
	return g.Player1 == g.Player2
}

func (g Game) IsWon() bool {
	switch g.Player1 {
	case ROCK:
		return g.Player2 == PAPER
	case PAPER:
		return g.Player2 == SCISSORS
	case SCISSORS:
		return g.Player2 == ROCK
	}

	return false
}

func (g Game) Score() int {
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

type ChoiceMap map[string]Choice

func CreateGame(s string) Game {
	parts := strings.Split(s, " ")
	p1 := GetMappedChoice(parts[0])
	p2 := GetMappedChoice(parts[1])

	return Game{p1, p2}
}

func GetMappedChoice(s string) Choice {
	mapped := ChoiceMap{
		"A": ROCK,
		"B": PAPER,
		"C": SCISSORS,
		"X": ROCK,
		"Y": PAPER,
		"Z": SCISSORS,
	}

	return mapped[s]
}

func CreateRiggedGame(s string) Game {
	parts := strings.Split(s, " ")

	p1 := GetMappedChoice(parts[0])

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
		winMap := map[Choice]Choice{
			ROCK:     PAPER,
			PAPER:    SCISSORS,
			SCISSORS: ROCK,
		}

		p2 = winMap[p1]
	case LOSS:
		lossMap := map[Choice]Choice{
			ROCK:     SCISSORS,
			PAPER:    ROCK,
			SCISSORS: PAPER,
		}

		p2 = lossMap[p1]
	}

	return Game{p1, p2}
}

type Outcome int

const (
	LOSS Outcome = -1
	DRAW         = 0
	WIN          = 1
)
