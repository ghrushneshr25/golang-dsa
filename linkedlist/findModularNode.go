/*
Problem-47
	Find modular node: Given a singly linked list, write a function to find the last element from the beginning whose n % k == 0,
	where n is the number of elements in the list and k is an integer constant. For example, if n = 19 and k = 3 then we should return
	18th node.

	Time Complexity  : O(n)
	Space Complexity : O(1)
*/

package linkedlist

type FindModularNode struct {
	head   *SingleLinkedList
	k      int
	Result int
}

func NewFindModularNode(val []int, k int) *FindModularNode {
	return &FindModularNode{
		head: ConvertArrayToSingleLinkedList(val),
		k:    k,
	}
}

func (obj *FindModularNode) FindModularNode() int {
	var modularNode *SingleLinkedList
	if obj.k <= 0 {
		return -1
	}
	i := 0
	temp := obj.head
	for temp != nil {
		if i%obj.k == 0 {
			modularNode = temp
		}
		i++
		temp = temp.GetNext()
	}
	return modularNode.GetData()
}

func (obj *FindModularNode) Execute() {
	obj.Result = obj.FindModularNode()
}
