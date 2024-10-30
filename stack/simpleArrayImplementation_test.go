package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FixedSizeArrayStackTestSuite struct {
	suite.Suite
	stack *FixedSizeArrayStack
}

func TestFixedSizeArrayStack(t *testing.T) {
	suite.Run(t, new(FixedSizeArrayStackTestSuite))
}

func (suite *FixedSizeArrayStackTestSuite) SetupTest() {
	suite.stack = NewFixedSizeArrayStack()
}

func (suite *FixedSizeArrayStackTestSuite) TestPush() {
	tests := []struct {
		values       []int
		expectedData []int
		expectedErr  error
	}{
		{
			values:       []int{1},
			expectedData: []int{1},
			expectedErr:  nil,
		},
		{
			values:       []int{1, 2},
			expectedData: []int{1, 2},
			expectedErr:  nil,
		},
		{
			values:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectedData: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectedErr:  nil,
		},
		{
			values:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			expectedData: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectedErr:  ErrStackFull,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Push %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.values {
				err := suite.stack.Push(value)
				if err != nil {
					suite.Equal(test.expectedErr, err)
					break
				}
			}
			for i, expected := range test.expectedData {
				suite.Equal(expected, suite.stack.stack[i])
			}
			suite.stack = NewFixedSizeArrayStack()
		})
	}
}

func (suite *FixedSizeArrayStackTestSuite) TestPop() {
	tests := []struct {
		initialValues []int
		expectedData  []int
		expectedErr   error
	}{
		{
			initialValues: []int{1, 2, 3},
			expectedData:  []int{3, 2, 1},
			expectedErr:   nil,
		},
		{
			initialValues: []int{1, 2},
			expectedData:  []int{2, 1},
			expectedErr:   nil,
		},
		{
			initialValues: []int{1},
			expectedData:  []int{1},
			expectedErr:   nil,
		},
		{
			initialValues: []int{},
			expectedData:  []int{},
			expectedErr:   ErrStackEmpty,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Pop %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.initialValues {
				suite.stack.Push(value)
			}
			for _, expected := range test.expectedData {
				val, err := suite.stack.Pop()
				if err != nil {
					suite.Equal(test.expectedErr, err)
					break
				}
				suite.Equal(expected, val)
			}
			_, err := suite.stack.Pop()
			suite.Equal(ErrStackEmpty, err)
			suite.stack = NewFixedSizeArrayStack()
		})
	}
}

func (suite *FixedSizeArrayStackTestSuite) TestTop() {
	tests := []struct {
		initialValues []int
		expectedData  int
		expectedErr   error
	}{
		{
			initialValues: []int{1, 2, 3},
			expectedData:  3,
			expectedErr:   nil,
		},
		{
			initialValues: []int{1, 2},
			expectedData:  2,
			expectedErr:   nil,
		},
		{
			initialValues: []int{1},
			expectedData:  1,
			expectedErr:   nil,
		},
		{
			initialValues: []int{},
			expectedData:  0,
			expectedErr:   ErrStackEmpty,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Top %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.initialValues {
				suite.stack.Push(value)
			}
			val, err := suite.stack.Top()
			suite.Equal(test.expectedErr, err)
			if err == nil {
				suite.Equal(test.expectedData, val)
			}
			suite.stack = NewFixedSizeArrayStack()
		})
	}
}

func (suite *FixedSizeArrayStackTestSuite) TestIsEmpty() {
	suite.True(suite.stack.IsEmpty())

	suite.stack.Push(1)
	suite.False(suite.stack.IsEmpty())

	suite.stack.Pop()
	suite.True(suite.stack.IsEmpty())
}

func (suite *FixedSizeArrayStackTestSuite) TestSize() {
	suite.Equal(0, suite.stack.Size())

	suite.stack.Push(1)
	suite.Equal(1, suite.stack.Size())

	suite.stack.Push(2)
	suite.Equal(2, suite.stack.Size())

	suite.stack.Pop()
	suite.Equal(1, suite.stack.Size())

	suite.stack.Pop()
	suite.Equal(0, suite.stack.Size())
}

func (suite *FixedSizeArrayStackTestSuite) TestToString() {
	tests := []struct {
		initialValues []int
		expectedStr   string
	}{
		{
			initialValues: []int{},
			expectedStr:   "[]",
		},
		{
			initialValues: []int{1},
			expectedStr:   "[1]",
		},
		{
			initialValues: []int{1, 2, 3},
			expectedStr:   "[1 , 2 , 3]",
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("ToString %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.initialValues {
				suite.stack.Push(value)
			}
			str := suite.stack.ToString()
			suite.Equal(test.expectedStr, str)
			suite.stack = NewFixedSizeArrayStack()
		})
	}
}

func (suite *FixedSizeArrayStackTestSuite) TestNewFixedSizeArrayStackWithCapacity() {
	tests := []struct {
		capacity      int
		expectedStack *FixedSizeArrayStack
	}{
		{
			capacity: 5,
			expectedStack: &FixedSizeArrayStack{
				stack:    make([]int, 5),
				capacity: 5,
				top:      -1,
			},
		},
		{
			capacity: 10,
			expectedStack: &FixedSizeArrayStack{
				stack:    make([]int, 10),
				capacity: 10,
				top:      -1,
			},
		},
		{
			capacity: 0,
			expectedStack: &FixedSizeArrayStack{
				stack:    make([]int, 0),
				capacity: 0,
				top:      -1,
			},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("NewFixedSizeArrayStackWithCapacity %v", index+1)
		suite.Run(testName, func() {
			stack := NewFixedSizeArrayStackWithCapacity(test.capacity)
			suite.Equal(test.expectedStack.capacity, stack.capacity)
			suite.Equal(test.expectedStack.top, stack.top)
			suite.Equal(len(test.expectedStack.stack), len(stack.stack))
		})
	}
}
