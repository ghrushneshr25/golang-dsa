package recursion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SortedArrayTestSuite struct {
	suite.Suite
}

func TestSortedArray(t *testing.T) {
	suite.Run(t, &SortedArrayTestSuite{})
}

func (testSuite *SortedArrayTestSuite) TestCases() {
	tests := []struct {
		array          []int
		expectedResult bool
	}{
		{
			array:          []int{1, 2, 3, 4, 5, 6},
			expectedResult: true,
		},
		{
			array:          []int{1, 2, 3, 4, 5, 4},
			expectedResult: false,
		},
		{
			array:          []int{10, 2, 3, 4, 5, 6},
			expectedResult: false,
		},
		{
			array:          []int{1, 20, 3, 4, 5, 6},
			expectedResult: false,
		},
		{
			array:          []int{1, 2, 30, 4, 5, 6},
			expectedResult: false,
		},
		{
			array:          []int{-998, -957, -936, -926, -923, -912, -902, 
				-881, -833, -773, -768, -768, -765, -750, -742, -731, -693, 
				-665, -655, -571, -561, -521, -493, -480, -480, -462, -421, 
				-413, -381, -368, -338, -319, -298, -282, -275, -265, -256, 
				-245, -236, -235, -218, -205, -176, -176, -127, -91, -78, -61, 
				-50, -50, -44, -25, 3, 31, 50, 64, 74, 83, 104, 125, 131, 165, 
				178, 196, 206, 224, 258, 285, 292, 307, 327, 337, 344, 365, 385, 
				395, 448, 482, 490, 521, 521, 535, 573, 594, 601, 613, 667, 734, 
				738, 744, 747, 783, 791, 867, 869, 897, 930, 956, 983, 984},
			expectedResult: true,
		},
		{
			array: []int{-229, 657, 440, 561, -655, -317, -440, 14, 438, -248, 444,
				-834, -393, -197, -679, -953, -584, -111, 923, -458, -132, 62,
				729, -930, -773, -652, -298, -720, -937, 78, -753, -381, -327,
				884, -303, -516, 432, -206, -776, -704, -952, 311, 572, 176,
				609, 395, -878, 283, 509, -325},
			expectedResult: false,
		},
	}

	for i, test := range tests {
		testSuite.Run(fmt.Sprintf("Running test %v", i), func() {
			soa := NewSortedArray(test.array)
			soa.Execute()

			testSuite.Equalf(test.expectedResult, soa.Result, "expected the result differs")
		})
	}
}
