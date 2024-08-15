package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MiddleOfLinkedListtTestSuite struct {
	suite.Suite
}

func TestMiddleOfLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(MiddleOfLinkedListtTestSuite))
}

func (s *MiddleOfLinkedListtTestSuite) TestCase() {
	methods := []string{
		"BruteForceMethod",
		"LengthMethod",
		"TwoPointerMethod",
		"HashMapMethod",
	}

	tests := []struct {
		val            []int
		expectedResult int
	}{
		{
			val:            []int{1, 2, 3, 4, 5},
			expectedResult: 3,
		},
		{
			val:            []int{1, 2, 3, 4, 5, 6},
			expectedResult: 4,
		},
		{
			val:            []int{1, 2, 3, 4, 5, 6, 7},
			expectedResult: 4,
		},
		{
			val:            []int{1},
			expectedResult: 1,
		},
		{
			val:            []int{},
			expectedResult: -1,
		},
	}

	for _, methodName := range methods {
		for i, test := range tests {
			testName := fmt.Sprintf("Method-%s Test-%d", methodName, i)
			s.Run(testName, func() {
				head := s.InitTestCase(test.val)
				mll := NewMiddleOfLinkedList(head)
				mll.ExecuteByMethod(methodName)
				s.Equalf(test.expectedResult, mll.Result, "Method : %v | Expected : %v | Actual : %v", methodName, test.expectedResult, mll.Result)
			})
		}
	}

}

func (s *MiddleOfLinkedListtTestSuite) InitTestCase(val []int) *SingleLinkedList {
	if len(val) == 0 {
		return nil
	}

	head := NewSingleLinkedList(val[0])
	current := head
	for i := 1; i < len(val); i++ {
		current.SetNext(NewSingleLinkedList(val[i]))
		current = current.GetNext()
	}
	return head
}
