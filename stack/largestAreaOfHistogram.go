/*
	Problem-25
		Largest rectangle under histogram: A histogram is a polygon composed of a sequence of rectangles
		aligned at a common base line. For simplicity, assume that the rectangles have equal widths but
		may have different heights. For example, the figure on the left  shows a histogram that consists
		of rectangles with the heights 3,2 ,5,6,1,4,4, measured in units where 1 is the width of the
		rectangles. Here our problem is: given an array with heights of rectangles (assuming width is 1),
		we need to find the largest rectangle possible. For the given example, the largest rectangle
		is the shared part.

	Linear search using a stack of incomplete subproblems: There are many ways of solving this problem.
	Judge has given a nice algorithm for this problem which is based on stack. We process the elements
	in left-to-right order and maintain a stack of information about started but yet unfinished sub
	histograms. If the stack is empty, open a new subproblem by pushing the element onto the stack.
	Otherwise compare it to the element on top of the stack. If the new one is greater we again push it.
	If the new one is equal we skip it. In all these cases, we continue with the next new element.

	If the new one is less, we finish the topmost subproblem by updating the maximum area with respect to the element at the top of the stack.
	Then, we discard the element at the top, and repeat the procedure keeping the current new element. This way, all subproblems are finished
	when the stack becomes empty, or its top element is less than or equal to the new element, leading to the actions described above. If all
	elements have been processed, and the stack is not yet empty, we finish the remaining subproblems by updating the maximum area with
	respect to the elements at the top.

	Space Complexity: O(n)
	Time Complexity: O(n)
*/

package stack

type MaxAreaOfHistogram struct {
	stack  *LinkedListStack
	input  []int
	Result int
}

func NewMaxAreaOfHistogram(input []int) *MaxAreaOfHistogram {
	return &MaxAreaOfHistogram{
		stack: NewLinkedListStack(),
		input: input,
	}
}

func (f *MaxAreaOfHistogram) Execute() {
	f.Result = f.maxArea()
}

func (f *MaxAreaOfHistogram) maxArea() int {
	if len(f.input) == 0 {
		return 0
	}

	maxArea, i := 0, 0
	for i < len(f.input) {
		peek, _ := f.stack.Peek()
		if f.stack.IsEmpty() || f.input[i] >= f.input[peek.(int)] {
			f.stack.Push(i)
			i++
		} else {
			top, _ := f.stack.Pop()
			width := i
			if !f.stack.IsEmpty() {
				peek, _ := f.stack.Peek()
				width = i - peek.(int) - 1
			}
			maxArea = max(maxArea, f.input[top.(int)]*width)
		}
	}

	for !f.stack.IsEmpty() {
		top, _ := f.stack.Pop()
		width := i
		if !f.stack.IsEmpty() {
			peek, _ := f.stack.Peek()
			width = i - peek.(int) - 1
		}
		maxArea = max(maxArea, f.input[top.(int)]*width)
	}

	return maxArea
}
