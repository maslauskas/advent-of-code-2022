package day10

import (
	"adventofcode/helpers"
	"fmt"
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

func TestCrt(t *testing.T) {
	t.Run("will draw pixel at the start of first cycle", func(t *testing.T) {
		tube := MakeTube()
		tube.IncrementCycle()

		AssertPixelVisible(t, tube, 0, 0)
	})

	t.Run("will draw pixels after first instruction", func(t *testing.T) {
		tube := MakeTube()
		tube.RunInstruction("addx 15")

		AssertPixelVisible(t, tube, 0, 0)
		AssertPixelVisible(t, tube, 1, 0)
	})

	t.Run("will draw pixels after first two instructions", func(t *testing.T) {
		tube := MakeTube()
		tube.RunInstruction("addx 15")
		tube.RunInstruction("addx -11")

		AssertPixelVisible(t, tube, 0, 0)
		AssertPixelVisible(t, tube, 1, 0)
		AssertPixelNotVisible(t, tube, 2, 0)
		AssertPixelNotVisible(t, tube, 3, 0)
	})

	t.Run("will draw pixels after first two instructions on the second row", func(t *testing.T) {
		tube := MakeTube()
		tube.Cycle = 40
		tube.RunInstruction("addx 15")
		tube.RunInstruction("addx -11")

		AssertPixelVisible(t, tube, 0, 1)
		AssertPixelVisible(t, tube, 1, 1)
		AssertPixelNotVisible(t, tube, 2, 1)
		AssertPixelNotVisible(t, tube, 3, 1)
	})

	t.Run("part 2 example case", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		got := Part2(input)
		want := `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`

		if want != got {
			t.Errorf("expected output not matched")
			fmt.Println(got)
			fmt.Println("===")
			fmt.Println(want)
		}
	})
}

func AssertPixelVisible(t *testing.T, tube Tube, x int, y int) {
	pixel := tube.GetPixel(x, y)

	if pixel != "#" {
		t.Errorf("expected pixel to be visible at {%d %d}", x, y)
	}
}

func AssertPixelNotVisible(t *testing.T, tube Tube, x int, y int) {
	pixel := tube.GetPixel(x, y)

	if pixel != "." {
		t.Errorf("expected pixel to be not visible at {%d %d}", x, y)
	}
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
