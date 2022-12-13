package day12

import (
	"container/heap"
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
}
