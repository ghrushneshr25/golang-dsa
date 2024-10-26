package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PalindromeLinkedListTestSuite struct {
	suite.Suite
}

func TestPalindromeLinkedListTestSuite(t *testing.T) {
	suite.Run(t, &PalindromeLinkedListTestSuite{})
}

func (testSuite *PalindromeLinkedListTestSuite) TestCase() {
	tests := []struct {
		values         []int
		expectedResult bool
	}{
		{
			values:         []int{1, 2, 3, 2, 1},
			expectedResult: true,
		},
		{
			values:         []int{1},
			expectedResult: true,
		},
		{
			values:         []int{1, 2, 3, 2, 2},
			expectedResult: false,
		},
		{
			values:         []int{1, 2},
			expectedResult: false,
		},
		{
			values:         []int{},
			expectedResult: true,
		},
	}

	for index, test := range tests {
		testSuite.Run(fmt.Sprintf("Test Case: %v", index), func() {
			cc := NewPalindromeLinkedList(test.values)
			cc.Execute()
			testSuite.Equal(test.expectedResult, cc.Result, "Expected: %v Actual: %v", test.expectedResult, cc.Result)
		})
	}
}
