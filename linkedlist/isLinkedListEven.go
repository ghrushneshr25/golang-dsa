package linkedlist

type IsLinkedListEven struct {
	Head   *SingleLinkedList
	Result bool
}

func NewIsLinkedListEven(head *SingleLinkedList) *IsLinkedListEven {
	return &IsLinkedListEven{Head: head}
}

func (ille *IsLinkedListEven) TwoPointerMethod() bool {
	slow := ille.Head
	fast := ille.Head

	for fast != nil && fast.GetNext() != nil {
		fast = fast.GetNext().GetNext()
		slow = slow.GetNext()
	}
	return fast == nil
}

func (ille *IsLinkedListEven) Execute() {
	ille.Result = ille.TwoPointerMethod()
}
