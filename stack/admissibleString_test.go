package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AdmissibleStringTestSuite struct {
	suite.Suite
}

func TestAdmissibleString(t *testing.T) {
	suite.Run(t, new(AdmissibleStringTestSuite))
}

func (suite *AdmissibleStringTestSuite) TestIsAdmissible() {
	tests := []struct {
		input          string
		expectedResult bool
	}{
		{
			input:          "SSPP",
			expectedResult: true,
		},
		{
			input:          "SSPPSP",
			expectedResult: true,
		},
		{
			input:          "SSSSPPPP",
			expectedResult: true,
		},
		{
			input:          "SSSPPP",
			expectedResult: true,
		},
		{
			input:          "SSSPPPP",
			expectedResult: false,
		},
		{
			input:          "SP",
			expectedResult: true,
		},
		{
			input:          "P",
			expectedResult: false,
		},
		{
			input:          "SSP",
			expectedResult: true,
		},
		{
			input:          "SSPPPP",
			expectedResult: false,
		},
		{
			input:          "",
			expectedResult: true,
		},
	}

	for index, test := range tests {
		testName := fmt.Sprintf("IsAdmissible %v", index+1)
		suite.Run(testName, func() {
			obj := NewAdmissibleString(test.input)
			obj.Execute()
			suite.Equal(test.expectedResult, obj.Result)
		})
	}
}
