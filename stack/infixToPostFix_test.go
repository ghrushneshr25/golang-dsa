package stack

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/suite"
)

type InfixToPostFixTestSuite struct {
    suite.Suite
}

func TestInfixToPostFix(t *testing.T) {
    suite.Run(t, new(InfixToPostFixTestSuite))
}

func (suite *InfixToPostFixTestSuite) TestConvert() {
    tests := []struct {
        input         string
        expectedResult string
    }{
        {
            input:         "a+b",
            expectedResult: "ab+",
        },
        {
            input:         "a+b*c",
            expectedResult: "abc*+",
        },
        {
            input:         "(a+b)*c",
            expectedResult: "ab+c*",
        },
        {
            input:         "a+b*(c^d-e)^(f+g*h)-i",
            expectedResult: "abcd^e-fgh*+^*+i-",
        },
        {
            input:         "A*(B+C)/D",
            expectedResult: "ABC+*D/",
        },
        {
            input:         "A+B*(C-D)",
            expectedResult: "ABCD-*+",
        },
        {
            input:         "",
            expectedResult: "",
        },
        {
            input:         "a+b*(c-d)",
            expectedResult: "abcd-*+",
        },
        {
            input:         "a+b+c",
            expectedResult: "ab+c+",
        },
        {
            input:         "a*b+c*d",
            expectedResult: "ab*cd*+",
        },
    }

    for index, test := range tests {
        testName := fmt.Sprintf("Convert %v", index+1)
        suite.Run(testName, func() {
            itp := NewInfixToPostFix(test.input)
            result := itp.Convert()
            suite.Equal(test.expectedResult, result)
        })
    }
}