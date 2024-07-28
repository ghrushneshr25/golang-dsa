package linkedlist

import "errors"

const MIN_NODE_CAPACITY = 4

type UnrolledLinkedList struct {
	NodeCapacity int
	CurrentSize  int
	FirstNode    *ULLNode
	LastNode     *ULLNode
}

type ULLNode struct {
	Data         []int
	NextNode     *ULLNode
	PreviousNode *ULLNode
	Size         int
}

func NewUnrolledLinkedList(nodeCapacity int) (ull *UnrolledLinkedList, err error) {
	if nodeCapacity < MIN_NODE_CAPACITY {
		err = errors.New("node capacity should be greater than min node capacity")
		return
	}
	ull = &UnrolledLinkedList{
		NodeCapacity: nodeCapacity,
	}

	firstNode := &ULLNode{
		Data: make([]int, nodeCapacity),
	}
	ull.FirstNode = firstNode
	ull.LastNode = firstNode
	return
}

func (list *UnrolledLinkedList) Size() int {
	return list.CurrentSize
}

func (list *UnrolledLinkedList) IsEmpty() bool {
	return list.CurrentSize == 0
}

func (list *UnrolledLinkedList) Contains(o int) bool {
	return list.Search(o) != -1
}

func (list *UnrolledLinkedList) Add(val int) bool {
	list.insertIntoNode(list.LastNode, list.LastNode.Size, val)
	return true
}

func (list *UnrolledLinkedList) Remove(val int) bool {
	node := list.FirstNode
	for node != nil {
		for i := 0; i < node.Size; i++ {
			if node.Data[i] == val {
				list.removeFromNode(node, i)
				return true
			}
		}
		node = node.NextNode
	}
	return false
}

func (list *UnrolledLinkedList) Clear() {
	node := list.FirstNode.NextNode
	for node != nil {
		next := node.NextNode
		node.NextNode = nil
		node.PreviousNode = nil
		node.Data = nil
		node = next
	}
	list.LastNode = list.FirstNode
	list.FirstNode.Data = make([]int, list.NodeCapacity)
	list.FirstNode.Size = 0
	list.FirstNode.NextNode = nil
	list.CurrentSize = 0
}

func (list *UnrolledLinkedList) Get(index int) (int, error) {
	if index < 0 || index >= list.CurrentSize {
		return 0, errors.New("index out of bounds")
	}
	node, p := list.findNode(index)
	return node.Data[index-p], nil
}

func (list *UnrolledLinkedList) Set(index int, element int) (int, error) {
	if index < 0 || index >= list.CurrentSize {
		return 0, errors.New("index out of bounds")
	}
	node, p := list.findNode(index)
	oldElement := node.Data[index-p]
	node.Data[index-p] = element
	return oldElement, nil
}

func (list *UnrolledLinkedList) AddAt(index int, element int) error {
	if index < 0 || index > list.CurrentSize {
		return errors.New("index out of bounds")
	}
	node, p := list.findNode(index)
	list.insertIntoNode(node, index-p, element)
	return nil
}

func (list *UnrolledLinkedList) RemoveAt(index int) (int, error) {
	if index < 0 || index >= list.CurrentSize {
		return 0, errors.New("index out of bounds")
	}
	node, p := list.findNode(index)
	element := node.Data[index-p]
	list.removeFromNode(node, index-p)
	return element, nil
}

func (list *UnrolledLinkedList) findNode(index int) (*ULLNode, int) {
	node := list.FirstNode
	p := 0
	if list.CurrentSize-index > index {
		for p <= index-node.Size {
			p += node.Size
			node = node.NextNode
		}
	} else {
		node = list.LastNode
		p = list.CurrentSize
		for p -= node.Size; p > index; {
			node = node.PreviousNode
		}
	}
	return node, p
}

func (list *UnrolledLinkedList) insertIntoNode(node *ULLNode, ptr int, element int) {
	if node.Size == list.NodeCapacity {
		newNode := &ULLNode{
			Data: make([]int, list.NodeCapacity),
		}
		elementsToMove := list.NodeCapacity / 2
		startIndex := list.NodeCapacity - elementsToMove
		copy(newNode.Data, node.Data[startIndex:])
		for i := startIndex; i < list.NodeCapacity; i++ {
			node.Data[i] = 0
		}
		node.Size -= elementsToMove
		newNode.Size = elementsToMove
		newNode.NextNode = node.NextNode
		newNode.PreviousNode = node
		if node.NextNode != nil {
			node.NextNode.PreviousNode = newNode
		}
		node.NextNode = newNode
		if node == list.LastNode {
			list.LastNode = newNode
		}
		if ptr > node.Size {
			node = newNode
			ptr -= node.Size
		}
	}
	copy(node.Data[ptr+1:], node.Data[ptr:node.Size])
	node.Data[ptr] = element
	node.Size++
	list.CurrentSize++
}

func (list *UnrolledLinkedList) removeFromNode(node *ULLNode, ptr int) {
	copy(node.Data[ptr:], node.Data[ptr+1:node.Size])
	node.Data[node.Size-1] = 0
	node.Size--
	if node.NextNode != nil && node.NextNode.Size+node.Size <= list.NodeCapacity {
		list.mergeWithNextNode(node)
	} else if node.PreviousNode != nil && node.PreviousNode.Size+node.Size <= list.NodeCapacity {
		list.mergeWithNextNode(node.PreviousNode)
	}
	list.CurrentSize--
}

func (list *UnrolledLinkedList) mergeWithNextNode(node *ULLNode) {
	next := node.NextNode
	copy(node.Data[node.Size:], next.Data[:next.Size])
	node.Size += next.Size
	if next.NextNode != nil {
		next.NextNode.PreviousNode = node
	}
	node.NextNode = next.NextNode
	if next == list.LastNode {
		list.LastNode = node
	}
}

func (list *UnrolledLinkedList) Search(o int) int {
	node := list.FirstNode
	index := 0
	for node != nil {
		for i := 0; i < node.Size; i++ {
			if node.Data[i] == o {
				return index + i
			}
		}
		index += node.Size
		node = node.NextNode
	}
	return -1
}
