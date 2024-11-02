package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TwoStackWithOneArrayTestSuite struct {
	suite.Suite
	stack *TwoStackWithOneArray
}

func TestTwoStackWithOneArray(t *testing.T) {
	suite.Run(t, new(TwoStackWithOneArrayTestSuite))
}

func (suite *TwoStackWithOneArrayTestSuite) SetupTest() {
	suite.stack = NewTwoStackWithOneArray(10)
}

func (suite *TwoStackWithOneArrayTestSuite) TestPeek() {
	tests := []struct {
		stackId        int
		operations     []func()
		expectedResult int
		expectedErr    error
	}{
		{
			stackId:        3,
			operations:     []func(){},
			expectedResult: 0,
			expectedErr:    ErrInvalidStackId,
		},
		{
			stackId:        1,
			operations:     []func(){},
			expectedResult: 0,
			expectedErr:    ErrStackEmpty,
		},
		{
			stackId:        2,
			operations:     []func(){},
			expectedResult: 0,
			expectedErr:    ErrStackEmpty,
		},
		{
			stackId: 1,
			operations: []func(){
				func() { suite.stack.Push(1, 10) },
			},
			expectedResult: 10,
			expectedErr:    nil,
		},
		{
			stackId: 2,
			operations: []func(){
				func() { suite.stack.Push(2, 20) },
			},
			expectedResult: 20,
			expectedErr:    nil,
		},
		{
			stackId: 1,
			operations: []func(){
				func() { suite.stack.Push(1, 10) },
				func() { suite.stack.Push(1, 20) },
			},
			expectedResult: 20,
			expectedErr:    nil,
		},
		{
			stackId: 2,
			operations: []func(){
				func() { suite.stack.Push(2, 30) },
				func() { suite.stack.Push(2, 40) },
			},
			expectedResult: 40,
			expectedErr:    nil,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Peek %v", index+1)
		suite.Run(testName, func() {
			for _, operation := range test.operations {
				operation()
			}
			result, err := suite.stack.Peek(test.stackId)
			suite.Equal(test.expectedErr, err)
			suite.Equal(test.expectedResult, result)
		})
	}
}

func (suite *TwoStackWithOneArrayTestSuite) TestIsEmpty() {
	tests := []struct {
		stackId        int
		operations     []func()
		expectedResult bool
	}{
		{
			stackId:        1,
			operations:     []func(){},
			expectedResult: true,
		},
		{
			stackId: 1,
			operations: []func(){
				func() { suite.stack.Push(1, 10) },
			},
			expectedResult: false,
		},
		{
			stackId: 2,
			operations: []func(){
				func() { suite.stack.Push(2, 20) },
			},
			expectedResult: false,
		},
		{
			stackId: 1,
			operations: []func(){
				func() { suite.stack.Push(1, 10) },
				func() { suite.stack.Pop(1) },
			},
			expectedResult: true,
		},
		{
			stackId: 2,
			operations: []func(){
				func() { suite.stack.Push(2, 30) },
				func() { suite.stack.Pop(2) },
			},
			expectedResult: true,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("IsEmpty %v", index+1)
		suite.Run(testName, func() {
			suite.stack = NewTwoStackWithOneArray(10)
			for _, operation := range test.operations {
				operation()
			}
			result := suite.stack.IsEmpty(test.stackId)
			suite.Equal(test.expectedResult, result)
		})
	}
}

func (suite *TwoStackWithOneArrayTestSuite) TestPush() {
	tests := []struct {
		stackId     int
		values      []int
		expectedErr error
	}{
		{
			stackId:     1,
			values:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectedErr: nil,
		},
		{
			stackId:     1,
			values:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			expectedErr: ErrStackFull,
		},
		{
			stackId:     2,
			values:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectedErr: nil,
		},
		{
			stackId:     2,
			values:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			expectedErr: ErrStackFull,
		},
		{
			stackId:     3,
			values:      []int{1},
			expectedErr: ErrInvalidStackId,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Push %v", index+1)
		suite.Run(testName, func() {
			var err error
			for _, value := range test.values {
				err = suite.stack.Push(test.stackId, value)
				if err != nil {
					break
				}
			}
			suite.stack = NewTwoStackWithOneArray(10)
			suite.Equal(test.expectedErr, err)
		})
	}
}

func (suite *TwoStackWithOneArrayTestSuite) TestPop() {
	tests := []struct {
		stackId        int
		operations     []func()
		expectedResult int
		expectedErr    error
	}{
		{
			stackId:        1,
			operations:     []func(){},
			expectedResult: 0,
			expectedErr:    ErrStackEmpty,
		},
		{
			stackId:        2,
			operations:     []func(){},
			expectedResult: 0,
			expectedErr:    ErrStackEmpty,
		},
		{
			stackId: 1,
			operations: []func(){
				func() { suite.stack.Push(1, 10) },
			},
			expectedResult: 10,
			expectedErr:    nil,
		},
		{
			stackId: 2,
			operations: []func(){
				func() { suite.stack.Push(2, 20) },
			},
			expectedResult: 20,
			expectedErr:    nil,
		},
		{
			stackId: 1,
			operations: []func(){
				func() { suite.stack.Push(1, 10) },
				func() { suite.stack.Push(1, 20) },
			},
			expectedResult: 20,
			expectedErr:    nil,
		},
		{
			stackId: 2,
			operations: []func(){
				func() { suite.stack.Push(2, 30) },
				func() { suite.stack.Push(2, 40) },
			},
			expectedResult: 40,
			expectedErr:    nil,
		},
		{
			stackId:        3,
			operations:     []func(){},
			expectedResult: 0,
			expectedErr:    ErrInvalidStackId,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Pop %v", index+1)
		suite.Run(testName, func() {
			for _, operation := range test.operations {
				operation()
			}
			result, err := suite.stack.Pop(test.stackId)
			suite.Equal(test.expectedErr, err)
			suite.Equal(test.expectedResult, result)
		})
	}
}
