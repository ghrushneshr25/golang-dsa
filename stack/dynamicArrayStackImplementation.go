package stack

import (
	"fmt"
	"math"
)

const (
	DefaultDynamicCapacity = 16
	MaxDynamicCapacity     = 1 << 15
)

type DynamicArrayStack struct {
	stack    []int
	top      int
	capacity int
}

func NewDynamicArrayStack() *DynamicArrayStack {
	return &DynamicArrayStack{
		stack:    make([]int, DefaultDynamicCapacity),
		top:      -1,
		capacity: DefaultDynamicCapacity,
	}
}

func NewDynamicArrayStackWithCapacity(cap int) *DynamicArrayStack {
	return &DynamicArrayStack{
		stack:    make([]int, cap),
		top:      -1,
		capacity: cap,
	}
}

func (s *DynamicArrayStack) IsEmpty() bool {
	return s.top < 0
}

func (s *DynamicArrayStack) Size() int {
	return s.top + 1
}

func (s *DynamicArrayStack) Push(value int) {
	if s.Size() == s.capacity {
		s.expand()
	}
	s.top++
	s.stack[s.top] = value
}

func (s *DynamicArrayStack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrStackEmpty
	}
	value := s.stack[s.top]
	s.stack[s.top] = math.MinInt
	s.top--
	if s.Size() <= s.capacity/4 {
		s.shrink()
	}
	return value, nil
}

func (s *DynamicArrayStack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, ErrStackEmpty
	}
	return s.stack[s.top], nil
}

func (s *DynamicArrayStack) Clear() {
	for i := 0; i < s.top; i++ {
		s.stack[i] = math.MinInt
	}
	s.top = -1
}

func (s *DynamicArrayStack) ToString() string {
	if s.IsEmpty() {
		return "[]"
	}
	str := fmt.Sprintf("%+v", s.stack[0])
	for i := 1; i <= s.top; i++ {
		str = fmt.Sprintf("%+v , %+v", str, s.stack[i])
	}
	return fmt.Sprintf("[%+v]", str)
}

func (s *DynamicArrayStack) expand() {
	newCapacity := s.capacity << 1
	if newCapacity > MaxDynamicCapacity {
		newCapacity = MaxDynamicCapacity
	}

	newStack := make([]int, newCapacity)
	copy(newStack, s.stack)
	s.stack = newStack
	s.capacity = newCapacity
}

func (s *DynamicArrayStack) shrink() {
	newCapacity := s.capacity >> 1
	if newCapacity < DefaultDynamicCapacity {
		newCapacity = DefaultDynamicCapacity
	}
	newStack := make([]int, newCapacity)
	copy(newStack, s.stack)
	s.stack = newStack
	s.capacity = newCapacity
}
