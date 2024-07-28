package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CheckCycleTestSuite struct {
	suite.Suite
}

func TestCheckCycleSuite(t *testing.T) {
	suite.Run(t, &CheckCycleTestSuite{})
}

func (testSuite *CheckCycleTestSuite) TestCase() {
	tests := []struct {
		values         []int
		expectedResult bool
	}{
		{
			values:         []int{1, 2, 3, 4, 5},
			expectedResult: false,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 1},
			expectedResult: true,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 5},
			expectedResult: true,
		},
		{
			values:         []int{1, 2, 3, 4, 5, 3},
			expectedResult: true,
		},
		{
			values:         []int{},
			expectedResult: false,
		},
	}

	methodNames := []string{
		"HashMapMethod",
		"TwoPointerMethod",
	}

	for _, methodName := range methodNames {
		for index, test := range tests {
			testSuite.Run(fmt.Sprintf("Test Case: %v using Method: %v", index, methodName), func() {
				head := testSuite.InitTestCase(test.values)
				cc := NewCheckCycle(head)
				cc.Execute(methodName)
				testSuite.Equal(test.expectedResult, cc.Result, "Method: %s, Values: %v", methodName, test.values)
			})
		}
	}
}

func (testSuite *CheckCycleTestSuite) InitTestCase(values []int) (head *SingleLinkedList) {
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
