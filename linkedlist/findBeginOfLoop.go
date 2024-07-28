/*
	Problem 12: Check whether the given linked list is NULL-terminated or not.
				If there is a cycle find the start node of the loop.
*/

package linkedlist

type FindBeginOfLoop struct {
	Head   *SingleLinkedList
	Result int
}

func NewFindBeginOfLoop(head *SingleLinkedList) *FindBeginOfLoop {
	return &FindBeginOfLoop{
		Head: head,
	}
}

/*
	Time Complexity: O(n).
	Space Complexity: O(1).
*/

func (fbl *FindBeginOfLoop) FloydCycle() int {
	fastPtr := fbl.Head
	slowPtr := fbl.Head
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
		slowPtr = fbl.Head
		for slowPtr != fastPtr {
			slowPtr = slowPtr.GetNext()
			fastPtr = fastPtr.GetNext()
		}
		return slowPtr.GetData()
	}
	return -1
}

func (fbl *FindBeginOfLoop) Execute() {
	fbl.Result = fbl.FloydCycle()
}
