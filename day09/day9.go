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

func (tail *Point) CatchUp(head Point) {
	deltaY := math.Abs(float64(tail.PosY - head.PosY))
	deltaX := math.Abs(float64(tail.PosX - head.PosX))

	if deltaY == 2 {
		if deltaX == 1 {
			tail.PosX = head.PosX
		}
		if head.PosY > tail.PosY {
			tail.PosY++
		} else {
			tail.PosY--
		}
	}

	if deltaX == 2 {
		if deltaY == 1 {
			tail.PosY = head.PosY
		}
		if head.PosX > tail.PosX {
			tail.PosX++
		} else {
			tail.PosX--
		}
	}
}

type Rope struct {
	Segments []Point
	Visits   map[string]int
}

func MakeRope(head Point, tailSegments int) Rope {
	segments := make([]Point, tailSegments)
	segments[0] = head
	return Rope{
		Segments: segments,
	}
}

func (r *Rope) Move(row string) {
	parts := strings.Split(row, " ")
	direction := parts[0]
	dist, _ := strconv.Atoi(parts[1])

	for i := dist; i > 0; i-- {
		r.MoveOnce(direction)
	}
}

func (r *Rope) MoveOnce(direction string) {
	var head *Point
	head = &r.Segments[0]

	switch direction {
	case "R":
		head.PosX += 1
	case "L":
		head.PosX -= 1
	case "U":
		head.PosY += 1
	case "D":
		head.PosY -= 1
	}

	for i := 1; i < len(r.Segments); i++ {
		tail := &r.Segments[i]
		tail.CatchUp(r.Segments[i-1])
	}

	r.LogTailVisit()

}

func (r *Rope) LogTailVisit() {
	key := fmt.Sprintf("%d:%d", r.GetTail().PosY, r.GetTail().PosX)
	if nil == r.Visits {
		r.Visits = map[string]int{}
	}
	r.Visits[key] = 1
}

func (r *Rope) GetHead() Point {
	return r.Segments[0]
}

func (r *Rope) GetTail() Point {
	return r.Segments[len(r.Segments)-1]
}

func Part1(input []string) int {
	game := MakeRope(Point{0, 0}, 2)
	game.Visits = map[string]int{
		"0:0": 1,
	}

	for _, row := range input {
		game.Move(row)
	}

	return len(game.Visits)
}
