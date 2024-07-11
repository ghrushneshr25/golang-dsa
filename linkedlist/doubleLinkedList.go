package linkedlist

type DoubleLinkedList struct {
	data int
	next *DoubleLinkedList
	prev *DoubleLinkedList
}

func NewDoubleLinkedList(data int, next, prev *DoubleLinkedList) *DoubleLinkedList {
	return &DoubleLinkedList{
		data: data,
		next: next,
		prev: prev,
	}
}

func (dll *DoubleLinkedList) GetData() int {
	return dll.data
}

func (dll *DoubleLinkedList) SetData(data int) {
	dll.data = data
}

func (dll *DoubleLinkedList) GetPrev() *DoubleLinkedList {
	return dll.prev
}

func (dll *DoubleLinkedList) SetPrev(prev *DoubleLinkedList) {
	dll.prev = prev
}

func (dll *DoubleLinkedList) GetNext() *DoubleLinkedList {
	return dll.next
}

func (dll *DoubleLinkedList) SetNext(next *DoubleLinkedList) {
	dll.next = next
}

func GetDoubleLinkedListLength(head *DoubleLinkedList) int {
	var length int

	temp := head
	for temp.GetNext() != nil {
		length++
		temp = temp.GetNext()
	}

	return length
}

func InsertAtStartDoubleLL(head **DoubleLinkedList, data int) {
	newNode := NewDoubleLinkedList(data, *head, nil)
	if *head != nil {
		(*head).SetPrev(newNode)
	}
	*head = newNode
}

func InsertAtEndDoubleLL(head **DoubleLinkedList, data int) {
	newNode := NewDoubleLinkedList(data, nil, nil)
	if *head == nil {
		*head = newNode
		return
	}

	temp := *head

	for temp.GetNext() != nil {
		temp = temp.GetNext()
	}
	temp.SetNext(newNode)
	newNode.SetPrev(temp)
}

func InsertAtPositionOfDoubleLL(head **DoubleLinkedList, data, position int) {
	if position == 0 {
		InsertAtStartDoubleLL(head, data)
		return
	}

	newNode := NewDoubleLinkedList(data, nil, nil)
	temp := *head

	for i := 0; i < position-1 && temp != nil; i++ {
		temp = temp.GetNext()
	}

	if temp == nil {
		return
	}

	newNode.SetNext(temp.GetNext())
	newNode.SetPrev(temp)
	if temp.GetNext() != nil {
		temp.GetNext().SetPrev(newNode)
	}
	temp.SetNext(newNode)
}

func DeleteAtStartOfDoubleLL(head **DoubleLinkedList) {
	if *head == nil {
		return
	}
	*head = (*head).GetNext()
	if *head != nil {
		(*head).SetPrev(nil)
	}
}

func DeleteAtEndOfDoubleLL(head **DoubleLinkedList) {
	if *head == nil {
		return
	}

	temp := *head
	for temp.GetNext() != nil {
		temp = temp.GetNext()
	}

	if temp.GetPrev() != nil {
		temp.GetPrev().SetNext(nil)
	} else {
		*head = nil
	}
}

func DeleteAtPositionOfDoubleLL(head **DoubleLinkedList, position int) {
	if position == 0 {
		DeleteAtStartOfDoubleLL(head)
		return
	}

	temp := *head
	for i := 0; i < position && temp != nil; i++ {
		temp = temp.GetNext()
	}

	if temp == nil {
		return
	}

	if temp.GetPrev() != nil {
		temp.GetPrev().SetNext(temp.GetNext())
	}
	if temp.GetNext() != nil {
		temp.GetNext().SetPrev(temp.GetPrev())
	}
}
