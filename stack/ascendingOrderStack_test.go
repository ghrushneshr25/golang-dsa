package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AscendingOrderStackTestSuite struct {
	suite.Suite
}

func TestAscendingOrderStack(t *testing.T) {
	suite.Run(t, new(AscendingOrderStackTestSuite))
}

func (suite *AscendingOrderStackTestSuite) TestSortStack() {
	tests := []struct {
		input          []int
		expectedResult []int
	}{
		{
			input:          []int{3, 1, 4, 2},
			expectedResult: []int{1, 2, 3, 4},
		},
		{
			input:          []int{5, 1, 3, 2, 4},
			expectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			input:          []int{1, 2, 3, 4, 5},
			expectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			input:          []int{5, 4, 3, 2, 1},
			expectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			input:          []int{2, 2, 2, 2, 2},
			expectedResult: []int{2, 2, 2, 2, 2},
		},
		{
			input:          []int{1},
			expectedResult: []int{1},
		},
		{
			input:          []int{},
			expectedResult: []int{},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("SortStack %v", index+1)
		suite.Run(testName, func() {
			obj := NewAscendingOrderStack(test.input)
			obj.Execute()
			for i := len(test.expectedResult) - 1; i >= 0; i-- {
				result, _ := obj.Result.Pop()
				suite.Equal(test.expectedResult[i], result)

			}
		})
	}
}
