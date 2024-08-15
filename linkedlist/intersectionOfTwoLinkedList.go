package linkedlist

type IntersectionOfTwoLinkedList struct {
	headOne *SingleLinkedList
	headTwo *SingleLinkedList
	Result  int
}

func NewIntersectionOfTwoLinkedList(headOne, headTwo *SingleLinkedList) *IntersectionOfTwoLinkedList {
	return &IntersectionOfTwoLinkedList{
		headOne: headOne,
		headTwo: headTwo,
		Result:  -1,
	}
}

func (i2ll *IntersectionOfTwoLinkedList) BruteForceMethod() *SingleLinkedList {
	for headOne := i2ll.headOne; headOne != nil; headOne = headOne.GetNext() {
		for headTwo := i2ll.headTwo; headTwo != nil; headTwo = headTwo.GetNext() {
			if headOne == headTwo {
				return headOne
			}
		}
	}
	return nil
}

func (i2ll *IntersectionOfTwoLinkedList) StackMethod() *SingleLinkedList {
	stack1 := []*SingleLinkedList{}
	stack2 := []*SingleLinkedList{}

	for headOne := i2ll.headOne; headOne != nil; headOne = headOne.GetNext() {
		stack1 = append([]*SingleLinkedList{headOne}, stack1...)
	}

	for headTwo := i2ll.headTwo; headTwo != nil; headTwo = headTwo.GetNext() {
		stack2 = append([]*SingleLinkedList{headTwo}, stack2...)
	}

	for i := 0; i < len(stack1) && i < len(stack2); i++ {
		if stack1[i] != stack2[i] {
			if i == 0 {
				return nil
			}
			return stack1[i-1]
		}
	}
	return nil
}

func (i2ll *IntersectionOfTwoLinkedList) TwoPointerMethod() *SingleLinkedList {
	var len1, len2, diff int
	currentOne, currentTwo := i2ll.headOne, i2ll.headTwo

	for currentOne != nil {
		len1++
		currentOne = currentOne.GetNext()
	}

	for currentTwo != nil {
		len2++
		currentTwo = currentTwo.GetNext()
	}

	if len2 > len1 {
		currentOne, currentTwo = i2ll.headTwo, i2ll.headOne
		diff = len2 - len1
	} else {
		currentOne, currentTwo = i2ll.headOne, i2ll.headTwo
		diff = len1 - len2
	}

	for i := 0; i < diff; i++ {
		currentOne = currentOne.GetNext()
	}

	for currentOne != nil && currentTwo != nil {
		if currentOne == currentTwo {
			return currentOne
		}
		currentOne = currentOne.GetNext()
		currentTwo = currentTwo.GetNext()
	}
	return nil
}

func (i2ll *IntersectionOfTwoLinkedList) ExecuteByName(methodName string) {
	methods := map[string]func() *SingleLinkedList{
		"BruteForceMethod": i2ll.BruteForceMethod,
		"StackMethod":      i2ll.StackMethod,
		"TwoPointerMethod": i2ll.TwoPointerMethod,
	}

	result := methods[methodName]()
	if result != nil {
		i2ll.Result = result.GetData()
	}
}
