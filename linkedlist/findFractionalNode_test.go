package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestFindFractionalNodeSuite struct {
	suite.Suite
}

func TestFindFractionalNode(t *testing.T) {
	suite.Run(t, &TestFindFractionalNodeSuite{})
}

func (testSuite *TestFindFractionalNodeSuite) TestCase() {
	tests := []struct {
		n              int
		val            []int
		k              int
		expectedResult int
	}{
		{
			n:              10,
			k:              2,
			expectedResult: 5,
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
			obj := NewFindFractionalNode(test.val, test.k)
			obj.Execute()
			testSuite.Equal(test.expectedResult, obj.Result)
		})
	}
}
