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

		heap.Push(&h, Path{Score: 1, Coords: []string{""}})
		heap.Push(&h, Path{Score: 5, Coords: []string{""}})
		heap.Push(&h, Path{Score: 2, Coords: []string{""}})

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
}
