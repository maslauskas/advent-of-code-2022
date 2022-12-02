package day01

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 24000
	got := MostCaloriesPerElf("./example.txt")

	if want != got {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestPart2(t *testing.T) {
	want := 45000
	got := TopThreeCalories("./example.txt")

	if want != got {
		t.Errorf("expected %v, got %v", want, got)
	}
}
