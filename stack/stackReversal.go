/*
Problem-11
	Given a stack, how to reverse the contents of the stack using only stack operations (push and pop)?
Algorithm
- First pop all the elements of the stack till it becomes empty.
- For each upward step in recursion, insert the element at the bottom of the stack.
*/

package stack

type StackReversal struct {
	stack  *LinkedListStack
	Result []int
}

func NewStackReversal(values []int) *StackReversal {
	stack := NewLinkedListStack()
	for _, value := range values {
		stack.Push(value)
	}
	return &StackReversal{
		stack:  stack,
		Result: []int{},
	}
}

func (sr *StackReversal) ReverseStack() {
	temp, err := sr.stack.Pop()
	if err != nil {
		return
	}
	sr.ReverseStack()
	sr.InsertAtBottom(temp)
}

func (sr *StackReversal) InsertAtBottom(item interface{}) {
	if sr.stack.IsEmpty() {
		sr.stack.Push(item)
		return
	}
	temp, _ := sr.stack.Pop()
	sr.InsertAtBottom(item)
	sr.stack.Push(temp)
}

func (sr *StackReversal) Execute() {
	sr.ReverseStack()
	for !sr.stack.IsEmpty() {
		item, _ := sr.stack.Pop()
		sr.Result = append(sr.Result, item.(int))
	}
}
