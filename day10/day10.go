package day10

import "strings"

type Tube struct {
	Cycle         int
	RegisterValue int
}

func (t *Tube) RunInstruction(instr string) {
	t.Cycle++

	split := strings.Split(instr, " ")
	switch split[0] {
	case "noop":
		break // do nothing
	case "addx":
		t.Cycle++
	default:
		panic("unrecognised instructions")
	}
}

func CreateTube() Tube {
	return Tube{
		RegisterValue: 1,
	}
}
