package stack

import (
	"fmt"
	"golang-dsa/linkedlist"
)

type LinkedListStack struct {
	top    *linkedlist.SingleLinkedList
	length int
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		top:    nil,
		length: 0,
	}
}

func (stack *LinkedListStack) IsEmpty() bool {
	return stack.top == nil
}

func (stack *LinkedListStack) Size() int {
	return stack.length
}

func (stack *LinkedListStack) Push(data interface{}) {
	node := &linkedlist.SingleLinkedList{
		CustomData: data,
	}
	node.SetNext(stack.top)
	stack.top = node
	stack.length++
}

func (stack *LinkedListStack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return 0, ErrStackEmpty
	}
	value := stack.top.CustomData
	stack.top = stack.top.GetNext()
	stack.length--
	return value, nil
}

func (stack *LinkedListStack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return 0, ErrStackEmpty
	}
	return stack.top.CustomData, nil
}

func (stack *LinkedListStack) ToString() string {
	if stack.IsEmpty() {
		return "[]"
	}
	var result string
	temp := stack.top

	for temp != nil {
		result += fmt.Sprintf(" %+v ", temp.CustomData)
		temp = temp.GetNext()
	}
	return "[" + result + "]"
}
