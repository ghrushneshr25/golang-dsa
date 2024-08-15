package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type IntersectionOfTwoLinkedListTestSuite struct {
	suite.Suite
}

func TestIntersectionOfTwoLinkedListSuite(t *testing.T) {
	suite.Run(t, &IntersectionOfTwoLinkedListTestSuite{})
}

func (testSuite *IntersectionOfTwoLinkedListTestSuite) TestCase() {
	tests := []struct {
		headOne        []int
		headTwo        []int
		expectedResult int
	}{
		{
			headOne:        []int{4, 1, 8, 4, 5},
			headTwo:        []int{2, 0, 1, 8, 4, 5},
			expectedResult: 1,
		},
		{
			headOne:        []int{0, 9, 1, 2, 4},
			headTwo:        []int{3, 2, 4},
			expectedResult: 2,
		},
		{
			headOne:        []int{2, 6, 4},
			headTwo:        []int{1, 5},
			expectedResult: -1,
		},
		{
			headOne:        []int{},
			headTwo:        []int{},
			expectedResult: -1,
		},
	}

	methods := []string{
		"BruteForceMethod",
		"StackMethod",
		"TwoPointerMethod",
	}

	for _, method := range methods {
		for index, test := range tests {
			testSuite.Run(fmt.Sprintf("Test Case : %v | Method : %v", index, method), func() {
				headOne, headTwo := testSuite.InitTestCase(test.headOne, test.headTwo)
				cc := NewIntersectionOfTwoLinkedList(headOne, headTwo)
				cc.ExecuteByName(method)
				testSuite.Equal(test.expectedResult, cc.Result, "Method : %v | Expected : %v | Actual : %v", method, test.expectedResult, cc.Result)
			})
		}
	}
}

func (testSuite *IntersectionOfTwoLinkedListTestSuite) InitTestCase(values1, values2 []int) (headOne, headTwo *SingleLinkedList) {
	for _, value := range values1 {
		InsertAtEndofSingleLL(&headOne, value)
	}
	for _, value := range values2 {
		if node := SearchInSingleLL(headOne, value); node != nil {
			current := headTwo
			for current.GetNext() != nil {
				current = current.GetNext()
			}
			current.SetNext(node)
			break
		}
		InsertAtEndofSingleLL(&headTwo, value)

	}
	return
}
