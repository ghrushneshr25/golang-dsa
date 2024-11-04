/*
	Problem-29 	If the stack gets too high, it might overbalance.
				There-fore, in real life, we would likely start a new stack when the previous
				stack exceeds some threshold.  Implement a data structure that mimics this
				and composed of several stacks, and should create a new stack
				once the previous one exceeds capacity.
				push() and pop() of this class should behave identically to a regular stack.
*/

package stack

type StackForStackSets struct {
	stacks   []*FixedSizeArrayStack
	capacity int
}

func NewStackForStackSets(capacity int) *StackForStackSets {
	newStack := NewFixedSizeArrayStackWithCapacity(capacity)
	return &StackForStackSets{
		stacks:   []*FixedSizeArrayStack{newStack},
		capacity: capacity,
	}
}

func (s *StackForStackSets) getLastStack() *FixedSizeArrayStack {
	return s.stacks[len(s.stacks)-1]
}

func (s *StackForStackSets) Push(value int) error {
	lastStack := s.getLastStack()
	if lastStack.IsFull() {
		newStack := NewFixedSizeArrayStackWithCapacity(s.capacity)
		newStack.Push(value)
		s.stacks = append(s.stacks, newStack)
	} else {
		lastStack.Push(value)
	}
	return nil
}

func (s *StackForStackSets) Pop() (int, error) {
	lastStack := s.getLastStack()
	value, err := lastStack.Pop()
	if err != nil {
		return 0, err
	}
	if lastStack.IsEmpty() && len(s.stacks) > 1 {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}

	return value, nil
}

func (s *StackForStackSets) getNthStack(n int) *FixedSizeArrayStack {
	if n < 0 || n >= len(s.stacks) {
		return nil
	}
	return s.stacks[n]
}

func (s *StackForStackSets) PopFromNth(n int) (int, error) {
	stack := s.getNthStack(n)
	if stack == nil {
		return 0, ErrInvalidStackId
	}
	value, _ := stack.Pop()
	if stack.IsEmpty() {
		s.stacks = append(s.stacks[:n], s.stacks[n+1:]...)
	}
	return value, nil
}
