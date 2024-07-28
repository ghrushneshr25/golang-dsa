/*
	Problem-7: Check whether the given linked list is either NULL-terminated or ends in a cycle (cyclic)
	Problem 8: --------------""-----------------------
	Problem 9: --------------""-----------------------
	Problem 10:  ------------""-----------------------
	Problem 11:  We are given a pointer to the first element of a linked list L.
				There are two possibilities for L it either ends (snake) or 
				its last element points back to one of the earlier elements in the list (snail). 
				Give an algorithm that tests whether a given list is a snake or a snail.
*/

package linkedlist

type CheckCycle struct {
	Head   *SingleLinkedList
	Result bool
}

func NewCheckCycle(head *SingleLinkedList) *CheckCycle {
	return &CheckCycle{
		Head: head,
	}
}

func (cc *CheckCycle) Execute(methodName string) {
	methods := map[string]func() bool{
		"HashMapMethod":    cc.HashMapMethod,
		"TwoPointerMethod": cc.TwoPointerMethod,
	}
	cc.Result = methods[methodName]()
}

/*
	Algorithm:
		- Traverse the linked list nodes one by one.
		- Check if the address of the node is available in the hash table or not.
		- If it is already available in the hash table, that indicates that we are visiting the node that was already visited. This is possible only if
			the given linked list has a loop in it.
		- If the address of the node is not available in the hash table, insert that node’s address into the hash table.
		- Continue this process until we reach the end of the linked list ݎ݋ we find the loop.

	Time Complexity: O(n) for scanning the linked list
	Space Complexity: O(n)for hash table.
*/

func (cc *CheckCycle) HashMapMethod() bool {
	hashMap := map[*SingleLinkedList]bool{}
	current := cc.Head
	for current != nil {
		if _, ok := hashMap[current]; ok {
			return true
		}
		hashMap[current] = true
		current = current.GetNext()
	}
	return false
}

func (cc *CheckCycle) TwoPointerMethod() bool {
	fastPtr := cc.Head
	slowPtr := cc.Head

	for fastPtr != nil && fastPtr.GetNext() != nil {
		fastPtr = fastPtr.GetNext().GetNext()
		slowPtr = slowPtr.GetNext()
		if fastPtr == slowPtr {
			return true
		}
	}
	return false
}
