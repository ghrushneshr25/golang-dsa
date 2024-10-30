package stack

import (
	"errors"
	"fmt"
	"math"
)

const (
	DefaultCapacity = 10
)

var (
	ErrStackEmpty = errors.New("KeyStackEmpty")
	ErrStackFull  = errors.New("KeyStackFull")
)

type FixedSizeArrayStack struct {
	stack    []int
	capacity int
	top      int
}

func NewFixedSizeArrayStack() *FixedSizeArrayStack {
	return &FixedSizeArrayStack{
		stack:    make([]int, DefaultCapacity),
		capacity: DefaultCapacity,
		top:      -1,
	}
}

func NewFixedSizeArrayStackWithCapacity(capacity int) *FixedSizeArrayStack {
	return &FixedSizeArrayStack{
		stack:    make([]int, capacity),
		capacity: capacity,
		top:      -1,
	}
}

func (obj *FixedSizeArrayStack) Size() int {
	return obj.top + 1
}

func (obj *FixedSizeArrayStack) IsEmpty() bool {
	return obj.top < 0
}

func (obj *FixedSizeArrayStack) IsFull() bool {
	return obj.Size() == obj.capacity
}

func (obj *FixedSizeArrayStack) Push(val int) (err error) {
	if obj.IsFull() {
		return ErrStackFull
	}
	obj.top++
	obj.stack[obj.top] = val
	return
}

func (obj *FixedSizeArrayStack) Top() (int, error) {
	if obj.IsEmpty() {
		return 0, ErrStackEmpty
	}
	return obj.stack[obj.top], nil
}

func (obj *FixedSizeArrayStack) Pop() (int, error) {
	if obj.IsEmpty() {
		return 0, ErrStackEmpty
	}
	val := obj.stack[obj.top]
	obj.stack[obj.top] = math.MinInt
	obj.top--
	return val, nil
}

func (obj *FixedSizeArrayStack) ToString() string {
	if obj.IsEmpty() {
		return "[]"
	}
	str := fmt.Sprintf("%+v", obj.stack[0])
	for i := 1; i <= obj.top; i++ {
		str = fmt.Sprintf("%+v , %+v", str, obj.stack[i])
	}
	return fmt.Sprintf("[%+v]", str)
}
