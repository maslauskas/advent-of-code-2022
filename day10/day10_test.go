package day10

import "testing"

func TestProgram(t *testing.T) {
	t.Run("will run instructions", func(t *testing.T) {
		dataset := map[string][2]int{
			"noop":   {1, 1},
			"addx 1": {2, 1},
		}

		for instr, wants := range dataset {
			tube := CreateTube()
			tube.RunInstruction(instr)

			AssertTubeCycle(t, tube, wants[0], instr)
			AssertRegisterValue(t, tube, wants[1], instr)
		}

	})
}

func AssertRegisterValue(t *testing.T, tube Tube, want int, set string) {
	t.Helper()

	got := tube.RegisterValue

	if want != got {
		t.Errorf("%q: expected cycle to be %d, got %d", set, want, got)
	}
}

func AssertTubeCycle(t *testing.T, tube Tube, want int, set string) {
	t.Helper()

	got := tube.Cycle

	if want != got {
		t.Errorf("%q: expected cycle to be %d, got %d", set, want, got)
	}
}
