package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StackForStackSetsTestSuite struct {
	suite.Suite
	stackSet *StackForStackSets
}

func TestStackForStackSets(t *testing.T) {
	suite.Run(t, new(StackForStackSetsTestSuite))
}

func (suite *StackForStackSetsTestSuite) SetupTest() {
	suite.stackSet = NewStackForStackSets(3)
}

func (suite *StackForStackSetsTestSuite) TestPush() {
	tests := []struct {
		values         []int
		expectedStacks int
	}{
		{
			values:         []int{1, 2, 3},
			expectedStacks: 1,
		},
		{
			values:         []int{1, 2, 3, 4},
			expectedStacks: 2,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 6},
			expectedStacks: 2,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 6, 7},
			expectedStacks: 3,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Push %v", index+1)
		suite.Run(testName, func() {
			suite.stackSet = NewStackForStackSets(3)
			for _, value := range test.values {
				suite.stackSet.Push(value)
			}
			suite.Equal(test.expectedStacks, len(suite.stackSet.stacks))
		})
	}
}

func (suite *StackForStackSetsTestSuite) TestPop() {
	tests := []struct {
		values         []int
		popCount       int
		expectedResult []int
		expectedErr    error
	}{
		{
			values:         []int{1, 2, 3},
			popCount:       3,
			expectedResult: []int{3, 2, 1},
			expectedErr:    nil,
		},
		{
			values:         []int{1, 2, 3, 4},
			popCount:       4,
			expectedResult: []int{4, 3, 2, 1},
			expectedErr:    nil,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 6},
			popCount:       6,
			expectedResult: []int{6, 5, 4, 3, 2, 1},
			expectedErr:    nil,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 6, 7},
			popCount:       7,
			expectedResult: []int{7, 6, 5, 4, 3, 2, 1},
			expectedErr:    nil,
		},
		{
			values:         []int{1, 2, 3},
			popCount:       4,
			expectedResult: []int{3, 2, 1, 0},
			expectedErr:    ErrStackEmpty,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Pop %v", index+1)
		suite.Run(testName, func() {
			suite.stackSet = NewStackForStackSets(3)
			for _, value := range test.values {
				suite.stackSet.Push(value)
			}
			var result []int
			var err error
			for i := 0; i < test.popCount; i++ {
				var value int
				value, err = suite.stackSet.Pop()
				result = append(result, value)
				if err != nil {
					break
				}
			}
			suite.Equal(test.expectedErr, err)
			suite.Equal(test.expectedResult, result)
		})
	}
}

func (suite *StackForStackSetsTestSuite) TestPopFromNth() {
	tests := []struct {
		values         []int
		popCount       int
		n              int
		expectedResult int
		expectedErr    error
	}{
		{
			values:         []int{1, 2, 3, 4, 5, 6},
			popCount:       1,
			n:              0,
			expectedResult: 3,
			expectedErr:    nil,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 6},
			n:              1,
			popCount:       1,
			expectedResult: 6,
			expectedErr:    nil,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 6},
			n:              2,
			popCount:       1,
			expectedResult: 0,
			expectedErr:    ErrInvalidStackId,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 6, 7},
			n:              1,
			popCount:       3,
			expectedResult: 4,
			expectedErr:    nil,
		},
		{
			values:         []int{1, 2, 3},
			n:              0,
			popCount:       1,
			expectedResult: 3,
			expectedErr:    nil,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("PopFromNth %v", index+1)
		suite.Run(testName, func() {
			suite.stackSet = NewStackForStackSets(3)
			for _, value := range test.values {
				suite.stackSet.Push(value)
			}
			result := -1
			var err error
			for i := 0; i < test.popCount; i++ {
				result, err = suite.stackSet.PopFromNth(test.n)
			}
			suite.Equal(test.expectedErr, err)
			suite.Equal(test.expectedResult, result)
		})
	}
}
