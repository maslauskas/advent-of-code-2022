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
	t.Cycle++

	split := strings.Split(instr, " ")
	switch split[0] {
	case "noop":
		break // do nothing
	case "addx":
		t.Cycle++
		val, _ := strconv.Atoi(split[1])
		t.RegisterValue += val
	default:
		panic("unrecognised instructions")
	}
}

func MakeTube() Tube {
	return Tube{
		RegisterValue: 1,
	}
}
