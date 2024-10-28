package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RotateByKTestSuite struct {
	suite.Suite
}

func TestRotateByKSuite(t *testing.T) {
	suite.Run(t, &RotateByKTestSuite{})
}

func (testSuite *RotateByKTestSuite) TestCase() {
	tests := []struct {
		val            []int
		k              int
		expectedResult []int
	}{
		{
			val:            []int{1, 2, 3, 4, 5},
			k:              2,
			expectedResult: []int{4, 5, 1, 2, 3},
		},
		{
			val:            []int{1, 2, 3, 4, 5},
			k:              3,
			expectedResult: []int{3, 4, 5, 1, 2},
		},
		{
			val:            []int{1, 2, 3, 4, 5},
			k:              5,
			expectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			val:            []int{1, 2, 3, 4, 5},
			k:              1,
			expectedResult: []int{5, 1, 2, 3, 4},
		},
		{
			val:            []int{1, 2, 3, 4, 5},
			k:              0,
			expectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			val:            []int{},
			k:              3,
			expectedResult: []int{},
		},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("TestCase %v", i)
		testSuite.Run(testname, func() {
			obj := NewRotateByK(test.val, test.k)
			obj.Execute()
			testSuite.Equal(test.expectedResult, obj.Result)
		})
	}
}
