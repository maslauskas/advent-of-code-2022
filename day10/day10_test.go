package day10

import (
	"adventofcode/helpers"
	"testing"
)

func TestTubeInstructions(t *testing.T) {
	t.Run("will run instructions", func(t *testing.T) {
		dataset := map[string][2]int{
			"noop":    {1, 1},
			"addx 1":  {2, 2},
			"addx -5": {2, -4},
		}

		for instr, wants := range dataset {
			tube := MakeTube()
			tube.RunInstruction(instr)

			AssertTubeCycle(t, tube, wants[0], instr)
			AssertRegisterValue(t, tube, wants[1], instr)
		}
	})
}

func TestSignalStrength(t *testing.T) {
	t.Run("initial strength will be 0", func(t *testing.T) {
		tube := MakeTube()
		AssertSignalStrength(t, tube, 0)
	})

	t.Run("will not increase when executing instructions", func(t *testing.T) {
		tube := MakeTube()
		tube.RunInstruction("addx 5")
		AssertSignalStrength(t, tube, 0)
	})

	t.Run("will increase when cycle reaches 20", func(t *testing.T) {
		tube := MakeTube()
		tube.Cycle = 17

		tube.RunInstruction("addx 4")
		tube.RunInstruction("noop")
		AssertSignalStrength(t, tube, 100)
	})

	t.Run("will increase when cycle reaches 20 in the middle of instruction", func(t *testing.T) {
		tube := MakeTube()
		tube.Cycle = 18

		tube.RunInstruction("addx 4")
		AssertSignalStrength(t, tube, 20)
	})

	t.Run("will increase when cycle reaches 60", func(t *testing.T) {
		tube := MakeTube()
		tube.Cycle = 57

		tube.RunInstruction("addx 4")
		tube.RunInstruction("noop")

		AssertSignalStrength(t, tube, 300)
	})

	t.Run("part 1 example case", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		got := Part1(input)
		want := 13140

		if want != got {
			t.Errorf("expected result to be %d, got %d", want, got)
		}
	})
}

func AssertSignalStrength(t *testing.T, tube Tube, want int) {
	t.Helper()
	got := tube.SignalStrength

	if want != got {
		t.Errorf("expected signal strength of %d, got %d", want, got)
	}
}

func AssertRegisterValue(t *testing.T, tube Tube, want int, set string) {
	t.Helper()

	got := tube.RegisterValue

	if want != got {
		t.Errorf("%q: expected register value to be %d, got %d", set, want, got)
	}
}

func AssertTubeCycle(t *testing.T, tube Tube, want int, set string) {
	t.Helper()

	got := tube.Cycle

	if want != got {
		t.Errorf("%q: expected cycle to be %d, got %d", set, want, got)
	}
}
