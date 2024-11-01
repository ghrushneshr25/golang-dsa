package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GetMinimumTestSuite struct {
	suite.Suite
}

func TestGetMinimum(t *testing.T) {
	suite.Run(t, new(GetMinimumTestSuite))
}

func (suite *GetMinimumTestSuite) TestGetMinimum() {
	tests := []struct {
		input          []int
		expectedResult int
	}{
		{
			input:          []int{3, 5, 2, 1, 4},
			expectedResult: 1,
		},
		{
			input:          []int{10, 20, 30, 40, 50},
			expectedResult: 10,
		},
		{
			input:          []int{5, 4, 3, 2, 1},
			expectedResult: 1,
		},
		{
			input:          []int{1, 2, 3, 4, 5},
			expectedResult: 1,
		},
		{
			input:          []int{2, 2, 2, 2, 2},
			expectedResult: 2,
		},
		{
			input:          []int{3, 3, 1, 3, 3},
			expectedResult: 1,
		},
		{
			input:          []int{7, 8, 9, 1, 2, 3},
			expectedResult: 1,
		},
		{
			input:          []int{1},
			expectedResult: 1,
		},
		{
			input:          []int{2, 1},
			expectedResult: 1,
		},
		{
			input:          []int{1, 2},
			expectedResult: 1,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("GetMinimum %v", index+1)
		suite.Run(testName, func() {
			gm := NewGetMinimum(test.input)
			gm.Execute()
			suite.Equal(test.expectedResult, gm.Result)
		})
	}
}
