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

type QueueElement struct {
}

func TestPriorityQueue(t *testing.T) {
	t.Run("will add element", func(t *testing.T) {
		queue := Queue{}
		queue.Push(QueueElement{}, 1)

		want := 1
		got := queue.Len()

		if want != got {
			t.Errorf("expected queue size of %d, got %d", want, got)
		}
	})
}
