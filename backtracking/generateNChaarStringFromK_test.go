package backtracking

import (
	"fmt"
	"golang-dsa/common"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NCharStringFromKTestSuite struct {
	suite.Suite
}

func TestNCharStringFromKTestSuite(t *testing.T) {
	suite.Run(t, &NCharStringFromKTestSuite{})
}

func (testSuite *NCharStringFromKTestSuite) TestCase() {
	tests := []struct {
		poolSize             int
		length               int
		expectedCombinations []string
	}{
		{
			poolSize:             2,
			length:               3,
			expectedCombinations: []string{"000", "100", "010", "110", "001", "101", "011", "111"},
		},
		{
			poolSize:             3,
			length:               2,
			expectedCombinations: []string{"00", "10", "20", "01", "11", "21", "02", "12", "22"},
		},
		{
			poolSize:             0,
			length:               3,
			expectedCombinations: []string{},
		},
		{
			poolSize:             10,
			length:               0,
			expectedCombinations: []string{},
		},
	}

	for _, test := range tests {
		testSuite.Run(fmt.Sprintf("Running for %v chars and %v poolsize", test.length, test.poolSize), func() {
			check := NewNCharStringFromK(test.poolSize, test.length)
			check.Execute()

			testSuite.Truef(common.CompareStringArrays(test.expectedCombinations, check.Combinations), "expected array to be same")
		})
	}
}
