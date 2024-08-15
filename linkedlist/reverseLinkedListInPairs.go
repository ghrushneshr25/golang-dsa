package linkedlist

type ReverseLinkedListInPairs struct {
	Head        *SingleLinkedList
	ReverseHead *SingleLinkedList
	Result      []int
}

func NewReverseLinkedListInPairs(values []int) *ReverseLinkedListInPairs {
	head := ConvertArrayToSingleLinkedList(values)
	return &ReverseLinkedListInPairs{
		Head: head,
	}
}

func (rllip *ReverseLinkedListInPairs) RecursiveMethod() {
	rllip.ReverseHead = rllip.recursion(rllip.Head)
}

func (rllp *ReverseLinkedListInPairs) recursion(head *SingleLinkedList) *SingleLinkedList {
	var temp *SingleLinkedList
	if head == nil || head.GetNext() == nil {
		return head
	}
	temp = head.GetNext()
	head.SetNext(temp.GetNext())
	temp.SetNext(head)
	head = temp

	head.GetNext().SetNext(rllp.recursion(head.GetNext().GetNext()))
	return head
}

func (rllip *ReverseLinkedListInPairs) IterativeMethod() {
	var tempOne, tempTwo *SingleLinkedList
	head := rllip.Head
	for head != nil && head.GetNext() != nil {
		if tempOne != nil {
			tempOne.GetNext().SetNext(head.GetNext())
		}
		tempOne = head.GetNext()
		head.SetNext(head.GetNext().GetNext())
		tempOne.SetNext(head)
		if tempTwo == nil {
			tempTwo = tempOne
		}
		head = head.GetNext()
	}
	rllip.ReverseHead = tempTwo
}

func (rllip *ReverseLinkedListInPairs) ExecuteByMethodName(methodName string) {
	switch methodName {
	case "RecursiveMethod":
		rllip.RecursiveMethod()
	case "IterativeMethod":
		rllip.IterativeMethod()
	}
	rllip.Result = ConvertSingleLinkedListToArray(rllip.ReverseHead)
}
