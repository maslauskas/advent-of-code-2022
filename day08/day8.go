package day08

import (
	"strconv"
)

func BuildBoard(input []string) Board {
	board := make(Board, len(input))
	for y, row := range input {
		board[y] = make([]int, 2)
		for i := 0; i < len(row); i++ {
			val, _ := strconv.Atoi(string(row[i]))
			board[y][i] = val
		}

	}

	return board
}

type Board [][]int
