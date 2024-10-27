/*
Problem-54
	Given a singly linked list L : L1->L2->L3->L4->L5->....->Ln-1->Ln
	Convert it to L1->Ln->L2->Ln-2......
*/

package linkedlist

type ReorderLinkedList struct {
	head   *SingleLinkedList
	Result []int
}

func NewReorderLinkedList(val []int) *ReorderLinkedList {
	return &ReorderLinkedList{
		head: ConvertArrayToSingleLinkedList(val),
	}
}

func (obj *ReorderLinkedList) ReorderLinkedList() {
	if obj.head == nil {
		return
	}
	slowPointer := obj.head
	fastPointer := obj.head.GetNext()

	for fastPointer != nil && fastPointer.GetNext() != nil {
		fastPointer = fastPointer.GetNext().GetNext()
		slowPointer = slowPointer.GetNext()
	}

	head2 := slowPointer.GetNext()
	slowPointer.SetNext(nil)

	// reverse second half

	head2 = obj.reverse(head2)
	obj.alternate(obj.head, head2)
}

func (obj *ReorderLinkedList) reverse(head *SingleLinkedList) *SingleLinkedList {
	if head == nil {
		return nil
	}

	reversedList := head
	pointer := head.GetNext()

	reversedList.SetNext(nil)

	for pointer != nil {
		temp := pointer
		pointer = pointer.GetNext()
		temp.SetNext(reversedList)
		reversedList = temp
	}
	return reversedList
}

func (obj *ReorderLinkedList) alternate(head1, head2 *SingleLinkedList) {
	pointer := head1
	head1 = head1.GetNext()

	nextList := false

	for head1 != nil && head2 != nil {
		if (head2 != nil && !nextList) || head1 == nil {
			pointer.SetNext(head2)
			head2 = head2.GetNext()
			nextList = true
			pointer = pointer.GetNext()
		} else {
			pointer.SetNext(head1)
			head1 = head1.GetNext()
			nextList = false
			pointer = pointer.GetNext()
		}
	}
	if head1 != nil {
		pointer.SetNext(head1)
	} else if head2 != nil {
		pointer.SetNext(head2)
	}
}

func (obj *ReorderLinkedList) Execute() {
	obj.ReorderLinkedList()
	obj.Result = ConvertSingleLinkedListToArray(obj.head)
}
