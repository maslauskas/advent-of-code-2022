package day11

import (
	"strconv"
	"strings"
)

type Monkey struct {
	Items     []int
	Operation string
	Test      int
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

func (m Monkey) ProcessItem(item int, t1 *Monkey, t2 *Monkey) {
	item = m.Investigate(item, m.Operation)
	item = m.GetBored(item)

	m.Throw(item, t1, t2, m.Test)
}
