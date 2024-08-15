package linkedlist

import "fmt"

type CircularLinkedList struct {
	data int
	next *CircularLinkedList
}

func NewCircularLinkedList(data int, next *CircularLinkedList) *CircularLinkedList {
	return &CircularLinkedList{
		data: data,
		next: next,
	}
}

func (cll *CircularLinkedList) GetData() int {
	return cll.data
}

func (cll *CircularLinkedList) SetData(data int) {
	cll.data = data
}

func (cll *CircularLinkedList) GetNext() *CircularLinkedList {
	return cll.next
}

func (cll *CircularLinkedList) SetNext(next *CircularLinkedList) {
	cll.next = next
}

func GetCircularLinkedListLength(cll *CircularLinkedList) int {
	length := 0
	if cll == nil {
		return length
	}
	length = 1
	temp := cll.GetNext()
	for ; temp != cll; temp = temp.GetNext() {
		length = length + 1
	}
	return length
}

func PrintContentOfCircularLinkedList(cll *CircularLinkedList) (data string) {
	if cll == nil {
		return
	}
	data = fmt.Sprintf("%v", cll.GetData())
	temp := cll.GetNext()
	for temp != cll {
		data = fmt.Sprintf("%v %v", data, temp.GetData())
		temp = temp.GetNext()
	}
	return
}

func InsertAtEndOfCircularLinkedList(head **CircularLinkedList, data int) {
	newNode := NewCircularLinkedList(data, nil)
	newNode.SetNext(newNode)
	if *head == nil {
		*head = newNode
		return
	}

	temp := *head
	for temp.GetNext() != *head {
		temp = temp.GetNext()
	}
	temp.SetNext(newNode)
	newNode.SetNext(*head)
}

func InsertAtFrontOfCircularLinkedList(head **CircularLinkedList, data int) {
	newNode := NewCircularLinkedList(data, nil)
	newNode.SetNext(newNode)
	if *head == nil {
		*head = newNode
		return
	}
	temp := *head
	for temp.GetNext() != *head {
		temp = temp.GetNext()
	}
	temp.SetNext(newNode)
	newNode.SetNext(*head)
	*head = newNode
}

func DeleteLastNodeOfCircularLinkedList(head **CircularLinkedList) {
	if *head == nil {
		return
	}
	temp := *head
	if temp.GetNext() == *head {
		*head = nil
		return
	}

	for temp.GetNext().GetNext() != *head {
		temp = temp.GetNext()
	}
	temp.SetNext(*head)
}

func DeleteFirstNodeOfCircularLinkedList(head **CircularLinkedList) {
	if *head == nil {
		return
	}
	temp := *head
	if temp.GetNext() == *head {
		*head = nil
		return
	}

	for temp.GetNext() != *head {
		temp = temp.GetNext()
	}
	temp.SetNext((*head).GetNext())
	*head = (*head).GetNext()
}

func ConvertArrayToCircularLinkedList(arr []int) *CircularLinkedList {
	var head *CircularLinkedList
	for _, v := range arr {
		InsertAtEndOfCircularLinkedList(&head, v)
	}
	return head
}

func ConvertCircularLinkedListToArray(head *CircularLinkedList) []int {
	var arr []int
	if head == nil {
		return arr
	}
	temp := head
	for {
		arr = append(arr, temp.GetData())
		temp = temp.GetNext()
		if temp == head {
			break
		}
	}
	return arr
}
