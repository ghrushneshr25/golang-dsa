/*
	Problem-2: Given an array, check whether the array is in sorted order with recursion.

	Time Complexity: O(n). Space Complexity: O(݊n) for recursive stack space
*/

package recursion

type SortedArray struct {
	Array  []int
	Result bool
}

func (soa *SortedArray) Recursion(index int) bool {
	if len(soa.Array) == 1 || index == 1 {
		return true
	}

	if soa.Array[index-1] < soa.Array[index-2] {
		return false
	}

	return soa.Recursion(index - 1)
}

func NewSortedArray(array []int) *SortedArray {
	return &SortedArray{
		Array: array,
	}
}

func (soa *SortedArray) Execute() {
	soa.Result = soa.Recursion(len(soa.Array))
}
