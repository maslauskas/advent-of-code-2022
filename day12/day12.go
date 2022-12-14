package day12

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func Part1(input []string) int {
	heightmap := CreateHeightMap(input)
	for heightmap.Queue.Len() > 0 {
		heightmap.Step()
	}

	min := 999999
	fmt.Println(len(heightmap.SuccessfulPaths))
	for i := 0; i < len(heightmap.SuccessfulPaths); i++ {
		if len(heightmap.SuccessfulPaths[i]) < min {
			min = len(heightmap.SuccessfulPaths[i])
		}
	}
	return min
}

type Path []Coord

type Coord struct {
	y int
	x int
}

type HeightMap struct {
	Rows            []string
	Start           [2]int
	End             [2]int
	Queue           Queue
	SuccessfulPaths []Path
}

func (m *HeightMap) Step() {
	elem := m.Queue.Pop()
	path := elem.Value.(Path)
	coord := path[len(path)-1]
	if reflect.DeepEqual(coord, Coord{m.End[0], m.End[1]}) {
		m.SuccessfulPaths = append(m.SuccessfulPaths, path)
		return
	}

	m.AddNeighbors(coord, path)
}

func (m *HeightMap) AddNeighbors(coord Coord, path []Coord) {
	m.AddNeighbor(coord, path, "TOP")
	m.AddNeighbor(coord, path, "RIGHT")
	m.AddNeighbor(coord, path, "BOTTOM")
	m.AddNeighbor(coord, path, "LEFT")
}

func (m *HeightMap) AddNeighbor(coord Coord, path Path, dir string) {
	newCoord := Coord{coord.y, coord.x}
	locChar := string(m.Rows[coord.y][coord.x])
	locVal := GetScore(locChar)

	switch dir {
	case "TOP":
		newCoord.y--
	case "RIGHT":
		newCoord.x++
	case "BOTTOM":
		newCoord.y++
	case "LEFT":
		newCoord.x--
	}

	// check out of bounds
	if newCoord.y < 0 || newCoord.x >= len(m.Rows[0]) || newCoord.y >= len(m.Rows) || newCoord.x < 0 {
		return
	}

	// check if coord was visited in path
	for i := 0; i < len(path)-1; i++ {
		if reflect.DeepEqual(path[i], newCoord) {
			return
		}
	}

	top := GetScore(string(m.Rows[newCoord.y][newCoord.x]))
	dist := math.Abs(float64(m.End[0]-newCoord.y)) + math.Abs(float64(m.End[1]-newCoord.x))
	if top-1 <= locVal {
		coords := append(path, newCoord)
		m.Queue.Push(coords, int(dist))
		//fmt.Println(locChar+"->"+string(m.Rows[newCoord.y][newCoord.x]), int(dist), coords)
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
