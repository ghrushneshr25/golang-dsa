package backtracking

import (
	"fmt"
	"golang-dsa/common"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NBitStringTestSuite struct {
	suite.Suite
}

func TestNBitStringTestSuite(t *testing.T) {
	suite.Run(t, &NBitStringTestSuite{})
}

func (testSuite *NBitStringTestSuite) TestCase() {
	tests := []struct {
		size                 int
		expectedCombinations []string
	}{
		{
			size:                 5,
			expectedCombinations: []string{"00000", "00001", "00010", "00011", "00100", "00101", "00110", "00111", "01000", "01001", "01010", "01011", "01100", "01101", "01110", "01111", "10000", "10001", "10010", "10011", "10100", "10101", "10110", "10111", "11000", "11001", "11010", "11011", "11100", "11101", "11110", "11111"},
		},
		{
			size:                 3,
			expectedCombinations: []string{"000", "001", "010", "011", "100", "101", "110", "111"},
		},
		{
			size:                 0,
			expectedCombinations: []string{},
		},
	}

	for _, test := range tests {
		testSuite.Run(fmt.Sprintf("Running for %v bits", test.size), func() {
			nbit := NewNBitString(test.size)
			nbit.Execute()
			
			testSuite.Truef(common.CompareStringArrays(test.expectedCombinations, nbit.Combinations), "expected array to be same")
		})
	}

}
