package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type IsPalindromeStackTestSuite struct {
	suite.Suite
}

func TestIsPalindromeStack(t *testing.T) {
	suite.Run(t, new(IsPalindromeStackTestSuite))
}

func (suite *IsPalindromeStackTestSuite) TestIsPalindrome() {
	tests := []struct {
		input          string
		expectedResult bool
	}{
		{
			input:          "abXba",
			expectedResult: true,
		},
		{
			input:          "abcXcbaa",
			expectedResult: false,
		},
		{
			input:          "abcXcba",
			expectedResult: true,
		},
		{
			input:          "abXbc",
			expectedResult: false,
		},
		{
			input:          "aXb",
			expectedResult: false,
		},
		{
			input:          "aXa",
			expectedResult: true,
		},
		{
			input:          "abccbaXabccba",
			expectedResult: true,
		},
		{
			input:          "abccbaXabccbb",
			expectedResult: false,
		},
		{
			input:          "X",
			expectedResult: true,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("IsPalindrome %v", index+1)
		suite.Run(testName, func() {
			obj := NewIsPalindromeStack(test.input)
			obj.Execute()
			suite.Equal(test.expectedResult, obj.Result)
		})
	}
}
