package linkedlist

type ReverseSingleLinkedList struct {
	Head   *SingleLinkedList
	Result []int
}

func NewReverseSingleLinkedList(values []int) *ReverseSingleLinkedList {
	reverseSingleLinkedList := ReverseSingleLinkedList{}
	for _, value := range values {
		InsertAtEndofSingleLL(&reverseSingleLinkedList.Head, value)
	}
	return &reverseSingleLinkedList
}

/*
	Time Complexity: O(n), for scanning the list of size n
	Space Complexity: O(1), for creating a temporary variable.
*/

func (rsll *ReverseSingleLinkedList) IterativeApproach() {
	current := rsll.Head
	var prev *SingleLinkedList

	for current != nil {
		next := current.GetNext()
		current.SetNext(prev)
		prev = current
		current = next
	}

	rsll.Head = prev
}

/*
	Time Complexity: O(n), for scanning the list of size n
	Space Complexity: O(n), for storing the list in the recursive stack.
*/

func (rsll *ReverseSingleLinkedList) RecursiveApproach() {
	rsll.Head = rsll.reverseRecursively(rsll.Head)
}

func (rsll *ReverseSingleLinkedList) reverseRecursively(head *SingleLinkedList) *SingleLinkedList {
	if head == nil || head.GetNext() == nil {
		return head
	}
	newHead := rsll.reverseRecursively(head.GetNext())
	head.GetNext().SetNext(head)
	head.SetNext(nil)
	return newHead
}

func (rsll *ReverseSingleLinkedList) ExecuteByName(methodName string) {
	reverse := map[string]func(){
		"IterativeApproach": rsll.IterativeApproach,
		"RecursiveApproach": rsll.RecursiveApproach,
	}
	reverse[methodName]()
	current := rsll.Head
	rsll.Result = []int{}
	for current != nil {
		rsll.Result = append(rsll.Result, current.GetData())
		current = current.GetNext()
	}
}
