package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type IsLinkedListEvenTestSuite struct {
	suite.Suite
}

func TestIsLinkedListEvenTestSuite(t *testing.T) {
	suite.Run(t, new(IsLinkedListEvenTestSuite))
}

func (s *IsLinkedListEvenTestSuite) TestCase() {
	tests := []struct {
		val            []int
		expectedResult bool
	}{
		{
			val:            []int{1, 2, 3, 4, 5},
			expectedResult: false,
		},
		{
			val:            []int{1, 2, 3, 4, 5, 6},
			expectedResult: true,
		},
		{
			val:            []int{1, 2, 3, 4, 5, 6, 7},
			expectedResult: false,
		},
		{
			val:            []int{1},
			expectedResult: false,
		},
		{
			val:            []int{},
			expectedResult: true,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("IsLinkedListEven TestCase- %+v", test.val)
		s.Run(testName, func() {
			head := s.InitTestCase(test.val)
			ille := NewIsLinkedListEven(head)
			ille.Execute()
			s.Equalf(test.expectedResult, ille.Result, "Expected : %v, Actual : %v", test.expectedResult, ille.Result)
		})
	}
}

func (s *IsLinkedListEvenTestSuite) InitTestCase(val []int) *SingleLinkedList {
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
