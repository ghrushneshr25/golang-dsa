package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RemoveDuplicatesFromUnorderedListTestSuite struct {
	suite.Suite
}

func TestRemoveDuplicatesFromUnorderedListSuite(t *testing.T) {
	suite.Run(t, &RemoveDuplicatesFromUnorderedListTestSuite{})
}

func (testSuite *RemoveDuplicatesFromUnorderedListTestSuite) TestCase() {
	tests := []struct {
		val            []int
		expectedResult []int
	}{
		{
			val:            []int{1, 2, 3, 2, 4, 3, 5},
			expectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			val:            []int{1, 1, 1, 1, 1},
			expectedResult: []int{1},
		},
		{
			val:            []int{1, 2, 3, 4, 5},
			expectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			val:            []int{5, 4, 3, 2, 1, 5, 4, 3, 2, 1},
			expectedResult: []int{5, 4, 3, 2, 1},
		},
		{
			val:            []int{},
			expectedResult: []int{},
		},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("TestCase %v", i)
		testSuite.Run(testname, func() {
			obj := NewRemoveDuplicatesFromUnorderedList(test.val)
			obj.Execute()
			testSuite.Equal(test.expectedResult, obj.Result)
		})
	}
}
