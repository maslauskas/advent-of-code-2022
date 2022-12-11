package day11

import (
	"strconv"
	"strings"
)

type Monkey struct {
}

func (m Monkey) Investigate(item int, operation string) int {
	parts := strings.Split(operation, " ")
	op := parts[0]
	value, _ := strconv.Atoi(parts[1])

	switch op {
	case "+":
		return item + value
	case "*":
		return item * value
	default:
		panic("unrecognised operation")
	}

}

func (m Monkey) GetBored(value int) int {
	result := value / 3
	return result
}
