package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FindBeginOfLoopTestSuite struct {
	suite.Suite
}

func TestFindBeginOfLoopSuite(t *testing.T) {
	suite.Run(t, &FindBeginOfLoopTestSuite{})
}

func (testSuite *FindBeginOfLoopTestSuite) TestCase() {
	tests := []struct {
		values         []int
		expectedResult int
	}{
		{
			values:         []int{1, 2, 3, 4, 5},
			expectedResult: -1,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 1},
			expectedResult: 1,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 5},
			expectedResult: 5,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 3},
			expectedResult: 3,
		},
		{
			values:         []int{},
			expectedResult: -1,
		},
	}

	for index, test := range tests {
		testSuite.Run(fmt.Sprintf("Test Case: %v", index), func() {
			head := testSuite.InitTestCase(test.values)
			cc := NewFindBeginOfLoop(head)
			cc.Execute()
			testSuite.Equal(test.expectedResult, cc.Result, "Expected: %v Actual: %v", test.expectedResult, cc.Result)
		})
	}
}

func (testSuite *FindBeginOfLoopTestSuite) InitTestCase(values []int) (head *SingleLinkedList) {
	if len(values) == 0 {
		return nil
	}
	head = NewSingleLinkedList(values[0])
	current := head
	for _, val := range values[1:] {
		alreadyExists := SearchInSingleLL(head, val)
		if alreadyExists != nil {
			current.SetNext(alreadyExists)
			return
		} else {
			InsertAtEndofSingleLL(&current, val)
		}
		current = current.GetNext()
	}
	return
}
