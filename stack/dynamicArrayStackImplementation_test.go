package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DynamicArrayStackTestSuite struct {
	suite.Suite
	stack *DynamicArrayStack
}

func TestDynamicArrayStack(t *testing.T) {
	suite.Run(t, new(DynamicArrayStackTestSuite))
}

func (suite *DynamicArrayStackTestSuite) SetupTest() {
	suite.stack = NewDynamicArrayStack()
}

func (suite *DynamicArrayStackTestSuite) TestPush() {
	tests := []struct {
		values       []int
		expectedData []int
	}{
		{
			values:       []int{1},
			expectedData: []int{1},
		},
		{
			values:       []int{1, 2},
			expectedData: []int{1, 2},
		},
		{
			values:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
			expectedData: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Push %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.values {
				suite.stack.Push(value)
			}
			for i, expected := range test.expectedData {
				suite.Equal(expected, suite.stack.stack[i])
			}
			suite.stack = NewDynamicArrayStack()
		})
	}
}

func (suite *DynamicArrayStackTestSuite) TestPop() {
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
			// Explicitly test popping from an empty stack
			_, err := suite.stack.Pop()
			suite.Equal(ErrStackEmpty, err)
			suite.stack = NewDynamicArrayStack()
		})
	}
}

func (suite *DynamicArrayStackTestSuite) TestTop() {
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
			suite.stack = NewDynamicArrayStack()
		})
	}
}

func (suite *DynamicArrayStackTestSuite) TestIsEmpty() {
	suite.True(suite.stack.IsEmpty())

	suite.stack.Push(1)
	suite.False(suite.stack.IsEmpty())

	suite.stack.Pop()
	suite.True(suite.stack.IsEmpty())
}

func (suite *DynamicArrayStackTestSuite) TestSize() {
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

func (suite *DynamicArrayStackTestSuite) TestToString() {
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
			suite.stack = NewDynamicArrayStack()
		})
	}
}

func (suite *DynamicArrayStackTestSuite) TestClear() {
	tests := []struct {
		initialValues []int
	}{
		{
			initialValues: []int{1, 2, 3},
		},
		{
			initialValues: []int{4, 5, 6, 7},
		},
		{
			initialValues: []int{},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Clear %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.initialValues {
				suite.stack.Push(value)
			}
			suite.stack.Clear()
			suite.True(suite.stack.IsEmpty())
			suite.stack = NewDynamicArrayStack()
		})
	}
}

func (suite *DynamicArrayStackTestSuite) TestExpand() {
	tests := []struct {
		initialCapacity  int
		initialValues    []int
		expectedCapacity int
	}{
		{
			initialCapacity:  4,
			initialValues:    []int{1, 2, 3, 4},
			expectedCapacity: 8,
		},
		{
			initialCapacity:  8,
			initialValues:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			expectedCapacity: 16,
		},
		{
			initialCapacity:  MaxDynamicCapacity/2 + 1,
			initialValues:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			expectedCapacity: MaxDynamicCapacity,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Expand %v", index+1)
		suite.Run(testName, func() {
			suite.stack = &DynamicArrayStack{
				stack:    make([]int, test.initialCapacity),
				top:      -1,
				capacity: test.initialCapacity,
			}
			for _, value := range test.initialValues {
				suite.stack.Push(value)
			}
			suite.stack.expand()
			suite.Equal(test.expectedCapacity, suite.stack.capacity)
			for i, value := range test.initialValues {
				suite.Equal(value, suite.stack.stack[i])
			}
			suite.stack = NewDynamicArrayStack()
		})
	}
}

func (suite *DynamicArrayStackTestSuite) TestNewDynamicArrayStackTestSuiteWithCapacity() {
	tests := []struct {
		capacity      int
		expectedStack *DynamicArrayStack
	}{
		{
			capacity: 5,
			expectedStack: &DynamicArrayStack{
				stack:    make([]int, 5),
				capacity: 5,
				top:      -1,
			},
		},
		{
			capacity: 10,
			expectedStack: &DynamicArrayStack{
				stack:    make([]int, 10),
				capacity: 10,
				top:      -1,
			},
		},
		{
			capacity: 0,
			expectedStack: &DynamicArrayStack{
				stack:    make([]int, 0),
				capacity: 0,
				top:      -1,
			},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("DynamicArrayStackWithCapacity %v", index+1)
		suite.Run(testName, func() {
			stack := NewDynamicArrayStackWithCapacity(test.capacity)
			suite.Equal(test.expectedStack.capacity, stack.capacity)
			suite.Equal(test.expectedStack.top, stack.top)
			suite.Equal(len(test.expectedStack.stack), len(stack.stack))
		})
	}
}
