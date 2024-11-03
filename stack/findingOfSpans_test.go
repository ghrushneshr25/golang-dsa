package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FindingOfSpansTestSuite struct {
	suite.Suite
}

func TestFindingOfSpans(t *testing.T) {
	suite.Run(t, new(FindingOfSpansTestSuite))
}

func (suite *FindingOfSpansTestSuite) TestComputeSpans() {
	tests := []struct {
		input          []int
		expectedResult []int
	}{
		{
			input:          []int{6, 3, 4, 5, 2},
			expectedResult: []int{1, 1, 2, 3, 1},
		},
		{
			input:          []int{10, 4, 5, 90, 120, 80},
			expectedResult: []int{1, 1, 2, 4, 5, 1},
		},
		{
			input:          []int{1, 2, 3, 4, 5},
			expectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			input:          []int{5, 4, 3, 2, 1},
			expectedResult: []int{1, 1, 1, 1, 1},
		},
		{
			input:          []int{2, 2, 2, 2, 2},
			expectedResult: []int{1, 2, 3, 4, 5},
		},
		{
			input:          []int{1},
			expectedResult: []int{1},
		},
		{
			input:          []int{},
			expectedResult: []int{},
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("ComputeSpans %v", index+1)
		suite.Run(testName, func() {
			obj := NewFindingOfSpans(test.input)
			obj.Execute()
			suite.Equal(test.expectedResult, obj.Result)
		})
	}
}
