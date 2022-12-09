package day09

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
}

func (r *Rope) Down(dist int) {
	r.Head.PosY -= dist
}

func (r *Rope) Right(dist int) {
	r.Head.PosX += dist
}

func (r *Rope) Left(dist int) {
	r.Head.PosX -= dist
}
