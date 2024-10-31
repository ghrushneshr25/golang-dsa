package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LinkedListStackTestSuite struct {
	suite.Suite
	stack *LinkedListStack
}

func TestLinkedListStack(t *testing.T) {
	suite.Run(t, new(LinkedListStackTestSuite))
}

func (suite *LinkedListStackTestSuite) SetupTest() {
	suite.stack = NewLinkedListStack()
}

func (suite *LinkedListStackTestSuite) TestPush() {
	tests := []struct {
		values       []interface{}
		expectedData []interface{}
	}{
		{
			values:       []interface{}{1},
			expectedData: []interface{}{1},
		},
		{
			values:       []interface{}{1, 2},
			expectedData: []interface{}{2, 1},
		},
		{
			values:       []interface{}{1, 2, 3},
			expectedData: []interface{}{3, 2, 1},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Push %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.values {
				suite.stack.Push(value)
			}
			current := suite.stack.top
			for _, expected := range test.expectedData {
				suite.Equal(expected, current.CustomData)
				current = current.GetNext()
			}
			suite.stack = NewLinkedListStack()
		})
	}
}

func (suite *LinkedListStackTestSuite) TestPop() {
	tests := []struct {
		initialValues []interface{}
		expectedData  []interface{}
		expectedErr   error
	}{
		{
			initialValues: []interface{}{1, 2, 3},
			expectedData:  []interface{}{3, 2, 1},
			expectedErr:   nil,
		},
		{
			initialValues: []interface{}{1, 2},
			expectedData:  []interface{}{2, 1},
			expectedErr:   nil,
		},
		{
			initialValues: []interface{}{1},
			expectedData:  []interface{}{1},
			expectedErr:   nil,
		},
		{
			initialValues: []interface{}{},
			expectedData:  []interface{}{},
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
			suite.stack = NewLinkedListStack()
		})
	}
}

func (suite *LinkedListStackTestSuite) TestPeek() {
	tests := []struct {
		initialValues []interface{}
		expectedData  interface{}
		expectedErr   error
	}{
		{
			initialValues: []interface{}{1, 2, 3},
			expectedData:  3,
			expectedErr:   nil,
		},
		{
			initialValues: []interface{}{1, 2},
			expectedData:  2,
			expectedErr:   nil,
		},
		{
			initialValues: []interface{}{1},
			expectedData:  1,
			expectedErr:   nil,
		},
		{
			initialValues: []interface{}{},
			expectedData:  0,
			expectedErr:   ErrStackEmpty,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Peek %v", index+1)
		suite.Run(testName, func() {
			for _, value := range test.initialValues {
				suite.stack.Push(value)
			}
			val, err := suite.stack.Peek()
			suite.Equal(test.expectedErr, err)
			if err == nil {
				suite.Equal(test.expectedData, val)
			}
			suite.stack = NewLinkedListStack()
		})
	}
}

func (suite *LinkedListStackTestSuite) TestIsEmpty() {
	suite.True(suite.stack.IsEmpty())

	suite.stack.Push(1)
	suite.False(suite.stack.IsEmpty())

	suite.stack.Pop()
	suite.True(suite.stack.IsEmpty())
}

func (suite *LinkedListStackTestSuite) TestSize() {
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

func (suite *LinkedListStackTestSuite) TestToString() {
	tests := []struct {
		initialValues []interface{}
		expectedStr   string
	}{
		{
			initialValues: []interface{}{},
			expectedStr:   "[]",
		},
		{
			initialValues: []interface{}{1},
			expectedStr:   "[ 1 ]",
		},
		{
			initialValues: []interface{}{1, 2, 3},
			expectedStr:   "[ 3  2  1 ]",
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
			suite.stack = NewLinkedListStack()
		})
	}
}
