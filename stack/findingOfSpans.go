/*
	Problem 22
	The problem is about finding the "span" for each element in an array ( A ). Let's break down the problem into parts to understand it more clearly.

	### Problem Explanation
		1. **Span ( S[i] ) Definition**:
		- For each element ( A[i] ) in the array ( A ), the span ( S[i] ) is defined as the **maximum number of consecutive elements** that immediately precede ( A[i] ) and satisfy the condition ( A[j] leq A[j+1] ).
		- In other words, it’s the length of the longest subarray ending at ( A[i] ) where each element is less than or equal to the next one.

		2. **Alternative Explanation**:
		- Another way to look at the problem is as follows: given an array ( A ), find the maximum difference ( j - i ) such that ( A[i] < A[j] ).
		- This means we're looking for two indices ( i ) and ( j ) (where ( j > i )) such that the element at index ( i ) is less than the element at index ( j ), and the difference between ( j ) and ( i ) (i.e., how far apart these indices are) is maximized.

	### Example to Clarify

		Let’s say we have an array:
			A = [6, 3, 4, 5, 2, 6, 4]

		For each element ( A[i] ), we want to find the span ( S[i] ).

		- For ( A[0] = 6 ): No previous elements, so ( S[0] = 1 ).
		- For ( A[1] = 3 ): No consecutive increasing elements before it, so ( S[1] = 1 ).
		- For ( A[2] = 4 ): ( A[1] leq A[2] ), so ( S[2] = 2 ).
		- For ( A[3] = 5 ): ( A[1] leq A[2] leq A[3] ), so ( S[3] = 3 ).
		- For ( A[4] = 2 ): No consecutive increasing elements before it, so ( S[4] = 1 ).
		- For ( A[5] = 6 ): ( A[4] leq A[5] ), so ( S[5] = 2 ).
		- For ( A[6] = 4 ): ( A[5] leq A[6] ), so ( S[6] = 2 ).

		So, the spans ( S ) for each element are:
			S = [1, 1, 2, 3, 1, 2, 2]

	### Goal of the Problem
		The main objective here is to compute this span ( S[i] ) for each element in the array efficiently,
		ideally in a way that doesn't involve examining every previous element for each ( i ) (which would be slow for large arrays). There may be an efficient approach involving stacks or specific algorithms that can help calculate these spans in linear time.
*/

package stack

type FindingOfSpans struct {
	stack  *LinkedListStack
	input  []int
	Result []int
}

func NewFindingOfSpans(input []int) *FindingOfSpans {
	return &FindingOfSpans{
		stack: NewLinkedListStack(),
		input: input,
	}
}

func (f *FindingOfSpans) findingSpans() []int {
	result, p := make([]int, len(f.input)), 0
	for i := 0; i < len(f.input); i++ {
		for !f.stack.IsEmpty() {
			top, _ := f.stack.Peek()
			if f.input[i] >= f.input[top.(int)] {
				f.stack.Pop()
			} else {
				break
			}
		}
		if f.stack.IsEmpty() {
			p = -1
		} else {
			top, _ := f.stack.Peek()
			p = top.(int)
		}
		result[i] = i - p
		f.stack.Push(i)
	}
	return result
}

func (f *FindingOfSpans) Execute() {
	f.Result = f.findingSpans()
}
