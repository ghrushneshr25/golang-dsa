package common

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CommonTestSuite struct {
	suite.Suite
}

func TestCommonTestSuite(t *testing.T) {
	suite.Run(t, &CommonTestSuite{})
}

func (suite *CommonTestSuite) TestCompareStringArrays() {
	tests := []struct {
		arr1     []string
		arr2     []string
		expected bool
	}{
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}, true},
		{[]string{"a", "b", "c"}, []string{"a", "b", "b"}, false},
		{[]string{"a", "b", "c", "d"}, []string{"a", "b", "d", "w"}, false},
		{[]string{"a", "a", "a"}, []string{"a", "a", "a"}, true},
		{[]string{"a", "b"}, []string{"a", "b", "c"}, false},
		{[]string{}, []string{}, true},
		{[]string{"a"}, []string{"a", "a"}, false},
	}

	for _, test := range tests {
		result := CompareStringArrays(test.arr1, test.arr2)
		suite.Equalf(test.expected, result, "CompareStringArrays(%v, %v) = %v; expected %v", test.arr1, test.arr2, result, test.expected)
	}
}
