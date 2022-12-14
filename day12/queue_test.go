package day12

import (
	"reflect"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	t.Run("will add element", func(t *testing.T) {
		queue := Queue{}
		queue.Push(QueueElement{}, 1)
		AssertQueueLength(t, queue, 1)
	})

	t.Run("will dequeue one by one", func(t *testing.T) {
		queue := Queue{}
		queue.Push([]Coord{{0, 0}}, 5)
		queue.Push([]Coord{{1, 1}}, 1)
		queue.Push([]Coord{{2, 2}}, 3)

		queue.Pop()
		AssertQueueLength(t, queue, 2)

		queue.Pop()
		AssertQueueLength(t, queue, 1)

		queue.Pop()
		AssertQueueLength(t, queue, 0)
	})

	t.Run("will dequeue in the right order", func(t *testing.T) {
		queue := Queue{}
		wantEl3 := []Coord{{0, 0}}
		queue.Push(wantEl3, 5)
		wantEl1 := []Coord{{1, 1}}
		queue.Push(wantEl1, 1)
		wantEl2 := []Coord{{2, 2}}
		queue.Push(wantEl2, 3)

		el1 := queue.Pop()
		AssertQueueElement(t, el1, QueueElement{wantEl1, 1})
	})

	t.Run("will not override elements", func(t *testing.T) {
		queue := Queue{}
		queue.Push(Path{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}}, 3)
		queue.Push(Path{{0, 0}, {0, 1}, {1, 1}}, 5)
		queue.Push(Path{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {1, 1}}, 5)
		queue.Push(Path{{0, 0}, {1, 0}}, 6)

		elem := queue.Pop()
		AssertQueueElement(t, elem, QueueElement{Path{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}}, 3})

		queue.Push(Path{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}, {3, 2}}, 4)
		queue.Push(Path{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}, {2, 1}}, 4)

		next := queue.Pop()
		AssertQueueElement(t, next, QueueElement{Path{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}, {3, 2}}, 4})

		next2 := queue.Pop()
		AssertQueueElement(t, next2, QueueElement{Path{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}, {2, 1}}, 4})

	})
}

func AssertQueueElement(t *testing.T, got QueueElement, want QueueElement) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected element to be %v, got %v", want, got)
	}
}

func AssertQueueLength(t *testing.T, queue Queue, want int) {
	t.Helper()
	got := queue.Len()

	if want != got {
		t.Errorf("expected queue size of %d, got %d", want, got)
	}
}
