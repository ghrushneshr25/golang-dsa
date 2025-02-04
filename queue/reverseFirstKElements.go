package queue

import "container/list"

func ReverseFirstKElements(q *list.List, k int) {

	if q == nil || q.Len() < k || k < 1 {
		return
	}

	stack := list.New()

	for i := 0; i < k; i++ {
		stack.PushBack(q.Remove(q.Front()))
	}

	for stack.Len() > 0 {
		q.PushBack(stack.Remove(stack.Back()))
	}

	for i := 0; i < q.Len()-k; i++ {
		q.PushBack(q.Remove(q.Front()))
	}
}
