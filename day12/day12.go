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
	return -1
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

func (m *HeightMap) Step() (Path, error) {
	path := heap.Pop(m.Queue).(Path)
	fmt.Println(path)
	coord := path.Coords[len(path.Coords)-1]
	if reflect.DeepEqual(coord, m.End) {
		return path, nil
	}
	y := coord.y
	x := coord.x
	locVal := GetScore(string(m.Rows[y][x]))
	fmt.Print(string(m.Rows[y][x]))

	// check and add top
	if y-1 > 0 {
		m.AddPath(y-1, x, locVal, path)
	}

	// check and add right
	if x+1 <= len(m.Rows[0]) {
		m.AddPath(y, x+1, locVal, path)
	}

	// check and add bottom
	if y+1 < len(m.Rows) {
		m.AddPath(y+1, x, locVal, path)
	}

	// check and add left
	if x-1 > 0 {
		m.AddPath(y, x-1, locVal, path)
	}

	fmt.Println()
	return Path{}, errors.New("Path not found")
}

func (m *HeightMap) AddPath(y int, x int, locVal int, path Path) {
	key := fmt.Sprintf("%d:%d", y, x)
	_, ok := m.Visited[key]
	if ok {
		// do not add visited routes
		return
	}
	top := GetScore(string(m.Rows[y][x]))
	fmt.Print(string(m.Rows[y][x]))
	dist := math.Abs(float64(m.End[0]-y)) + math.Abs(float64(m.End[1]-x))
	if top-1 <= locVal {
		coords := append(path.Coords, Coord{y, x})
		heap.Push(m.Queue, Path{Score: int(dist), Coords: coords})
		m.Visited[key] = Coord{y, x}
	}
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
