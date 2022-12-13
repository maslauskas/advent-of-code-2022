package day12

import (
	"container/heap"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
)

func Part1(input []string) int {
	heightmap := CreateHeightMap(input)
	path := heightmap.FindPath()
	fmt.Println(path)

	return len(path.Coords)
}

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
	a := h[i]
	b := h[j]

	h[j] = a
	h[i] = b
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

func (m *HeightMap) Step() (Path, error) {
	path := heap.Pop(m.Queue).(Path)
	fmt.Println(path)
	fmt.Println(m.Queue.Len())
	coord := path.Coords[len(path.Coords)-1]
	if reflect.DeepEqual(coord, Coord{m.End[0], m.End[1]}) {

		return path, nil
	}

	m.AddNeighbors(coord, path.Coords)

	return Path{}, errors.New("Path not found")
}

func (m *HeightMap) FindPath() Path {
	for m.Queue.Len() > 0 {
		found, err := m.Step()
		if err == nil {
			return found
		}
	}

	panic("no path found")
}

func (m *HeightMap) AddNeighbors(coord Coord, path []Coord) {
	m.AddNeighbor(coord, path, "TOP")
	m.AddNeighbor(coord, path, "RIGHT")
	m.AddNeighbor(coord, path, "BOTTOM")
	m.AddNeighbor(coord, path, "LEFT")
}

func (m *HeightMap) AddNeighbor(coord Coord, path []Coord, dir string) {
	y := coord.y
	x := coord.x
	locChar := string(m.Rows[y][x])
	locVal := GetScore(locChar)

	switch dir {
	case "TOP":
		y = y - 1
	case "RIGHT":
		x = x + 1
	case "BOTTOM":
		y = y + 1
	case "LEFT":
		x = x - 1
	}

	// check out of bounds
	if y < 0 || x >= len(m.Rows[0]) || y >= len(m.Rows) || x < 0 {
		return
	}

	// check if coord was visited in path
	visited := false
	for _, c := range path {
		if reflect.DeepEqual(c, Coord{y, x}) {
			visited = true
		}
	}
	if visited {
		return
	}

	top := GetScore(string(m.Rows[y][x]))
	dist := math.Abs(float64(m.End[0]-y)) + math.Abs(float64(m.End[1]-x))
	if top-1 <= locVal {
		coords := append(path, Coord{y, x})
		newPath := Path{Score: int(dist), Coords: coords}
		heap.Push(m.Queue, newPath)
		fmt.Println(locChar+"->"+string(m.Rows[y][x]), newPath)

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
		Rows:  input,
		Start: start,
		End:   end,
		Queue: &h,
		Visited: map[string]Coord{
			"0:0": Coord{0, 0},
		},
	}
}
