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
	Head     Point
	Tail     Point
	Segments []Point
	Visits   map[string]int
}

func MakeRope(head Point, tailSegments int) Rope {
	segments := make([]Point, tailSegments)
	segments[0] = head
	return Rope{
		// support old calculation
		Head: head,
		Tail: segments[len(segments)-1],
		// new property
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
	switch direction {
	case "R":
		r.Head.PosX += 1
		moveX := -1
		moveY := 0

		r.MoveSegment(moveY, moveX, &r.Tail, r.Head)
		r.MoveSegment(moveY, moveX, &r.Segments[len(r.Segments)-1], r.Head)
	case "L":
		r.Head.PosX -= 1
		moveX := 1
		moveY := 0

		r.MoveSegment(moveY, moveX, &r.Tail, r.Head)
		r.MoveSegment(moveY, moveX, &r.Segments[len(r.Segments)-1], r.Head)
	case "U":
		r.Head.PosY += 1
		moveX := 0
		moveY := -1

		r.MoveSegment(moveY, moveX, &r.Tail, r.Head)
		r.MoveSegment(moveY, moveX, &r.Segments[len(r.Segments)-1], r.Head)
	case "D":
		r.Head.PosY -= 1
		moveX := 0
		moveY := 1

		r.MoveSegment(moveY, moveX, &r.Tail, r.Head)
		r.MoveSegment(moveY, moveX, &r.Segments[len(r.Segments)-1], r.Head)
	}
}

func (r *Rope) MoveSegment(moveY int, moveX int, tail *Point, head Point) {
	diffX := math.Abs(float64(tail.PosX - head.PosX))
	diffY := math.Abs(float64(tail.PosY - head.PosY))

	if diffY > 1 || diffX > 1 {
		tail.PosY = head.PosY + moveY
		tail.PosX = head.PosX + moveX

		r.LogTailVisit()
	}
}

func (r *Rope) LogTailVisit() {
	key := fmt.Sprintf("%d:%d", r.GetTail().PosY, r.GetTail().PosX)
	if nil == r.Visits {
		r.Visits = map[string]int{}
	}
	r.Visits[key] = 1
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
