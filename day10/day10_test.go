package day10

import "testing"

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
