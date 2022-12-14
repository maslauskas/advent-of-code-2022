package day12

import (
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

	return len(path)
}

type Path []Coord

type Coord struct {
	y int
	x int
}

type HeightMap struct {
	Rows  []string
	Start [2]int
	End   [2]int
	Queue Queue
}

func (m *HeightMap) Step() (Path, error) {
	path := m.Queue.Pop().Value.(Path)
	coord := path[len(path)-1]
	if reflect.DeepEqual(coord, Coord{m.End[0], m.End[1]}) {
		return path, nil
	}

	m.AddNeighbors(coord, path)

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

func (m *HeightMap) AddNeighbor(coord Coord, path Path, dir string) {
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
		m.Queue.Push(coords, int(dist))
		//fmt.Println(locChar+"->"+string(m.Rows[y][x]), coords)

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

	queue := Queue{
		[]QueueElement{{Path{{0, 0}}, 0}},
	}

	return HeightMap{
		Rows:  input,
		Start: start,
		End:   end,
		Queue: queue,
	}
}
