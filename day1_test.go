package main

import "testing"

func TestPart1(t *testing.T) {
    want := 24000
    got := MostCaloriesPerElf("./examples/day1.txt")

    if want != got {
        t.Errorf("expected %v, got %v", want, got)
    }
}

func TestPart2(t *testing.T) {
    want := 45000
    got := TopThreeCalories("./examples/day1.txt")

    if want != got {
        t.Errorf("expected %v, got %v", want, got)
    }
}