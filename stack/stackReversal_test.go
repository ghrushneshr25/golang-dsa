package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StackReversalTestSuite struct {
	suite.Suite
}

func TestStackReversal(t *testing.T) {
	suite.Run(t, new(StackReversalTestSuite))
}

func (suite *StackReversalTestSuite) SetupTest() {
	// Setup code if needed
}

func (suite *StackReversalTestSuite) TestReverseStack() {
	tests := []struct {
		input          []int
		expectedResult []int
	}{
		{
			input:          []int{5, 4, 3, 2, 1},
			expectedResult: []int{5, 4, 3, 2, 1},
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
		testName := fmt.Sprintf("ReverseStack %v", index+1)
		suite.Run(testName, func() {
			sr := NewStackReversal(test.input)
			sr.Execute()
			suite.Equal(test.expectedResult, sr.Result)
		})
	}
}
