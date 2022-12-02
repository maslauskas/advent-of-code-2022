package day02

import (
	"adventofcode/helpers"
	"fmt"
	"strings"
)

func Part1(path string) {
	total := CalculateTotalScore(path)

	fmt.Println(total)
}

func Part2(path string) {
	total := PredictEndings(path)

	fmt.Println(total)
}

func CalculateTotalScore(path string) int {
	scores := CalculateRowsScore(path)

	total := 0

	for _, score := range scores {
		total += score
	}

	return total
}

func CalculateRowsScore(path string) []int {
	lines := helpers.ReadInput(path)

	var scores []int
	for _, line := range lines {
		res := strings.Split(line, " ")
		score := GetScoreFromLine(res)

		scores = append(scores, score)
	}

	return scores
}

func PredictEndings(path string) int {
	lines := helpers.ReadInput(path)

	score := 0
	for _, line := range lines {
		res := strings.Split(line, " ")

		riggedGame := RiggedGame{Player1: Sign{res[0]}, Outcome: res[1]}
		game := Game{riggedGame.Player1, Sign{riggedGame.player2Sign()}}

		score += game.score() + game.Player2.score()
	}

	return score
}

func GetScoreFromLine(res []string) int {
	player1 := Sign{res[0]}
	player2 := Sign{res[1]}

	game := Game{Player1: player1, Player2: player2}

	return game.score() + player2.score()
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
