package day12

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

//func Part1(input []string) int {
//	return -1
//}

type PathHeap []Path

type Path struct {
	Score  int
	Coords []Coord
}

func (h PathHeap) Len() int {
	return len(h)
}
func (h PathHeap) Less(i, j int) bool {
	return h[i].Score < h[j].Score
}
func (h PathHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PathHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Path))
}

func (h *PathHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

type Coord struct {
	y int
	x int
}

type HeightMap struct {
	Rows    []string
	Start   [2]int
	End     [2]int
	Queue   *PathHeap
	Visited map[string]Coord
}

func (m *HeightMap) Step() {
	path := heap.Pop(m.Queue).(Path)
	coord := path.Coords[len(path.Coords)-1]
	y := coord.y
	x := coord.x
	locVal := GetScore(string(m.Rows[y][x]))

	// check and add top
	if y-1 > 0 {
		m.AddPath(y-1, x, locVal, path)
	}

	// check and add right
	if x+1 <= len(m.Rows[0]) {
		m.AddPath(y, x+1, locVal, path)
	}

	// check and add bottom
	if y+1 <= len(m.Rows) {
		m.AddPath(y+1, x, locVal, path)
	}

	// check and add left
	if x-1 > 0 {
		m.AddPath(y, x-1, locVal, path)
	}
}

func (m *HeightMap) AddPath(y int, x int, locVal int, path Path) {
	top := GetScore(string(m.Rows[y][x]))
	dist := math.Abs(float64(m.End[0]-y)) + math.Abs(float64(m.End[1]-x))
	if top-1 <= locVal {
		coords := append(path.Coords, Coord{y, x})
		heap.Push(m.Queue, Path{Score: int(dist), Coords: coords})
		key := fmt.Sprintf("%d:%d", y, x)
		m.Visited[key] = Coord{y, x}
	}
}

func GetScore(char string) int {
	alpha := "Sabcdefghijklmnopqrstuvwxyz"
	return strings.Index(alpha, char)
}

func CreateHeightMap(input []string) HeightMap {
	// find start
	var start [2]int
	var end [2]int
	for y, row := range input {
		for x, char := range row {
			if string(char) == "S" {
				start = [2]int{y, x}
			}

			if string(char) == "E" {
				end = [2]int{y, x}
			}
		}
	}

	h := PathHeap{
		{Score: 0, Coords: []Coord{{0, 0}}},
	}
	heap.Init(&h)

	return HeightMap{
		Rows:    input,
		Start:   start,
		End:     end,
		Queue:   &h,
		Visited: map[string]Coord{},
	}
}
