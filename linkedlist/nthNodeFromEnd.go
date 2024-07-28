/*
Problem 2: Find nth node from the end of a Linked List.
*/
package linkedlist

type NthNodeFromEnd struct {
	Head   *SingleLinkedList
	Result int
	N      int
}

func NewNthNodeFromEnd(values []int, n int) *NthNodeFromEnd {
	nthNodeFromEnd := &NthNodeFromEnd{}
	for _, value := range values {
		InsertAtEndofSingleLL(&nthNodeFromEnd.Head, value)
	}
	nthNodeFromEnd.N = n
	return nthNodeFromEnd
}

func (nnfe *NthNodeFromEnd) Execute(methodName string) {
	methods := map[string]func() int{
		"BruteForceMethod": nnfe.BruteForceMethod,
		"HashMapMethod":    nnfe.HashMapMethod,
		"TwoPointerMethod": nnfe.TwoPointerMethod,
		"RecursionMethod":  nnfe.RecursionMethod,
	}
	nnfe.Result = methods[methodName]()
}

func (nnfe *NthNodeFromEnd) BruteForceMethod() int {
	/*
		Time Complexity: O(n^2).
		For scanning the remaining list (from current node) for each node.
		Space Complexity: O(1).
	*/
	current := nnfe.Head
	len := GetSingleLinkedListLength(current)
	if len < nnfe.N {
		return -1
	}
	for current.GetNext() != nil && len > nnfe.N {
		current = current.GetNext()
		len = GetSingleLinkedListLength(current)
	}
	if len == nnfe.N {
		return current.GetData()
	}
	return -1
}

func (nnfe *NthNodeFromEnd) HashMapMethod() int {
	/*
		M => Length of the linked list.
		Time Complexity: T(M) = O(M) Time for creating the hash table.
		Space Complexity: O(M) for hash table of size
	*/
	len := GetSingleLinkedListLength(nnfe.Head)
	if len < nnfe.N {
		return -1
	}
	hashMap := map[int]*SingleLinkedList{}
	current := nnfe.Head
	currentPosition := 1
	for current != nil {
		hashMap[currentPosition] = current
		current = current.GetNext()
		currentPosition++
	}
	if val, ok := hashMap[len-nnfe.N+1]; ok {
		return val.GetData()
	}
	return -1
}

func (nnfe *NthNodeFromEnd) TwoPointerMethod() int {
	/*
		Time Complexity: O(n)
		Space Complexity: O(1)
	*/

	if nnfe.N < 0 {
		return -1
	}
	var pTemp, pNthNode *SingleLinkedList
	pTemp = nnfe.Head
	for i := 1; i < nnfe.N && pTemp != nil; i++ {
		pTemp = pTemp.GetNext()
	}
	if pTemp == nil {
		return -1
	}
	pNthNode = nnfe.Head
	for pTemp.GetNext() != nil {
		pTemp = pTemp.GetNext()
		pNthNode = pNthNode.GetNext()
	}
	return pNthNode.GetData()
}

func (nnfe *NthNodeFromEnd) RecursionMethod() int {
	/*
		Time Complexity: O(n) for pre recursive calls and O(n) for post recursive calls, which is => O(2n) = O(n).
		Space Complexity: O(n) for recursive stack.
	*/
	counter := 0
	node := nnfe.recursionLogic(nnfe.Head, nnfe.N, &counter)
	if node == nil {
		return -1
	}
	return node.GetData()
}

func (nnfe *NthNodeFromEnd) recursionLogic(head *SingleLinkedList, n int, counter *int) *SingleLinkedList {
	if head != nil {
		result := nnfe.recursionLogic(head.GetNext(), n, counter)
		(*counter)++
		if n == *counter {
			return head
		}
		return result
	}
	return nil
}
