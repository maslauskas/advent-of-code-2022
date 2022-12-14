package day12

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
