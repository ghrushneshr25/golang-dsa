package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExchangeAdjacentNodesTestSuite struct {
	suite.Suite
}

func TestExchangeAdjacentNodes(t *testing.T) {
	suite.Run(t, &ExchangeAdjacentNodesTestSuite{})
}

func (testSuite *ExchangeAdjacentNodesTestSuite) TestCase() {
	tests := []struct {
		val            []int
		expectedResult []int
	}{
		{
			val:            []int{1, 2, 3, 4, 5, 6},
			expectedResult: []int{2, 1, 4, 3, 6, 5},
		},
		{
			val:            []int{1, 2, 3, 4, 5},
			expectedResult: []int{2, 1, 4, 3, 5},
		},
		{
			val:            []int{1, 2},
			expectedResult: []int{2, 1},
		},
		{
			val:            []int{1},
			expectedResult: []int{1},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Test Case :%v", index+1)
		testSuite.Run(testName, func() {
			obj := NewExchangeAdjacentNodes(test.val)
			obj.Execute()
			testSuite.Equal(test.expectedResult, obj.Result)
		})
	}
}
