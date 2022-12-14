package day08

import (
	"strconv"
)

func Part1(input []string) int {
	board := BuildBoard(input)
	total := 0

	for y, row := range board {
		for x := range row {
			if !board.CheckBlocked(x, y) {
				total++
			}
		}
	}

	return total
}

func Part2(input []string) int {
	board := BuildBoard(input)
	max := 0

	for y, row := range board {
		for x := range row {
			score := board.ScenicScore(y, x)
			if score > max {
				max = score
			}
		}
	}

	return max
}

func BuildBoard(input []string) Board {
	board := make(Board, len(input))
	for y, row := range input {
		board[y] = make([]int, len(row))
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

func (b Board) ScenicScore(y int, x int) int {
	top := b.CountTop(x, y)
	right := b.CountRight(x, y)
	bottom := b.CountBottom(x, y)
	left := b.CountLeft(x, y)

	return top * right * bottom * left
}

func (b Board) CountTop(x int, y int) int {
	score := 0

	val := b[y][x]

	for i := y - 1; i >= 0; i-- {
		neighbor := b[i][x]
		score++
		if neighbor >= val {
			break
		}
	}

	return score
}

func (b Board) CountRight(x int, y int) int {
	score := 0

	val := b[y][x]

	for i := x + 1; i < len(b[y]); i++ {
		neighbor := b[y][i]
		score++
		if neighbor >= val {
			break
		}
	}

	return score
}

func (b Board) CountBottom(x int, y int) int {
	score := 0

	val := b[y][x]

	for i := y + 1; i < len(b); i++ {
		neighbor := b[i][x]
		score++
		if neighbor >= val {
			break
		}
	}

	return score
}

func (b Board) CountLeft(x int, y int) int {
	score := 0

	val := b[y][x]

	for i := x - 1; i >= 0; i-- {
		neighbor := b[y][i]
		score++
		if neighbor >= val {
			break
		}
	}

	return score
}

func CheckRight(b Board, x int, y int) bool {
	val := b[y][x]

	for i := x + 1; i < len(b[y]); i++ {
		neighbor := b[y][i]
		if neighbor >= val {
			return true
		}
	}

	return false
}

func CheckLeft(b Board, x int, y int) bool {
	val := b[y][x]

	for i := x - 1; i >= 0; i-- {
		neighbor := b[y][i]
		if neighbor >= val {
			return true
		}
	}

	return false
}

func CheckTop(b Board, x int, y int) bool {
	val := b[y][x]

	for i := y - 1; i >= 0; i-- {
		neighbor := b[i][x]
		if neighbor >= val {
			return true
		}
	}

	return false
}

func CheckBottom(b Board, x int, y int) bool {
	val := b[y][x]

	for i := y + 1; i < len(b); i++ {
		neighbor := b[i][x]
		if neighbor >= val {
			return true
		}
	}

	return false
}
