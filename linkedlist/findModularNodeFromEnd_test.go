package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestFindModularNodeFromEndSuite struct {
	suite.Suite
}

func TestFindModularNodeFromEnd(t *testing.T) {
	suite.Run(t, &TestFindModularNodeFromEndSuite{})
}

func (testSuite *TestFindModularNodeFromEndSuite) TestCase() {
	tests := []struct {
		n              int
		val            []int
		k              int
		expectedResult int
	}{
		{
			n:              19,
			k:              3,
			expectedResult: 16,
		},
		{
			n:              19,
			k:              -1,
			expectedResult: -1,
		},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("TestCase %v", i)
		testSuite.Run(testname, func() {
			for i := 0; i < test.n; i++ {
				test.val = append(test.val, i)
			}
			obj := NewFindModularNodeFromEnd(test.val, test.k)
			obj.Execute()
			testSuite.Equal(test.expectedResult, obj.Result)
		})
	}
}
