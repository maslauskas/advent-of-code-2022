package day11

import (
	"strconv"
	"strings"
	"testing"
)

type Monkey struct {
}

func (m Monkey) Investigate(item int, operation string) int {
	parts := strings.Split(operation, " ")
	op := parts[0]
	value, _ := strconv.Atoi(parts[1])

	switch op {
	case "+":
		return item + value
	case "*":
		return item * value
	default:
		panic("unrecognised operation")
	}

}

func TestMonkeyBusiness(t *testing.T) {
	t.Run("will investigate item", func(t *testing.T) {
		monkey := Monkey{}

		dataset := map[string]int{
			"+ 1": 2,
			"+ 2": 3,
			"* 1": 1,
			"* 5": 5,
		}

		for operation, want := range dataset {
			got := monkey.Investigate(1, operation)

			if want != got {
				t.Errorf("expected item value to be %d, got %d", want, got)
			}
		}

	})
}
