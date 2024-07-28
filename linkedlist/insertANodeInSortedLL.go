/*
	Problem 16: Insert a node in a sorted linked list
*/

package linkedlist

type InsertANodeInSortedLL struct {
	Head   *SingleLinkedList
	Result []int
}

func NewInsertANodeInSortedLL(head *SingleLinkedList) *InsertANodeInSortedLL {
	return &InsertANodeInSortedLL{
		Head: head,
	}
}

/*
	Time Complexity: O(N)
	Space Complexity: O(1)
*/

func (sll *InsertANodeInSortedLL) InsertNode(data int) {
	newNode := NewSingleLinkedList(data)

	if sll.Head == nil {
		sll.Head = newNode
		return
	}

	current := sll.Head
	var temp *SingleLinkedList
	for current != nil && current.GetData() < data {
		temp = current
		current = current.GetNext()
	}

	newNode.SetNext(current)
	if temp != nil {
		temp.SetNext(newNode)
	} else {
		sll.Head = newNode
	}
}

func (sll *InsertANodeInSortedLL) Execute(data int) {
	sll.InsertNode(data)
	sll.Result = []int{}
	current := sll.Head
	for current != nil {
		sll.Result = append(sll.Result, current.GetData())
		current = current.GetNext()
	}
}
