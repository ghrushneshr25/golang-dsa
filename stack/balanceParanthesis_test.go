package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type BalanceParanthesisTestSuite struct {
	suite.Suite
}

func TestBalanceParanthesis(t *testing.T) {
	suite.Run(t, new(BalanceParanthesisTestSuite))
}

func (suite *BalanceParanthesisTestSuite) TestIsBalanced() {
	tests := []struct {
		input          string
		expectedResult bool
	}{
		{
			input:          "()",
			expectedResult: true,
		},
		{
			input:          "({[]})",
			expectedResult: true,
		},
		{
			input:          "({[)]}",
			expectedResult: false,
		},
		{
			input:          "(()))",
			expectedResult: false,
		},
		{
			input:          "((())",
			expectedResult: false,
		},
		{
			input:          "",
			expectedResult: true,
		},
		{
			input:          "{[()]}",
			expectedResult: true,
		},
		{
			input:          "{[(])}",
			expectedResult: false,
		},
		{
			input:          "([{}])",
			expectedResult: true,
		},
		{
			input:          "([)]",
			expectedResult: false,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("IsBalanced %v", test.input)
		suite.Run(testName, func() {
			bp := NewBalanceParanthesis(test.input)
			bp.Execute()
			suite.Equal(test.expectedResult, bp.Result)
		})
	}
}
