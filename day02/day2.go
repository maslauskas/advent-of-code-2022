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

func PredictEndings(lines []string) int {
	score := 0
	for _, line := range lines {
		res := strings.Split(line, " ")

		riggedGame := RiggedGame{Player1: Sign{res[0]}, Outcome: res[1]}
		game := Game{riggedGame.Player1, Sign{riggedGame.player2Sign()}}

		score += game.score() + game.Player2.score()
	}

	return score
}

type Sign struct {
	Input string
}

func (sign Sign) mappedSign() string {
	if sign.Input == "A" || sign.Input == "X" || sign.Input == "ROCK" {
		return "ROCK"
	}

	if sign.Input == "B" || sign.Input == "Y" || sign.Input == "PAPER" {
		return "PAPER"
	}

	if sign.Input == "C" || sign.Input == "Z" || sign.Input == "SCISSORS" {
		return "SCISSORS"
	}

	panic("bad input")
}

func (sign Sign) score() int {
	if sign.mappedSign() == "ROCK" {
		return 1
	}

	if sign.mappedSign() == "PAPER" {
		return 2
	}

	if sign.mappedSign() == "SCISSORS" {
		return 3
	}

	panic("bad input")
}

type Game struct {
	Player1 Sign
	Player2 Sign
}

type RiggedGame struct {
	Player1 Sign
	Outcome string
}

func (game RiggedGame) player2Sign() string {
	// if need to draw, return same
	if game.Outcome == "Y" {
		return game.Player1.mappedSign()
	}
	// if need to win, return stronger
	if game.Outcome == "Z" {
		if game.Player1.mappedSign() == "ROCK" {
			return "PAPER"
		}

		if game.Player1.mappedSign() == "PAPER" {
			return "SCISSORS"
		}

		if game.Player1.mappedSign() == "SCISSORS" {
			return "ROCK"
		}

		panic("bad sign")
	}

	// if need to lose, return weaker
	if game.Outcome == "X" {
		if game.Player1.mappedSign() == "ROCK" {
			return "SCISSORS"
		}

		if game.Player1.mappedSign() == "PAPER" {
			return "ROCK"
		}

		if game.Player1.mappedSign() == "SCISSORS" {
			return "PAPER"
		}

		panic("bad sign")
	}

	panic("no outcome")
}

func (game Game) isDraw() bool {
	return game.Player1.mappedSign() == game.Player2.mappedSign()
}

func (game Game) isWin() bool {
	if game.Player2.mappedSign() == "ROCK" {
		return game.Player1.mappedSign() == "SCISSORS"
	}

	if game.Player2.mappedSign() == "PAPER" {
		return game.Player1.mappedSign() == "ROCK"
	}

	if game.Player2.mappedSign() == "SCISSORS" {
		return game.Player1.mappedSign() == "PAPER"
	}

	panic("bad sign")
}

func (game Game) score() int {
	if game.isDraw() {
		return 3
	}

	if game.isWin() {
		return 6
	}

	// loss
	return 0
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

	outcomes := map[string]Choice{
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
	LOSS = -1
	DRAW = 0
	WIN  = 1
)
