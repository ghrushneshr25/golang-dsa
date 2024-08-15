package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ReverseSingleLinkedListSuite struct {
	suite.Suite
}

func TestReverseSingleLinkedListSuite(t *testing.T) {
	suite.Run(t, &ReverseSingleLinkedListSuite{})
}

func (testSuite *ReverseSingleLinkedListSuite) TestCase() {
	tests := []struct {
		values         []int
		expectedResult []int
	}{
		{
			values:         []int{1, 2, 3, 4, 5},
			expectedResult: []int{5, 4, 3, 2, 1},
		},
		{
			values:         []int{10, 20, 30, 40, 50},
			expectedResult: []int{50, 40, 30, 20, 10},
		},
		{
			values:         []int{},
			expectedResult: []int{},
		},
	}

	methodNames := []string{
		"RecursiveApproach",
		"IterativeApproach",
	}

	for _, methodName := range methodNames {
		for index, test := range tests {
			testSuite.Run(fmt.Sprintf("Test Case: %v using Method: %v", index, methodName), func() {
				rsll := NewReverseSingleLinkedList(test.values)
				rsll.ExecuteByName(methodName)
				testSuite.Equal(test.expectedResult, rsll.Result)
			})
		}
	}
}
