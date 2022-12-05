package day05

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(input []string) string {
	c := CreateStackCollection(input)
	c.MoveAll()

	res := ""
	for i := 1; i <= len(c.Stacks); i++ {
		stack := c.Stacks[i]
		lastIndex := len(stack) - 1
		res = res + stack[lastIndex]
	}

	return res
}

func Part2(input []string) string {
	c := CreateStackCollection(input)
	c.MoveAllStacks()

	res := ""
	for i := 1; i <= len(c.Stacks); i++ {
		stack := c.Stacks[i]
		lastIndex := len(stack) - 1
		res = res + stack[lastIndex]
	}

	return res
}

type StackCollection struct {
	Stacks       map[int][]string
	Instructions []string
}

func (c StackCollection) Move(step int) {
	for i := 0; i < step; i++ {
		count, from, to := ParseInstructions(c.Instructions[i])

		c.ExecuteInstruction(count, from, to)
	}
}

func (c StackCollection) MoveAll() {
	for _, instr := range c.Instructions {
		count, from, to := ParseInstructions(instr)

		c.ExecuteInstruction(count, from, to)
	}
}

func (c StackCollection) MoveAllStacks() {
	for _, instr := range c.Instructions {
		count, from, to := ParseInstructions(instr)

		c.ExecuteInstructionRetainingOrder(count, from, to)
	}
}

func (c StackCollection) ExecuteInstruction(count int, from int, to int) {
	for j := count; j > 0; j-- {
		index := len(c.Stacks[from]) - 1
		box := c.Stacks[from][index]

		c.Stacks[to] = append(c.Stacks[to], box)
		c.Stacks[from] = c.Stacks[from][:index]
	}
}

func (c StackCollection) ExecuteInstructionRetainingOrder(count int, from int, to int) {
	index := len(c.Stacks[from])
	boxes := c.Stacks[from][index-count:]

	c.Stacks[to] = append(c.Stacks[to], boxes...)
	c.Stacks[from] = c.Stacks[from][:index-count]
}

func CreateStackCollection(input []string) StackCollection {
	stacks := map[int][]string{}
	instructions := []string{}

	for _, line := range input {
		// skip empty line
		if len(line) == 0 {
			continue
		}

		// add instruction if starts with "move"
		if strings.HasPrefix(line, "move") {
			instructions = append(instructions, line)
			continue
		}

		// build stacks
		row := 0

		for i := 0; i < len(line); i = i + 4 {
			row++

			box := line[i+1]
			isAlpha := regexp.MustCompile(`[A-Z]`).MatchString
			if string(box) == " " || !isAlpha(string(box)) {
				continue
			}

			item := string(box)
			stacks[row] = append([]string{item}, stacks[row]...)
		}
	}

	return StackCollection{stacks, instructions}
}

func ParseInstructions(instr string) (int, int, int) {
	r, _ := regexp.Compile(`\d+`)
	result := r.FindAllString(instr, -1)

	count, _ := strconv.Atoi(result[0])
	from, _ := strconv.Atoi(result[1])
	to, _ := strconv.Atoi(result[2])

	return count, from, to
}
