package day11

import (
	"strconv"
	"strings"
)

type Monkey struct {
	Items []int
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

func (m Monkey) Throw(item int, target1 *Monkey, target2 *Monkey, test int) {
	if item%test == 0 {
		target1.Items = append(target1.Items, item)
	} else {
		target2.Items = append(target2.Items, item)
	}
}
