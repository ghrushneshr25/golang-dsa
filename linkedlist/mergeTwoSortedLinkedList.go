package linkedlist

type MergeTwoSortedLinkedList struct {
	headOne    *SingleLinkedList
	headTwo    *SingleLinkedList
	mergedHead *SingleLinkedList
	Result     []int
}

func NewMergeTwoSortedLinkedList(headOne, headTwo *SingleLinkedList) *MergeTwoSortedLinkedList {
	return &MergeTwoSortedLinkedList{
		headOne: headOne,
		headTwo: headTwo,
	}
}

func (mtsll *MergeTwoSortedLinkedList) recursion(headOne, headTwo *SingleLinkedList) *SingleLinkedList {
	if headOne == nil {
		return headTwo
	}
	if headTwo == nil {
		return headOne
	}
	var head *SingleLinkedList
	if headOne.GetData() < headTwo.GetData() {
		head = headOne
		head.SetNext(mtsll.recursion(headOne.GetNext(), headTwo))
	} else {
		head = headTwo
		head.SetNext(mtsll.recursion(headOne, headTwo.GetNext()))
	}
	return head
}

func (mtsll *MergeTwoSortedLinkedList) RecursionMethod() {
	mtsll.mergedHead = mtsll.recursion(mtsll.headOne, mtsll.headTwo)
}

func (mtsll *MergeTwoSortedLinkedList) IterativeMethod() {
	head := &SingleLinkedList{}
	current := head

	for mtsll.headOne != nil && mtsll.headTwo != nil {
		if mtsll.headOne.GetData() < mtsll.headTwo.GetData() {
			current.SetNext(mtsll.headOne)
			mtsll.headOne = mtsll.headOne.GetNext()
		} else {
			current.SetNext(mtsll.headTwo)
			mtsll.headTwo = mtsll.headTwo.GetNext()
		}
		current = current.GetNext()
	}

	if mtsll.headOne != nil {
		current.SetNext(mtsll.headOne)
	} else if mtsll.headTwo != nil {
		current.SetNext(mtsll.headTwo)
	}
	mtsll.mergedHead = head.GetNext()
}

func (mtsll *MergeTwoSortedLinkedList) ExecuteByMethodName(method string) {
	switch method {
	case "RecursionMethod":
		mtsll.RecursionMethod()
	case "IterativeMethod":
		mtsll.IterativeMethod()
	}
	mtsll.Result = ConvertSingleLinkedListToArray(mtsll.mergedHead)
}
