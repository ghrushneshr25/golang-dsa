package linkedlist

type SplitCircularLinkedList struct {
	head      *CircularLinkedList
	ResultOne []int
	ResultTwo []int
}

func NewSplitCircularLinkedList(values []int) *SplitCircularLinkedList {
	head := ConvertArrayToCircularLinkedList(values)
	return &SplitCircularLinkedList{head, []int{}, []int{}}
}

func (scll *SplitCircularLinkedList) SplitList(head *CircularLinkedList, resultOne, resultTwo **CircularLinkedList) {
	slowPtr, fastPtr := head, head
	if head == nil {
		return
	}

	for fastPtr.GetNext() != head && fastPtr.GetNext().GetNext() != head {
		fastPtr = fastPtr.GetNext().GetNext()
		slowPtr = slowPtr.GetNext()
	}

	if fastPtr.GetNext().GetNext() == head {
		fastPtr = fastPtr.GetNext()
	}

	resultOne = &head
	if head.GetNext() != head {
		*resultTwo = slowPtr.GetNext()
	}

	fastPtr.SetNext(slowPtr.GetNext())
	slowPtr.SetNext(head)
}

func (scll *SplitCircularLinkedList) Execute() {
	var resultOne, resultTwo *CircularLinkedList
	scll.SplitList(scll.head, &resultOne, &resultTwo)
	scll.ResultOne = ConvertCircularLinkedListToArray(resultOne)
	scll.ResultTwo = ConvertCircularLinkedListToArray(resultTwo)
}
