package day12

import "testing"

type Queue struct {
	Elements []QueueElement
}

func (q *Queue) Push(element QueueElement, priority int) {
	// less priority is better
	q.Elements = append(q.Elements, element)
}

func (q *Queue) Len() int {
	return len(q.Elements)
}

func (q *Queue) Pop() QueueElement {
	elem := q.Elements[q.Len()-1]
	q.Elements = q.Elements[:q.Len()-1]

	return elem
}

type QueueElement struct {
	Coords []Coord
}

func TestPriorityQueue(t *testing.T) {
	t.Run("will add element", func(t *testing.T) {
		queue := Queue{}
		queue.Push(QueueElement{}, 1)
		AssertQueueLength(t, queue, 1)
	})

	t.Run("will dequeue one by one", func(t *testing.T) {
		queue := Queue{}
		queue.Push(QueueElement{[]Coord{{0, 0}}}, 5)
		queue.Push(QueueElement{[]Coord{{1, 1}}}, 1)
		queue.Push(QueueElement{[]Coord{{2, 2}}}, 3)

		queue.Pop()
		AssertQueueLength(t, queue, 2)

		queue.Pop()
		AssertQueueLength(t, queue, 1)

		queue.Pop()
		AssertQueueLength(t, queue, 0)
	})
}

func AssertQueueLength(t *testing.T, queue Queue, want int) {
	t.Helper()
	got := queue.Len()

	if want != got {
		t.Errorf("expected queue size of %d, got %d", want, got)
	}
}
