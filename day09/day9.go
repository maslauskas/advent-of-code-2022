package day09

type Point struct {
	PosX int
	PosY int
}
type Rope struct {
	Head Point
	Tail Point
}

func (r *Rope) Up() {
	r.Head.PosY += 1
}

func (r *Rope) Down() {
	r.Head.PosY -= 1
}

func (r *Rope) Right() {
	r.Head.PosX += 1
}

func (r *Rope) Left() {
	r.Head.PosX -= 1
}
