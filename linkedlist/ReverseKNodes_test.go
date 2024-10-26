package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ReverseKNodesTestSuite struct {
	suite.Suite
}

func TestReverseKNodesSuite(t *testing.T) {
	suite.Run(t, &ReverseKNodesTestSuite{})
}

func (testSuite *ReverseKNodesTestSuite) TestCase() {

	tests := []struct {
		val            []int
		k              int
		expectedResult []int
	}{
		{
			val:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			k:              2,
			expectedResult: []int{2, 1, 4, 3, 6, 5, 8, 7, 10, 9},
		},
		{
			val:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			k:              3,
			expectedResult: []int{3, 2, 1, 6, 5, 4, 9, 8, 7, 10},
		},
		{
			val:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			k:              4,
			expectedResult: []int{4, 3, 2, 1, 8, 7, 6, 5, 9, 10},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Test Case :%v", index+1)
		testSuite.Run(testName, func() {
			obj := NewReverseKNodes(test.val, test.k)
			obj.Execute()
			testSuite.Equal(test.expectedResult, obj.Result)
		})
	}

}
