package recursion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TowerOfHanoiTestSuite struct {
	suite.Suite
}

func TestTowerOfHanoi(t *testing.T) {
	suite.Run(t, &TowerOfHanoiTestSuite{})
}

func (testSuite *TowerOfHanoiTestSuite) TestCases() {
	tests := []struct {
		disks         int
		expectedMoves []string
		expectedCount int
	}{
		{
			disks:         1,
			expectedMoves: []string{"A -> C"},
			expectedCount: 1,
		},
		{
			disks:         2,
			expectedMoves: []string{"A -> B", "A -> C", "B -> C"},
			expectedCount: 3,
		},
		{
			disks: 3,
			expectedMoves: []string{
				"A -> C", "A -> B", "C -> B", "A -> C", "B -> A", "B -> C", "A -> C",
			},
			expectedCount: 7,
		},
		{
			disks: 4,
			expectedMoves: []string{
				"A -> B", "A -> C", "B -> C", "A -> B", "C -> A", "C -> B", "A -> B",
				"A -> C", "B -> C", "B -> A", "C -> A", "B -> C", "A -> B", "A -> C", "B -> C",
			},
			expectedCount: 15,
		},
	}

	for _, test := range tests {
		testSuite.Run(fmt.Sprintf("Running for %v disks", test.disks), func() {
			toh := NewTowerOfHanoi(test.disks)
			toh.Execute()

			testSuite.Equalf(test.expectedCount, toh.MoveCount, "expected the count to be equal")
			testSuite.Equalf(test.expectedMoves, toh.Move, "expected moves are different")
		})
	}
}
