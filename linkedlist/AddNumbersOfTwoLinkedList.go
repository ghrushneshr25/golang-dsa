/*
	Problem-58
		You are given two linked lists representing two non-negative numbers.
		The digits are stored in reverse order and each of their nodes contain a single digit.
		Add the two numbers and return it as a linked list.
		Example:
			3->4->3 + 5->6->4 = 8->0->8

*/

package linkedlist

type AddNumbersOfTwoLinkedList struct {
	headOne *SingleLinkedList
	headTwo *SingleLinkedList
	Result  []int
}

func NewAddNumbersOfTwoLinkedList(valOne, valTwo []int) *AddNumbersOfTwoLinkedList {
	return &AddNumbersOfTwoLinkedList{
		headOne: ConvertArrayToSingleLinkedList(valOne),
		headTwo: ConvertArrayToSingleLinkedList(valTwo),
	}
}

func (obj *AddNumbersOfTwoLinkedList) addNumbersOfTwoLinkedList(headOne, headTwo *SingleLinkedList) *SingleLinkedList {
	if headOne == nil {
		return headTwo
	}

	if headTwo == nil {
		return headOne
	}

	result := &SingleLinkedList{}
	current := result

	carryForward := 0

	for headOne != nil && headTwo != nil {
		sum := headOne.GetData() + headTwo.GetData() + carryForward
		carryForward = sum / 10
		sum = sum % 10
		current.SetNext(NewSingleLinkedList(sum))
		current = current.GetNext()
		headOne = headOne.GetNext()
		headTwo = headTwo.GetNext()
	}

	if headOne != nil {
		if carryForward != 0 {
			current.SetNext(obj.addNumbersOfTwoLinkedList(headOne, NewSingleLinkedList(carryForward)))
		} else {
			current.SetNext(headOne)
		}
	} else if headTwo != nil {
		if carryForward != 0 {
			current.SetNext(obj.addNumbersOfTwoLinkedList(headTwo, NewSingleLinkedList(carryForward)))
		} else {
			current.SetNext(headTwo)
		}
	} else if carryForward != 0 {
		current.SetNext(NewSingleLinkedList(carryForward))
	}

	return result.GetNext()
}

func (obj *AddNumbersOfTwoLinkedList) Execute() {
	result := obj.addNumbersOfTwoLinkedList(obj.headOne, obj.headTwo)
	obj.Result = ConvertSingleLinkedListToArray(result)
}
