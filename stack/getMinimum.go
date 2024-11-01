/*
	Problem-6/7 How to design a stack such that getMinimum() should be O(1)?
	 Take an auxiliary stack that maintains the minimum of all values in the stack. 
	 Also, assume that each element of the stack is less than its below elements. 
	 For simplicity let us call the auxiliary stack min stack.   
*/

package stack

type GetMinimum struct {
	input  []int
	Result int
}

func NewGetMinimum(input []int) *GetMinimum {
	return &GetMinimum{
		input: input,
	}
}

func (gm *GetMinimum) GetMinimum() int {
	minStack := NewLinkedListStack()

	push := func(i int) {
		min, _ := minStack.Peek()
		if minStack.IsEmpty() || min.(int) > i {
			minStack.Push(i)
		}
	}

	for _, i := range gm.input {
		push(i)
	}

	min, _ := minStack.Peek()
	return min.(int)
}

func (gm *GetMinimum) Execute() {
	gm.Result = gm.GetMinimum()
}
