package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ReorderTestSuite struct {
	suite.Suite
}

func TestReorderSuite(t *testing.T) {
	suite.Run(t, &ReorderTestSuite{})
}

func (testSuite *ReorderTestSuite) TestCase() {
	tests := []struct {
		val            []int
		expectedResult []int
	}{
		{
			val:            []int{1, 2, 3, 4, 5},
			expectedResult: []int{1, 5, 2, 4, 3},
		},
		{
			val:            []int{1, 2, 3, 4},
			expectedResult: []int{1, 4, 2, 3},
		},
		{
			val:            []int{1, 2, 3},
			expectedResult: []int{1, 3, 2},
		},
		{
			val:            []int{1, 2},
			expectedResult: []int{1, 2},
		},
		{
			val:            []int{1},
			expectedResult: []int{1},
		},
		{
			val:            []int{},
			expectedResult: []int{},
		},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("Test Case %+v", i)
		testSuite.Run(testname, func() {
			obj := NewReorderLinkedList(test.val)
			obj.Execute()
			testSuite.Equal(test.expectedResult, obj.Result)
		})
	}
}
