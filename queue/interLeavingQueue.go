package queue

import "container/list"

func InterLeavingQueue(q *list.List) {

	if q.Len()%2 != 0 {
		return // return if the queue has an odd number of elements
	}

	stack := list.New()
	halfSize := q.Len() / 2
	for i := 0; i < halfSize; i++ {
		stack.PushBack(q.Remove(q.Front()))
	}

	for stack.Len() > 0 {
		q.PushBack(stack.Remove(stack.Back()))
	}

	for i := 0; i < halfSize; i++ {
		q.PushBack(q.Remove(q.Front()))
	}

	for i := 0; i < halfSize; i++ {
		stack.PushBack(q.Remove(q.Front()))
	}

	for stack.Len() > 0 {
		q.PushBack(stack.Remove(stack.Back()))
		q.PushBack(q.Remove(q.Front()))
	}
}
