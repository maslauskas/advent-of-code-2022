package day12

import (
	"adventofcode/helpers"
	"container/heap"
	"reflect"
	"testing"
)

func TestPathfinding(t *testing.T) {
	t.Run("priority queue", func(t *testing.T) {
		h := PathHeap{}
		heap.Init(&h)

		heap.Push(&h, Path{Score: 1, Coords: []Coord{{0, 0}}})
		heap.Push(&h, Path{Score: 5, Coords: []Coord{{0, 0}}})
		heap.Push(&h, Path{Score: 2, Coords: []Coord{{0, 0}}})

		i1 := heap.Pop(&h).(Path)
		if i1.Score != 1 {
			t.Errorf("expected item to have score 1, got %d", i1.Score)
		}

		i2 := heap.Pop(&h).(Path)
		if i2.Score != 2 {
			t.Errorf("expected item to have score 1, got %d", i2.Score)
		}

		i3 := heap.Pop(&h).(Path)
		if i3.Score != 5 {
			t.Errorf("expected item to have score 1, got %d", i3.Score)
		}
	})

	t.Run("will create heightmap", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		heightmap := CreateHeightMap(input)

		wantRows := 5
		gotRows := len(heightmap.Rows)

		if wantRows != gotRows {
			t.Errorf("expected rows count to be %d, got %d", wantRows, gotRows)
		}
	})

	t.Run("will find start point", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		heightmap := CreateHeightMap(input)

		want := [2]int{0, 0}
		got := heightmap.Start

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected start to be at %v, got %v", want, got)
		}
	})

	t.Run("will find end point", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		heightmap := CreateHeightMap(input)

		want := [2]int{2, 5}
		got := heightmap.End

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected end to be at %v, got %v", want, got)
		}
	})

	t.Run("will add possible paths to heap", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		heightmap := CreateHeightMap(input)
		heightmap.Step()

		want := 2
		got := heightmap.Queue.Len()

		if want != got {
			t.Errorf("expected queue length of %d, got %d", want, got)
		}
	})

	t.Run("will add visited coords", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		heightmap := CreateHeightMap(input)
		heightmap.Step()

		want := 2
		got := len(heightmap.Visited)
		if want != got {
			t.Errorf("expected visited count of %d, got %d", want, got)
		}
	})

	t.Run("will remove first path from the queue", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		heightmap := CreateHeightMap(input)

		// step twice
		heightmap.Step()
		heightmap.Step()

		want := 3
		got := heightmap.Queue.Len()

		if want != got {
			t.Errorf("expected queue length of %d, got %d", want, got)
		}
	})

	t.Run("will pop path that is closest to end", func(t *testing.T) {
		input := helpers.ReadInput("./example.txt")
		heightmap := CreateHeightMap(input)

		// step twice
		heightmap.Step()
		heightmap.Step()

		// pop last element
		path := heap.Pop(heightmap.Queue)
		want := Path{Score: 5, Coords: []Coord{{0, 0}, {0, 1}, {0, 2}}}

		if !reflect.DeepEqual(path, want) {
			t.Errorf("expected next queue element to be%v, got %v", want, path)
		}
	})

	//t.Run("part 1 example case", func(t *testing.T) {
	//	input := helpers.ReadInput("./example.txt")
	//	got := Part1(input)
	//
	//	want := 31
	//	if want != got {
	//		t.Errorf("expected result of %d, got %d", want, got)
	//	}
	//})
}
