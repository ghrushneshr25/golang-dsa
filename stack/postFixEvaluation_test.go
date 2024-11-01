package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PostFixEvaluationTestSuite struct {
	suite.Suite
}

func TestPostFixEvaluation(t *testing.T) {
	suite.Run(t, new(PostFixEvaluationTestSuite))
}

func (suite *PostFixEvaluationTestSuite) TestEvaluate() {
	tests := []struct {
		input          string
		expectedResult int
	}{
		{
			input:          "231*+9-",
			expectedResult: -4,
		},
		{
			input:          "123+*8-",
			expectedResult: -3,
		},
		{
			input:          "34+2*1+",
			expectedResult: 15,
		},
		{
			input:          "56+78+*",
			expectedResult: 165,
		},
		{
			input:          "12+34+*",
			expectedResult: 21,
		},
		{
			input:          "123*+45*+",
			expectedResult: 27,
		},
		{
			input:          "92-3*",
			expectedResult: 21,
		},
		{
			input:          "82/",
			expectedResult: 4,
		},
		{
			input:          "82/4*",
			expectedResult: 16,
		},
		{
			input:          "82/4*5+",
			expectedResult: 21,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("Evaluate %v", index+1)
		suite.Run(testName, func() {
			pfe := NewPostFixEvaluation(test.input)
			pfe.Execute()
			suite.Equal(test.expectedResult, test.expectedResult)
		})
	}
}
