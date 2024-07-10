/*
	Problem 3: Generate all the strings of ݊n bits. Assume A[0...n-1] is an array of size ݊n.

	Using Subtraction and Conquer Master theorem, we get T(n)=O(2^n).
	This means the algorithm for generating bit-strings is optimal.
*/

package backtracking

import "strings"

type NBitString struct {
	Size         int
	Combinations []string
	Content      []string
}

func (nbit *NBitString) Recursion(current int) {
	if current <= 0 {
		if len(nbit.Content) > 0 {
			nbit.Combinations = append(nbit.Combinations, strings.Join(nbit.Content, ""))
		}
		return
	}
	nbit.Content[current-1] = "0"
	nbit.Recursion(current - 1)
	nbit.Content[current-1] = "1"
	nbit.Recursion(current - 1)
}

func NewNBitString(size int) *NBitString {
	nbitString := &NBitString{}
	nbitString.Size = size
	nbitString.Content = make([]string, size)
	return nbitString
}

func (nbit *NBitString) Execute() {
	nbit.Recursion(nbit.Size)
}
