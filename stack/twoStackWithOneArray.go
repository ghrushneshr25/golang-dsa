package stack

import "errors"

var (
	ErrInvalidStackId = errors.New("InvalidStackId")
)

type TwoStackWithOneArray struct {
	stack  []int
	size   int
	topOne int
	topTwo int
}

func NewTwoStackWithOneArray(size int) *TwoStackWithOneArray {
	return &TwoStackWithOneArray{
		stack:  make([]int, size),
		size:   size,
		topOne: -1,
		topTwo: size,
	}
}

func (s *TwoStackWithOneArray) Push(stackId, value int) error {
	if s.topOne+1 == s.topTwo {
		return ErrStackFull
	}
	if stackId == 1 {
		s.topOne++
		s.stack[s.topOne] = value
	} else if stackId == 2 {
		s.topTwo--
		s.stack[s.topTwo] = value
	} else {
		return ErrInvalidStackId
	}
	return nil
}

func (s *TwoStackWithOneArray) Pop(stackId int) (int, error) {
	if stackId == 1 {
		if s.topOne == -1 {
			return 0, ErrStackEmpty
		}
		value := s.stack[s.topOne]
		s.topOne--
		return value, nil
	} else if stackId == 2 {
		if s.topTwo == s.size {
			return 0, ErrStackEmpty
		}
		value := s.stack[s.topTwo]
		s.topTwo++
		return value, nil
	}
	return 0, ErrInvalidStackId
}

func (s *TwoStackWithOneArray) Peek(stackId int) (int, error) {
	if stackId == 1 {
		if s.topOne == -1 {
			return 0, ErrStackEmpty
		}
		return s.stack[s.topOne], nil
	} else if stackId == 2 {
		if s.topTwo == s.size {
			return 0, ErrStackEmpty
		}
		return s.stack[s.topTwo], nil
	}
	return 0, ErrInvalidStackId
}

func (s *TwoStackWithOneArray) IsEmpty(stackId int) bool {
	if stackId == 1 {
		return s.topOne == -1
	}
	return s.topTwo == s.size
}
