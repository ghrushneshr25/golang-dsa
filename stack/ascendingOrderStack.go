/*
	Problem-26
				Give an algorithm for sorting a stack in ascending order.
				We should not make any assumptions about how the stack is implemented.

	Time Complexity: O(n^2).
	Space Complexity: O(n), for stack.
*/

package stack

type AscendingOrderStack struct {
	inputStack *LinkedListStack
	Result     *LinkedListStack
}

func NewAscendingOrderStack(input []int) *AscendingOrderStack {
	inputStack := NewLinkedListStack()
	index := len(input) - 1
	for index >= 0 {
		inputStack.Push(input[index])
		index--
	}
	return &AscendingOrderStack{
		inputStack: inputStack,
	}
}

func (obj *AscendingOrderStack) sortStack() *LinkedListStack {
	rstk := NewLinkedListStack()
	for !obj.inputStack.IsEmpty() {
		temp, _ := obj.inputStack.Pop()
		for !rstk.IsEmpty() {
			peek, _ := rstk.Peek()
			if peek.(int) > temp.(int) {
				pop, _ := rstk.Pop()
				obj.inputStack.Push(pop)
			} else {
				break
			}
		}
		rstk.Push(temp)
	}
	return rstk
}

func (obj *AscendingOrderStack) Execute() {
	obj.Result = obj.sortStack()
}
