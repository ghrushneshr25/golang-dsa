/*
Problem-59 	Given a linked list and a value K,
			partition it such that all nodes less than K come before nodes greater than or equal to K.
		   	You should preserve the original relative order of the nodes in each of the two partitions.
			For example,
				Given 1->4->3->2->5->2 and K = 3,
				return 1->2->2->4->3->5.
*/

package linkedlist

type PartitionByK struct {
	head   *SingleLinkedList
	K      int
	Result []int
}

func NewPartitionByK(val []int, K int) *PartitionByK {
	return &PartitionByK{
		head: ConvertArrayToSingleLinkedList(val),
		K:    K,
	}
}

func (obj *PartitionByK) partitionByK(head *SingleLinkedList, K int) *SingleLinkedList {
	if head == nil || head.GetNext() == nil {
		return head
	}

	root := NewSingleLinkedList(0)
	pivot := NewSingleLinkedList(K)

	rootNext, pivotNext, current := root, pivot, head

	for current != nil {
		next := current.GetNext()
		if current.GetData() >= K {
			pivotNext.SetNext(current)
			pivotNext = current
			pivotNext.SetNext(nil)
		} else {
			rootNext.SetNext(current)
			rootNext = current
		}
		current = next
	}
	rootNext.SetNext(pivot.GetNext())
	return root.GetNext()
}

func (obj *PartitionByK) Execute() {
	result := obj.partitionByK(obj.head, obj.K)
	obj.Result = ConvertSingleLinkedListToArray(result)
}
