package queue

import "errors"

type CircularQueueNode struct {
	data interface{}
	next *CircularQueueNode
}

type CircularQueueLinkedList struct {
	front *CircularQueueNode
	rear  *CircularQueueNode
	size  int
}

func (q *CircularQueueLinkedList) EnQueue(data interface{}) {
	rear := new(CircularQueueNode)
	rear.data = data

	if q.IsEmpty() {
		q.front = rear
	} else {
		oldLast := q.rear
		oldLast.next = rear
	}
	q.rear = rear
	q.size++
}

func (q *CircularQueueLinkedList) DeQueue() (data interface{}, err error) {
	if q.IsEmpty() {
		q.rear = nil
		return nil, errors.New("unable to dequeue element from empty list")
	}

	data = q.front.data
	q.front = q.front.next
	q.size--
	return data, nil
}

func (q *CircularQueueLinkedList) FrontElement() (data interface{}, err error) {
	if q.IsEmpty() {
		return nil, errors.New("unable to read from empty list")
	}
	return q.front.data, nil
}

func (q *CircularQueueLinkedList) IsEmpty() bool {
	return q.front == nil
}

func (q *CircularQueueLinkedList) Length() int {
	return q.size
}
