/*
	Problem-40 Exchange the adjacent elements in a link list.
	Time Complexity: O(n)
	Space Complexity: O(1)
*/

package linkedlist

type ExchangeAdjacentNodes struct {
	head   *SingleLinkedList
	Result []int
}

func NewExchangeAdjacentNodes(val []int) *ExchangeAdjacentNodes {
	return &ExchangeAdjacentNodes{
		head: ConvertArrayToSingleLinkedList(val),
	}
}

func (ean *ExchangeAdjacentNodes) ExchangeAdjacentNodes() {
	temp := &SingleLinkedList{}
	temp.SetNext(ean.head)

	previous, current := temp, ean.head
	for current != nil && current.GetNext() != nil {
		tmp := current.GetNext().GetNext()
		current.GetNext().SetNext(previous.GetNext())
		previous.SetNext(current.GetNext())
		current.SetNext(tmp)
		previous = current
		current = previous.GetNext()
	}

	ean.head = temp.GetNext()
}

func (ean *ExchangeAdjacentNodes) Execute() {
	ean.ExchangeAdjacentNodes()
	ean.Result = ConvertSingleLinkedListToArray(ean.head)
}
