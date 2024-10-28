package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PartitionByKTestSuite struct {
	suite.Suite
}

func TestPartitionByKSuite(t *testing.T) {
	suite.Run(t, &PartitionByKTestSuite{})
}

func (testSuite *PartitionByKTestSuite) TestCase() {
	tests := []struct {
		val            []int
		K              int
		expectedResult []int
	}{
		{
			val:            []int{1, 4, 3, 2, 5, 2},
			K:              3,
			expectedResult: []int{1, 2, 2, 4, 3, 5},
		},
		{
			val:            []int{1, 4, 3, 2, 5, 2},
			K:              5,
			expectedResult: []int{1, 4, 3, 2, 2, 5},
		},
		{
			val:            []int{1, 4, 3, 2, 5, 2},
			K:              1,
			expectedResult: []int{1, 4, 3, 2, 5, 2},
		},
		{
			val:            []int{1, 4, 3, 2, 5, 2},
			K:              2,
			expectedResult: []int{1, 4, 3, 2, 5, 2},
		},
		{
			val:            []int{1, 4, 3, 2, 5, 2},
			K:              4,
			expectedResult: []int{1, 3, 2, 2, 4, 5},
		},
		{
			val:            []int{},
			K:              3,
			expectedResult: []int{},
		},
		{
			val:            []int{1, 2, 2},
			K:              3,
			expectedResult: []int{1, 2, 2},
		},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("TestCase %v", i)
		testSuite.Run(testname, func() {
			obj := NewPartitionByK(test.val, test.K)
			obj.Execute()
			testSuite.Equal(test.expectedResult, obj.Result)
		})
	}
}
