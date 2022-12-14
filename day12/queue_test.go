package day12

import (
	"reflect"
	"testing"
)

type Queue struct {
	Elements []QueueElement
}

func (q *Queue) Push(value interface{}, priority int) {
	// less priority is better
	newElement := QueueElement{value, priority}
	if q.Len() == 0 {
		q.Elements = append(q.Elements, newElement)
		return
	}

	q.Elements = append(q.Elements, newElement)

	i := q.Len() - 2
	for i >= 0 {
		if q.Elements[i].Priority > priority {
			q.Elements[i+1] = q.Elements[i]
			i--
		} else {
			break
		}
	}
	q.Elements[i+1] = newElement
}

func (q *Queue) Len() int {
	return len(q.Elements)
}

func (q *Queue) Pop() QueueElement {
	elem := q.Elements[0]
	q.Elements = q.Elements[1:]

	return elem
}

type QueueElement struct {
	Value    interface{}
	Priority int
}

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
