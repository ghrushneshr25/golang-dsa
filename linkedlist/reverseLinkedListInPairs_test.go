package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ReverseLinkedListInPairsTestSuite struct {
	suite.Suite
}

func TestReverseLinkedListInPairsTestSuite(t *testing.T) {
	suite.Run(t, new(ReverseLinkedListInPairsTestSuite))
}

func (testSuite *ReverseLinkedListInPairsTestSuite) TestCase() {
	methods := []string{
		"RecursiveMethod",
		"IterativeMethod",
	}

	testCases := []struct {
		values   []int
		expected []int
	}{
		{
			values:   []int{1, 2, 3, 4, 5, 6},
			expected: []int{2, 1, 4, 3, 6, 5},
		},
		{
			values:   []int{1, 2, 3, 4, 5},
			expected: []int{2, 1, 4, 3, 5},
		},
		{
			values:   []int{},
			expected: []int{},
		},
	}

	for _, methodName := range methods {
		for index, test := range testCases {
			testName := fmt.Sprintf("TestCase %v %v", index+1, methodName)
			testSuite.Run(testName, func() {
				rllip := NewReverseLinkedListInPairs(test.values)
				rllip.ExecuteByMethodName(methodName)
				testSuite.Equal(test.expected, rllip.Result)
			})
		}
	}
}
