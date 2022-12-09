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

func (r *Rope) Up(dist int) {
	r.Move(fmt.Sprintf("U %d", dist))
}

func (r *Rope) Down(dist int) {
	r.Move(fmt.Sprintf("D %d", dist))
}

func (r *Rope) Right(dist int) {
	r.Move(fmt.Sprintf("R %d", dist))
}

func (r *Rope) Left(dist int) {
	r.Move(fmt.Sprintf("L %d", dist))
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

		diffX := math.Abs(float64(r.Tail.PosX - r.Head.PosX))
		diffY := math.Abs(float64(r.Tail.PosY - r.Head.PosY))

		if diffY > 1 || diffX > 1 {
			r.Tail.PosY = r.Head.PosY
			r.Tail.PosX = r.Head.PosX - 1

			r.LogTailVisit()
		}
	case "L":
		r.Head.PosX -= 1

		diffX := math.Abs(float64(r.Tail.PosX - r.Head.PosX))
		diffY := math.Abs(float64(r.Tail.PosY - r.Head.PosY))

		if diffY > 1 || diffX > 1 {
			r.Tail.PosY = r.Head.PosY
			r.Tail.PosX = r.Head.PosX + 1

			r.LogTailVisit()
		}
	case "U":
		r.Head.PosY += 1
		diffX := math.Abs(float64(r.Tail.PosX - r.Head.PosX))
		diffY := math.Abs(float64(r.Tail.PosY - r.Head.PosY))

		if diffY > 1 || diffX > 1 {
			r.Tail.PosX = r.Head.PosX
			r.Tail.PosY = r.Head.PosY - 1

			r.LogTailVisit()
		}
	case "D":
		r.Head.PosY -= 1
		diffX := math.Abs(float64(r.Tail.PosX - r.Head.PosX))
		diffY := math.Abs(float64(r.Tail.PosY - r.Head.PosY))

		if diffY > 1 || diffX > 1 {
			r.Tail.PosX = r.Head.PosX
			r.Tail.PosY = r.Head.PosY + 1

			r.LogTailVisit()
		}
	}
}

func (r *Rope) LogTailVisit() {
	key := fmt.Sprintf("%d:%d", r.Tail.PosY, r.Tail.PosX)
	if nil == r.Visits {
		r.Visits = map[string]int{}
	}
	r.Visits[key] = 1
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
