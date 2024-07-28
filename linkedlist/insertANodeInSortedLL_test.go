package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type InsertANodeInSortedLLSuite struct {
	suite.Suite
}

func TestInsertANodeInSortedLLSuite(t *testing.T) {
	suite.Run(t, &InsertANodeInSortedLLSuite{})
}

func (testSuite *InsertANodeInSortedLLSuite) TestCase() {
	tests := []struct {
		values         []int
		insertValue    int
		expectedResult []int
	}{
		{
			values:         []int{1, 2, 3, 4, 5},
			insertValue:    6,
			expectedResult: []int{1, 2, 3, 4, 5, 6},
		},
		{
			values:         []int{1, 2, 3, 4, 5},
			insertValue:    3,
			expectedResult: []int{1, 2, 3, 3, 4, 5},
		},
		{
			values:         []int{1, 2, 3, 4, 5},
			insertValue:    0,
			expectedResult: []int{0, 1, 2, 3, 4, 5},
		},
		{
			values:         []int{1, 2, 3, 4, 5},
			insertValue:    2,
			expectedResult: []int{1, 2, 2, 3, 4, 5},
		},
		{
			values:         []int{1, 2, 3, 4, 5},
			insertValue:    5,
			expectedResult: []int{1, 2, 3, 4, 5, 5},
		},
		{
			values:         []int{},
			insertValue:    1,
			expectedResult: []int{1},
		},
	}

	for index, test := range tests {
		testSuite.Run(fmt.Sprintf("Test Case %v", index), func() {
			head := testSuite.InitTestCase(test.values)
			cc := NewInsertANodeInSortedLL(head)
			cc.Execute(test.insertValue)
			testSuite.Equal(test.expectedResult, cc.Result, "Expected: %v Actual: %v", test.expectedResult, cc.Result)
		})
	}
}

func (testSuite *InsertANodeInSortedLLSuite) InitTestCase(values []int) (head *SingleLinkedList) {
	if len(values) == 0 {
		return nil
	}
	head = NewSingleLinkedList(values[0])
	current := head
	for _, val := range values[1:] {
		current.SetNext(NewSingleLinkedList(val))
		current = current.GetNext()
	}
	return
}
