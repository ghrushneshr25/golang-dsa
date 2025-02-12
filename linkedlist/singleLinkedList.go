package linkedlist

type SingleLinkedList struct {
	CustomData interface{}
	data       int
	nextNode   *SingleLinkedList
}

func NewSingleLinkedList(data int) *SingleLinkedList {
	return &SingleLinkedList{
		data: data,
	}
}

func (sll *SingleLinkedList) SetData(data int) {
	sll.data = data
}

func (sll *SingleLinkedList) GetData() int {
	return sll.data
}

func (sll *SingleLinkedList) SetNext(nextNode *SingleLinkedList) {
	sll.nextNode = nextNode
}

func (sll *SingleLinkedList) GetNext() *SingleLinkedList {
	return sll.nextNode
}

/*
	Time Complexity: O(n), for scanning the list of size n Space Complexity: O(1), for creating a temporary variable.
*/

func GetSingleLinkedListLength(head *SingleLinkedList) int {
	var length int
	currentNode := head
	for currentNode != nil {
		length++
		currentNode = currentNode.GetNext()
	}
	return length
}

func SearchInSingleLL(head *SingleLinkedList, data int) *SingleLinkedList {
	currentNode := head
	for currentNode != nil {
		if currentNode.GetData() == data {
			return currentNode
		}
		currentNode = currentNode.GetNext()
	}
	return nil
}

func InsertAtStartofSingleLL(head **SingleLinkedList, data int) {
	newNode := NewSingleLinkedList(data)
	newNode.SetNext(*head)
	*head = newNode
}

func InsertAtEndofSingleLL(head **SingleLinkedList, data int) {
	newNode := NewSingleLinkedList(data)
	if *head == nil {
		*head = newNode
		return
	}
	currentNode := *head
	for currentNode.GetNext() != nil {
		currentNode = currentNode.GetNext()
	}
	currentNode.SetNext(newNode)
}

func InsertAtPositionOfSingleLL(head **SingleLinkedList, data, position int) {
	if position == 0 {
		InsertAtStartofSingleLL(head, data)
		return
	}

	newNode := NewSingleLinkedList(data)

	currentNode := *head
	for i := 0; i < position-1 && currentNode != nil; i++ {
		currentNode = currentNode.GetNext()
	}

	if currentNode == nil {
		return // out of range
	}
	newNode.SetNext(currentNode.GetNext())
	currentNode.SetNext(newNode)
}

func DeleteAtStartOfSingleLL(head **SingleLinkedList) {
	if *head != nil {
		*head = (*head).GetNext()
	}
}

func DeleteAtEndSingleLL(head **SingleLinkedList) {
	if *head == nil {
		return
	}
	if (*head).GetNext() == nil {
		*head = nil
		return
	}
	currentNode := *head
	for currentNode.GetNext().GetNext() != nil {
		currentNode = currentNode.GetNext()
	}
	currentNode.SetNext(nil)
}

func DeleteAtPositionOfSingleLL(head **SingleLinkedList, position int) {
	if position == 0 {
		DeleteAtStartOfSingleLL(head)
		return
	}
	currentNode := *head
	for i := 0; i < position-1 && currentNode != nil; i++ {
		currentNode = currentNode.GetNext()
	}
	if currentNode == nil || currentNode.GetNext() == nil {
		return // position is out of bounds
	}
	currentNode.SetNext(currentNode.GetNext().GetNext())
}

func ConvertArrayToSingleLinkedList(val []int) *SingleLinkedList {
	if len(val) == 0 {
		return nil
	}
	var head *SingleLinkedList
	for i := len(val) - 1; i >= 0; i-- {
		InsertAtStartofSingleLL(&head, val[i])
	}
	return head
}

func ConvertSingleLinkedListToArray(head *SingleLinkedList) []int {
	result := []int{}
	currentNode := head
	for currentNode != nil {
		result = append(result, currentNode.GetData())
		currentNode = currentNode.GetNext()
	}
	return result
}
