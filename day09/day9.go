package day09

import (
	"math"
	"strconv"
	"strings"
)

type Point struct {
	PosX int
	PosY int
}
type Rope struct {
	Head Point
	Tail Point
}

func (r *Rope) Up(dist int) {
	r.Head.PosY += dist
	diffX := math.Abs(float64(r.Tail.PosX - r.Head.PosX))
	diffY := math.Abs(float64(r.Tail.PosY - r.Head.PosY))

	if diffY+diffX > 2 {
		r.Tail.PosX = r.Head.PosX
		r.Tail.PosY = r.Head.PosY - 1
	}

}

func (r *Rope) Down(dist int) {
	r.Head.PosY -= dist
	diffX := math.Abs(float64(r.Tail.PosX - r.Head.PosX))
	diffY := math.Abs(float64(r.Tail.PosY - r.Head.PosY))

	if diffY+diffX > 2 {
		r.Tail.PosX = r.Head.PosX
		r.Tail.PosY = r.Head.PosY + 1
	}
}

func (r *Rope) Right(dist int) {
	r.Head.PosX += dist

	diffX := math.Abs(float64(r.Tail.PosX - r.Head.PosX))
	diffY := math.Abs(float64(r.Tail.PosY - r.Head.PosY))

	if diffY+diffX > 2 {
		r.Tail.PosY = r.Head.PosY
		r.Tail.PosX = r.Head.PosX - 1
	}
}

func (r *Rope) Left(dist int) {
	r.Head.PosX -= dist

	diffX := math.Abs(float64(r.Tail.PosX - r.Head.PosX))
	diffY := math.Abs(float64(r.Tail.PosY - r.Head.PosY))

	if diffY+diffX > 2 {
		r.Tail.PosY = r.Head.PosY
		r.Tail.PosX = r.Head.PosX + 1
	}
}

func (r *Rope) Move(row string) {
	parts := strings.Split(row, " ")
	direction := parts[0]
	dist, _ := strconv.Atoi(parts[1])

	switch direction {
	case "R":
		r.Right(dist)
	case "L":
		r.Left(dist)
	case "U":
		r.Up(dist)
	case "D":
		r.Down(dist)
	}
}
