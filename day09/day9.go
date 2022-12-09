package day09

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	PosX int
	PosY int
}
type Rope struct {
	Head   Point
	Tail   Point
	Visits map[string]int
}

func (r *Rope) Up(dist int) {
	r.Head.PosY += dist
	diffX := math.Abs(float64(r.Tail.PosX - r.Head.PosX))
	diffY := math.Abs(float64(r.Tail.PosY - r.Head.PosY))

	if diffY > 1 || diffX > 1 {
		r.Tail.PosX = r.Head.PosX
		r.Tail.PosY = r.Head.PosY - 1

		key := fmt.Sprintf("%d:%d", r.Tail.PosY, r.Tail.PosX)
		if nil == r.Visits {
			r.Visits = map[string]int{}
		}
		r.Visits[key] = 1
	}

}

func (r *Rope) Down(dist int) {
	r.Head.PosY -= dist
	diffX := math.Abs(float64(r.Tail.PosX - r.Head.PosX))
	diffY := math.Abs(float64(r.Tail.PosY - r.Head.PosY))

	if diffY > 1 || diffX > 1 {
		r.Tail.PosX = r.Head.PosX
		r.Tail.PosY = r.Head.PosY + 1

		key := fmt.Sprintf("%d:%d", r.Tail.PosY, r.Tail.PosX)
		if nil == r.Visits {
			r.Visits = map[string]int{}
		}
		r.Visits[key] = 1
	}
}

func (r *Rope) Right(dist int) {
	r.Head.PosX += dist

	diffX := math.Abs(float64(r.Tail.PosX - r.Head.PosX))
	diffY := math.Abs(float64(r.Tail.PosY - r.Head.PosY))

	if diffY > 1 || diffX > 1 {
		r.Tail.PosY = r.Head.PosY
		r.Tail.PosX = r.Head.PosX - 1

		key := fmt.Sprintf("%d:%d", r.Tail.PosY, r.Tail.PosX)
		if nil == r.Visits {
			r.Visits = map[string]int{}
		}
		r.Visits[key] = 1
	}
}

func (r *Rope) Left(dist int) {
	r.Head.PosX -= dist

	diffX := math.Abs(float64(r.Tail.PosX - r.Head.PosX))
	diffY := math.Abs(float64(r.Tail.PosY - r.Head.PosY))

	if diffY > 1 || diffX > 1 {
		r.Tail.PosY = r.Head.PosY
		r.Tail.PosX = r.Head.PosX + 1
		key := fmt.Sprintf("%d:%d", r.Tail.PosY, r.Tail.PosX)
		if nil == r.Visits {
			r.Visits = map[string]int{}
		}
		r.Visits[key] = 1
	}
}

func (r *Rope) Move(row string) {
	parts := strings.Split(row, " ")
	direction := parts[0]
	dist, _ := strconv.Atoi(parts[1])

	for i := dist; i > 0; i-- {
		switch direction {
		case "R":
			r.Right(1)
		case "L":
			r.Left(1)
		case "U":
			r.Up(1)
		case "D":
			r.Down(1)
		}

		fmt.Println(row, r.Tail)
	}
}

func Part1(input []string) int {
	game := Rope{
		Visits: map[string]int{
			"0:0": 1,
		},
	}

	for _, row := range input {
		game.Move(row)
	}

	return len(game.Visits)
}
