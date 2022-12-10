package day10

import (
	"math"
	"strconv"
	"strings"
)

type Tube struct {
	Cycle          int
	RegisterValue  int
	SignalStrength int
	Crt            [6][40]string
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
	x := t.Cycle % 40
	y := math.Floor(float64(t.Cycle) / 40.0)

	if x-1 == t.RegisterValue || x == t.RegisterValue || x+1 == t.RegisterValue {
		t.Crt[int(y)][x] = "#"
	} else {
		t.Crt[int(y)][x] = "."
	}

	t.Cycle++
	if t.Cycle%40 == 20 {
		t.SignalStrength += t.RegisterValue * t.Cycle
	}
}

func (t *Tube) GetPixel(x int, y int) string {
	return t.Crt[y][x]
}

func (t *Tube) Display() string {
	output := ""
	for y, row := range t.Crt {
		for _, val := range row {
			output += val
		}

		if y < len(t.Crt)-1 {
			output += "\n"
		}
	}

	return output
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

func Part2(input []string) string {
	tube := MakeTube()

	for _, instr := range input {
		tube.RunInstruction(instr)
	}

	return tube.Display()
}
