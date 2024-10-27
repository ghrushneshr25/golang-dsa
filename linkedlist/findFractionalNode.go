/*
Problem-48
	Find fractional node:
	Given a singly linked list, write a function to find the n/k th element, where n is the number of elements in the list.
*/

package linkedlist

type FindFractionalNode struct {
	head   *SingleLinkedList
	k      int
	Result int
}

func NewFindFractionalNode(val []int, k int) *FindFractionalNode {
	return &FindFractionalNode{
		head: ConvertArrayToSingleLinkedList(val),
		k:    k,
	}
}

func (obj *FindFractionalNode) FindFractionalNode() int {
	var fractionalNode *SingleLinkedList
	if obj.k <= 0 {
		return -1
	}
	i := 0
	temp := obj.head
	for temp != nil {
		if i%obj.k == 0 {
			if fractionalNode == nil {
				fractionalNode = obj.head
			} else {
				fractionalNode = fractionalNode.GetNext()
			}
		}
		i++
		temp = temp.GetNext()
		if temp == nil {
			if i%obj.k == 0 && fractionalNode != nil {
				fractionalNode = fractionalNode.GetNext()
			}
		}
	}

	return fractionalNode.GetData()
}

func (obj *FindFractionalNode) Execute() {
	obj.Result = obj.FindFractionalNode()
}
