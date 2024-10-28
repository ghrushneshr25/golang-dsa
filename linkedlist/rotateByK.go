/*
	Problem-57
		Given a list, rotate the list to the right by k places, where k is non-negative.
		Example: 1->->2->3->4->5->NULL, k = 2
		Output: 4->5->1->2->3->NULL
*/

package linkedlist

type RotateByK struct {
	head   *SingleLinkedList
	k      int
	Result []int
}

func NewRotateByK(val []int, k int) *RotateByK {
	return &RotateByK{
		k:    k,
		head: ConvertArrayToSingleLinkedList(val),
	}
}

func (obj *RotateByK) rotateByK() *SingleLinkedList {
	if obj.head == nil || obj.head.GetNext() == nil {
		return obj.head
	}

	n := obj.k
	rotateStart, rotateEnd := obj.head, obj.head

	for n > 0 {
		n--
		rotateEnd = rotateEnd.GetNext()
		if rotateEnd == nil {
			rotateEnd = obj.head
		}
	}

	if rotateStart == rotateEnd {
		return obj.head
	}

	for rotateEnd.GetNext() != nil {
		rotateStart = rotateStart.GetNext()
		rotateEnd = rotateEnd.GetNext()
	}

	temp := rotateStart.GetNext()
	rotateEnd.SetNext(obj.head)
	rotateStart.SetNext(nil)
	return temp
}

func (obj *RotateByK) Execute() {
	result := obj.rotateByK()
	obj.Result = ConvertSingleLinkedListToArray(result)
}
