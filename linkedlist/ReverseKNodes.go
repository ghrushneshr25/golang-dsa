/*
	For a given K value (K>0) reverse blocks of K nodes in a list.
 	Example: Input: 1 2 3 4 5 6 7 8 9 10,

	Output for different K values:
 		For K = 2 => 2 1 4 3 6 5 8 7 10 9
		For K = 3 => 3 2 1 6 5 4 9 8 7 10
		For K = 4 => 4 3 2 1 8 7 6 5 9 10

	This is an extension of swapping nodes in a linked list.
	1) Check if remaining list has K nodes.
		a. If yes get the pointer of K+1 th  node.
		b. Else return.
	2) Reverse first K nodes.
	3) Set next of last node (after reversal) to K+1  node.
	4) Move to K+1  node.
	5) Go to step 1.
	6) K−1  node of first K nodes becomes the new head if available. Otherwise, we can return the head.

*/

package linkedlist

type ReverseKNodes struct {
	head   *SingleLinkedList
	k      int
	Result []int
}

func NewReverseKNodes(val []int, k int) *ReverseKNodes {
	return &ReverseKNodes{
		head: ConvertArrayToSingleLinkedList(val),
		k:    k,
	}
}

func (rkn *ReverseKNodes) ReverseNodes() {
	rkn.head = rkn.reverseKNodes(rkn.head, rkn.k)
}

func (rkn *ReverseKNodes) reverseKNodes(head *SingleLinkedList, k int) *SingleLinkedList {
	current := head
	var prev, next *SingleLinkedList
	count := 0

	// Check if there are at least k nodes left to reverse
	temp := head
	for i := 0; i < k; i++ {
		if temp == nil {
			return head
		}
		temp = temp.GetNext()
	}

	// Reverse first k nodes of the linked list
	for current != nil && count < k {
		next = current.GetNext()
		current.SetNext(prev)
		prev = current
		current = next
		count++
	}

	// next is now pointing to (k+1)th node
	if next != nil {
		head.SetNext(rkn.reverseKNodes(next, k))
	}

	// prev is now the head of the reversed list
	return prev
}

func (rkn *ReverseKNodes) Execute() {
	rkn.ReverseNodes()
	rkn.Result = ConvertSingleLinkedListToArray(rkn.head)
}
