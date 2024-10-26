package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type JosephusCircleTestSuite struct {
	suite.Suite
}

func TestJosephusCirclet(t *testing.T) {
	suite.Run(t, &JosephusCircleTestSuite{})
}

func (jc *JosephusCircleTestSuite) TestCase() {
	tests := []struct {
		N              int
		M              int
		expectedResult int
	}{
		{
			N:              7,
			M:              3,
			expectedResult: 5,
		},
		{
			N:              5,
			M:              2,
			expectedResult: 4,
		},
		{
			N:              10,
			M:              1,
			expectedResult: 1,
		},
		{
			N:              6,
			M:              7,
			expectedResult: 6,
		},
	}

	for _, test := range tests {
		jc.Run("TestCase", func() {
			obj := NewJosephusCircle(test.N, test.M)
			obj.Execute()
			jc.Equal(test.expectedResult, obj.Result)
		})
	}
}
