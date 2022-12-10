package day10

import (
	"strconv"
	"strings"
)

type Tube struct {
	Cycle          int
	RegisterValue  int
	SignalStrength int
}

func (t *Tube) RunInstruction(instr string) {
	t.IncrementCycle()

	split := strings.Split(instr, " ")
	switch split[0] {
	case "noop":
		break // do nothing
	case "addx":
		t.IncrementCycle()
		val, _ := strconv.Atoi(split[1])
		t.RegisterValue += val
	default:
		panic("unrecognised instructions")
	}
}

func (t *Tube) IncrementCycle() {
	t.Cycle++

	if t.Cycle%40 == 20 {
		t.SignalStrength += t.RegisterValue * t.Cycle
	}
}

func (t *Tube) GetPixel(x int, y int) string {
	return "#"
}

func MakeTube() Tube {
	return Tube{
		RegisterValue: 1,
	}
}

func Part1(input []string) int {
	tube := MakeTube()

	for _, instr := range input {
		tube.RunInstruction(instr)
	}

	return tube.SignalStrength
}
