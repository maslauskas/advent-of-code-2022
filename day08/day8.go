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

func (b Board) CheckBlocked(x int, y int) bool {
	top := CheckTop(b, x, y)
	right := CheckRight(b, x, y)
	bottom := CheckBottom(b, x, y)
	left := CheckLeft(b, x, y)

	return top && right && bottom && left
}

func CheckRight(b Board, x int, y int) bool {
	val := b[y][x]

	for i := x + 1; i < len(b[y]); i++ {
		neighbor := b[y][i]
		if neighbor > val {
			return true
		}
	}

	return false
}

func CheckLeft(b Board, x int, y int) bool {
	val := b[y][x]

	for i := x - 1; i >= 0; i-- {
		neighbor := b[y][i]
		if neighbor > val {
			return true
		}
	}

	return false
}

func CheckTop(b Board, x int, y int) bool {
	val := b[y][x]

	for i := y - 1; i >= 0; i-- {
		neighbor := b[i][x]
		if neighbor > val {
			return true
		}
	}

	return false
}

func CheckBottom(b Board, x int, y int) bool {
	val := b[y][x]

	for i := y + 1; i < len(b); i++ {
		neighbor := b[i][x]
		if neighbor > val {
			return true
		}
	}

	return false
}
