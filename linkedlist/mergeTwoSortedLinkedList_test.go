package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MergeTwoSortedLinkedListTestSuite struct {
	suite.Suite
}

func TestMergeTwoSortedLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(MergeTwoSortedLinkedListTestSuite))
}

func (mtsll *MergeTwoSortedLinkedListTestSuite) TestCase() {
	methods := []string{
		"RecursionMethod",
		"IterativeMethod",
	}

	testCases := []struct {
		headOne  []int
		headTwo  []int
		expected []int
	}{
		{
			headOne:  []int{1, 3, 5},
			headTwo:  []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			headOne:  []int{1, 2, 3},
			headTwo:  []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			headOne:  []int{4, 5, 6},
			headTwo:  []int{1, 2, 3},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			headOne:  []int{1, 2, 3},
			headTwo:  []int{},
			expected: []int{1, 2, 3},
		},
		{
			headOne:  []int{},
			headTwo:  []int{4, 5, 6},
			expected: []int{4, 5, 6},
		},
		{
			headOne:  []int{},
			headTwo:  []int{},
			expected: []int{},
		},
	}

	for _, method := range methods {
		for index, test := range testCases {
			testName := fmt.Sprintf("Method Name %v | Test Case :%v", method, index+1)
			mtsll.Run(testName, func() {
				headOne := ConvertArrayToSingleLinkedList(test.headOne)
				headTwo := ConvertArrayToSingleLinkedList(test.headTwo)
				mtslll := NewMergeTwoSortedLinkedList(headOne, headTwo)
				mtslll.ExecuteByMethodName(method)
				mtsll.Equalf(test.expected, mtslll.Result, "Expected : %v, Got: %v", test.expected, mtslll.Result)
			})
		}
	}
}
