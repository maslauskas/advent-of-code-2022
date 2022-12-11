package day11

import (
	"regexp"
	"strconv"
	"strings"
)

type Monkey struct {
	Items     []int
	Operation string
	Test      int
	Target1   *Monkey
	Target2   *Monkey
}

func (m Monkey) Investigate(item int, operation string) int {
	parts := strings.Split(operation, " ")
	op := parts[0]
	value := parts[1]
	var val int

	if value == "old" {
		val = item
	} else {
		v, _ := strconv.Atoi(parts[1])
		val = v
	}

	switch op {
	case "+":
		return item + val
	case "*":
		return item * val
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

func (m *Monkey) ProcessAllItems() {
	for len(m.Items) > 0 {
		item := m.Items[0]
		m.Items = m.Items[1:]

		m.ProcessItem(item, m.Target1, m.Target2)
	}
}

func CreateMonkeySquad(input []string) []Monkey {
	count := 1

	// count monkeys
	for _, line := range input {
		if line == "" {
			count++
		}
	}

	// create monkey instruction input
	monkeyInput := make([][]string, count)
	index := 0
	for _, line := range input {
		if line == "" {
			index++
			continue
		}
		monkeyInput[index] = append(monkeyInput[index], line)
	}

	// create correct monkey squad size
	var squad []Monkey

	for _, m := range monkeyInput {
		monkey := Monkey{}
		r, _ := regexp.Compile(`\d+`)
		items := r.FindAllString(m[1], -1)
		for _, item := range items {
			val, _ := strconv.Atoi(item)
			monkey.Items = append(monkey.Items, val)
		}
		squad = append(squad, monkey)
	}

	return squad
}
