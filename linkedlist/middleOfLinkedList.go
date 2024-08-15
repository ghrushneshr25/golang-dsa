package linkedlist

type MiddleOfLinkedList struct {
	Head   *SingleLinkedList
	Result int
}

func NewMiddleOfLinkedList(head *SingleLinkedList) *MiddleOfLinkedList {
	return &MiddleOfLinkedList{Head: head}
}

func (mll *MiddleOfLinkedList) BruteForceMethod() int {
	len := GetSingleLinkedListLength(mll.Head)
	mid := len / 2
	current := mll.Head

	for current != nil {
		currentLength := GetSingleLinkedListLength(current)
		if currentLength == len-mid {
			return current.GetData()
		}
		current = current.GetNext()
	}
	return -1
}

func (mll *MiddleOfLinkedList) LengthMethod() int {
	len := GetSingleLinkedListLength(mll.Head)
	if len == 0 {
		return -1
	}
	mid := len / 2
	current := mll.Head
	for i := 0; i < mid; i++ {
		current = current.GetNext()
	}
	return current.GetData()
}

func (mll *MiddleOfLinkedList) TwoPointerMethod() int {
	slow := mll.Head
	fast := mll.Head
	for fast != nil && fast.GetNext() != nil {
		fast = fast.GetNext().GetNext()
		slow = slow.GetNext()
	}
	if slow == nil {
		return -1
	}
	return slow.GetData()
}

func (mll *MiddleOfLinkedList) HashMapMethod() int {
	count := 0
	nodeMap := make(map[int]*SingleLinkedList)
	current := mll.Head

	for current != nil {
		count++
		nodeMap[count] = current
		current = current.GetNext()
	}
	if count == 0 {
		return -1
	}
	mid := count / 2
	return nodeMap[mid+1].GetData()
}

func (mll *MiddleOfLinkedList) ExecuteByMethod(method string) {
	methods := map[string]func() int{
		"BruteForceMethod": mll.BruteForceMethod,
		"LengthMethod":     mll.LengthMethod,
		"TwoPointerMethod": mll.TwoPointerMethod,
		"HashMapMethod":    mll.HashMapMethod,
	}
	mll.Result = methods[method]()
}
