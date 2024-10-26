/* Problem-39 How will you check if the linked list is palindrome or not?

Algorithm
1. Get the middle of the linked list.
2. Reverse the second half of the linked list.
3. Compare the first half and second half.
4. Construct the original linked list by reversing the second half again and attaching it back to the first half.

*/

package linkedlist

type PalindromeLinkedList struct {
	head   *SingleLinkedList
	Result bool
}

func NewPalindromeLinkedList(val []int) *PalindromeLinkedList {
	head := ConvertArrayToSingleLinkedList(val)
	return &PalindromeLinkedList{
		head: head,
	}
}

func (pll *PalindromeLinkedList) IsPalindrome() bool {
	if pll.head == nil || pll.head.GetNext() == nil {
		return true
	}

	// Step 1: Find the middle of the linkedlist
	slow, fast := pll.head, pll.head
	for fast != nil && fast.GetNext() != nil {
		slow = slow.GetNext()
		fast = fast.GetNext().GetNext()
	}

	// Step 2: Reverse the second half of the linkedlist
	var prev *SingleLinkedList
	current := slow
	for current != nil {
		next := current.GetNext()
		current.SetNext(prev)
		prev = current
		current = next
	}

	// Step 3: Compare the first half and second half of the linkedlist
	firstHalf, secondHalf := pll.head, prev
	for secondHalf != nil {
		if firstHalf.GetData() != secondHalf.GetData() {
			return false
		}
		firstHalf = firstHalf.GetNext()
		secondHalf = secondHalf.GetNext()
	}

	// Step 4: Restore the original linkedlist
	// reverse the second half again and attach it to restore the original list

	var prevRestore *SingleLinkedList
	currentRestore := prev
	for currentRestore != nil {
		next := currentRestore.GetNext()
		currentRestore.SetNext(prevRestore)
		prevRestore = currentRestore
		currentRestore = next
	}

	return true
}

func (pll *PalindromeLinkedList) Execute() {
	pll.Result = pll.IsPalindrome()
}
