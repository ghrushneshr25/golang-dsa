package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AddNumbersOfTwoLinkedListTestSuite struct {
	suite.Suite
}

func TestAddNumbersOfTwoLinkedListSuite(t *testing.T) {
	suite.Run(t, &AddNumbersOfTwoLinkedListTestSuite{})
}

func (testSuite *AddNumbersOfTwoLinkedListTestSuite) TestCase() {
	tests := []struct {
		valOne         []int
		valTwo         []int
		expectedResult []int
	}{
		{
			valOne:         []int{3, 4, 3},
			valTwo:         []int{5, 6, 4},
			expectedResult: []int{8, 0, 8},
		},
		{
			valOne:         []int{1, 8},
			valTwo:         []int{0},
			expectedResult: []int{1, 8},
		},
		{
			valOne:         []int{0},
			valTwo:         []int{1, 8},
			expectedResult: []int{1, 8},
		},
		{
			valOne:         []int{9, 9, 9},
			valTwo:         []int{1},
			expectedResult: []int{0, 0, 0, 1},
		},
		{
			valOne:         []int{1},
			valTwo:         []int{9, 9, 9},
			expectedResult: []int{0, 0, 0, 1},
		},
		{
			valOne:         []int{2, 4, 3},
			valTwo:         []int{5, 6, 4},
			expectedResult: []int{7, 0, 8},
		},
		{
			valOne:         []int{0},
			valTwo:         []int{0},
			expectedResult: []int{0},
		},
		{
			valOne:         []int{1},
			valTwo:         []int{},
			expectedResult: []int{1},
		},
		{
			valOne:         []int{},
			valTwo:         []int{1},
			expectedResult: []int{1},
		},
	}

	for i, test := range tests {
		testname := fmt.Sprintf("TestCase %v", i)
		testSuite.Run(testname, func() {
			obj := NewAddNumbersOfTwoLinkedList(test.valOne, test.valTwo)
			obj.Execute()
			testSuite.Equal(test.expectedResult, obj.Result)
		})
	}
}
