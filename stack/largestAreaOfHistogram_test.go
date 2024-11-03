package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LargestAreaOfHistogramTestSuite struct {
	suite.Suite
}

func TestLargestAreaOfHistogram(t *testing.T) {
	suite.Run(t, new(LargestAreaOfHistogramTestSuite))
}

func (suite *LargestAreaOfHistogramTestSuite) TestLargestArea() {
	tests := []struct {
		input          []int
		expectedResult int
	}{
		{
			input:          []int{6, 2, 5, 4, 5, 1, 6},
			expectedResult: 12,
		},
		{
			input:          []int{2, 1, 5, 6, 2, 3},
			expectedResult: 10,
		},
		{
			input:          []int{2, 4},
			expectedResult: 4,
		},
		{
			input:          []int{6, 2, 5, 4, 5, 1, 6, 3},
			expectedResult: 12,
		},
		{
			input:          []int{1, 2, 3, 4, 5},
			expectedResult: 9,
		},
		{
			input:          []int{5, 4, 3, 2, 1},
			expectedResult: 9,
		},
		{
			input:          []int{2, 1, 2},
			expectedResult: 3,
		},
		{
			input:          []int{2, 1, 2, 3, 1},
			expectedResult: 5,
		},
		{
			input:          []int{1, 1, 1, 1, 1},
			expectedResult: 5,
		},
		{
			input:          []int{},
			expectedResult: 0,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("LargestArea %v", index+1)
		suite.Run(testName, func() {
			obj := NewMaxAreaOfHistogram(test.input)
			obj.Execute()
			suite.Equal(test.expectedResult, obj.Result)
		})
	}
}
