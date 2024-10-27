/*
Problem-48
	Find modular node from the end: Given a singly linked list,
	write a function to find the first element from the end whose
	n%k == 0, where n is the number of elements in the list and k is an integer constant.
	For example, if n=19 and k=3 then we should return 16 node.
*/

package linkedlist

type FindModularNodeFromEnd struct {
	head   *SingleLinkedList
	k      int
	Result int
}

func NewFindModularNodeFromEnd(val []int, k int) *FindModularNodeFromEnd {
	return &FindModularNodeFromEnd{
		head: ConvertArrayToSingleLinkedList(val),
		k:    k,
	}
}

func (obj *FindModularNodeFromEnd) FindModularNodeFromEnd() int {
	var modularNode *SingleLinkedList
	if obj.k <= 0 {
		return -1
	}
	temp := obj.head

	for i := 0; i < obj.k; i++ {
		if temp != nil {
			temp = temp.GetNext()
		} else {
			return -1
		}
	}

	modularNode = obj.head

	for temp != nil {
		temp = temp.GetNext()
		modularNode = modularNode.GetNext()
	}
	return modularNode.GetData()
}

func (obj *FindModularNodeFromEnd) Execute() {
	obj.Result = obj.FindModularNodeFromEnd()
}
