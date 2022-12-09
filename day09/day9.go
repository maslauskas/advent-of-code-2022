package day09

import "math"

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
	r.Tail.PosX = r.Head.PosX
	if math.Abs(float64(r.Tail.PosY-r.Head.PosY)) > 1 {
		r.Tail.PosY = r.Head.PosY - 1
	}
}

func (r *Rope) Down(dist int) {
	r.Head.PosY -= dist
	r.Tail.PosX = r.Head.PosX
	if math.Abs(float64(r.Tail.PosY-r.Head.PosY)) > 1 {
		r.Tail.PosY = r.Head.PosY + 1
	}
}

func (r *Rope) Right(dist int) {
	r.Head.PosX += dist
	r.Tail.PosY = r.Head.PosY

	if math.Abs(float64(r.Tail.PosX-r.Head.PosX)) > 1 {
		r.Tail.PosX = r.Head.PosX - 1
	}
}

func (r *Rope) Left(dist int) {
	r.Head.PosX -= dist
	r.Tail.PosY = r.Head.PosY

	if math.Abs(float64(r.Tail.PosX-r.Head.PosX)) > 1 {
		r.Tail.PosX = r.Head.PosX + 1
	}
}
