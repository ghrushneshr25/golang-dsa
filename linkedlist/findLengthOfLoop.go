/*
	Problem 15: Check whether the given linked list is NULL-terminated.
				If there is a cycle, find the length of the loop.
*/

package linkedlist

type FindLengthOfLoop struct {
	Head   *SingleLinkedList
	Result int
}

func NewFindLengthOfLoop(head *SingleLinkedList) *FindLengthOfLoop {
	return &FindLengthOfLoop{
		Head: head,
	}
}

func (fll *FindLengthOfLoop) FloydCycle() int {
	fastPtr := fll.Head
	slowPtr := fll.Head
	var loopExists bool
	for fastPtr != nil && fastPtr.GetNext() != nil {
		fastPtr = fastPtr.GetNext().GetNext()
		slowPtr = slowPtr.GetNext()
		if fastPtr == slowPtr {
			loopExists = true
			break
		}
	}
	if loopExists {
		count := 1
		fastPtr = fastPtr.GetNext()
		for slowPtr != fastPtr {
			fastPtr = fastPtr.GetNext()
			count++
		}
		return count
	}
	return -1
}

func (fll *FindLengthOfLoop) Execute() {
	fll.Result = fll.FloydCycle()
}
