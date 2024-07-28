package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NthNodeFromEndSuite struct {
	suite.Suite
}

func TestNthNodeFromEndSuite(t *testing.T) {
	suite.Run(t, &NthNodeFromEndSuite{})
}

func (testSuite *NthNodeFromEndSuite) TestCase() {
	tests := []struct {
		values         []int
		n              int
		expectedResult int
	}{
		{
			values:         []int{1, 2, 3, 4, 5},
			n:              2,
			expectedResult: 4,
		},
		{
			values:         []int{10, 20, 30, 40, 50},
			n:              1,
			expectedResult: 50,
		},
		{
			values:         []int{10, 20, 30, 40, 50},
			n:              5,
			expectedResult: 10,
		},
		{
			values:         []int{1, 2, 3},
			n:              4,
			expectedResult: -1, // n is greater than the length of the list
		},
		{
			values:         []int{1},
			n:              1,
			expectedResult: 1,
		},
		{
			values:         []int{},
			n:              1,
			expectedResult: -1, // empty list
		},
		{
			values:         []int{1, 2, 3},
			n:              3,
			expectedResult: 1,
		},
		{
			values:         []int{1, 2, 3},
			n:              -2,
			expectedResult: -1,
		},
	}

	methods := []string{
		"BruteForceMethod",
		"HashMapMethod",
		"TwoPointerMethod",
		"RecursionMethod",
	}

	for _, method := range methods {
		for index, test := range tests {
			testSuite.Run(fmt.Sprintf("Test Case: %v using Method: %v", index, method), func() {
				nthNodeFromEnd := NewNthNodeFromEnd(test.values, test.n)
				nthNodeFromEnd.Execute(method)
				testSuite.Equal(test.expectedResult, nthNodeFromEnd.Result)
			})
		}
	}
}
