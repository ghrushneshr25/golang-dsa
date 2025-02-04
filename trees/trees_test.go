package trees

import (
	"math"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TreesTestSuite struct {
	suite.Suite
}

func TestTreesSuite(t *testing.T) {
	suite.Run(t, new(TreesTestSuite))
}

func (suite *TreesTestSuite) TestFindMax() {
	tests := []struct {
		name string
		root *BinaryTreeNode
		want int
	}{
		{
			name: "single node",
			root: &BinaryTreeNode{Data: 10},
			want: 10,
		},
		{
			name: "left heavy tree",
			root: &BinaryTreeNode{
				Data:  10,
				Left:  &BinaryTreeNode{Data: 20},
				Right: &BinaryTreeNode{Data: 5},
			},
			want: 20,
		},
		{
			name: "right heavy tree",
			root: &BinaryTreeNode{
				Data:  10,
				Left:  &BinaryTreeNode{Data: 5},
				Right: &BinaryTreeNode{Data: 20},
			},
			want: 20,
		},
		{
			name: "balanced tree",
			root: &BinaryTreeNode{
				Data: 10,
				Left: &BinaryTreeNode{
					Data:  20,
					Left:  &BinaryTreeNode{Data: 30},
					Right: &BinaryTreeNode{Data: 25},
				},
				Right: &BinaryTreeNode{
					Data:  15,
					Left:  &BinaryTreeNode{Data: 5},
					Right: &BinaryTreeNode{Data: 35},
				},
			},
			want: 35,
		},
		{
			name: "empty tree",
			root: nil,
			want: math.MinInt32,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.want, findMax(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestNonRecursiveFindMax() {
	tests := []struct {
		name string
		root *BinaryTreeNode
		want int
	}{
		{
			name: "single node",
			root: &BinaryTreeNode{Data: 10},
			want: 10,
		},
		{
			name: "left heavy tree",
			root: &BinaryTreeNode{
				Data:  10,
				Left:  &BinaryTreeNode{Data: 20},
				Right: &BinaryTreeNode{Data: 5},
			},
			want: 20,
		},
		{
			name: "right heavy tree",
			root: &BinaryTreeNode{
				Data:  10,
				Left:  &BinaryTreeNode{Data: 5},
				Right: &BinaryTreeNode{Data: 20},
			},
			want: 20,
		},
		{
			name: "balanced tree",
			root: &BinaryTreeNode{
				Data: 10,
				Left: &BinaryTreeNode{
					Data:  20,
					Left:  &BinaryTreeNode{Data: 30},
					Right: &BinaryTreeNode{Data: 25},
				},
				Right: &BinaryTreeNode{
					Data:  15,
					Left:  &BinaryTreeNode{Data: 5},
					Right: &BinaryTreeNode{Data: 35},
				},
			},
			want: 35,
		},
		{
			name: "empty tree",
			root: nil,
			want: math.MinInt32,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.want, nonRecursiveFindMax(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestFind() {
	root := &BinaryTreeNode{
		Data: 1,
		Left: &BinaryTreeNode{
			Data: 2,
			Left: &BinaryTreeNode{
				Data: 4,
			},
			Right: &BinaryTreeNode{
				Data: 5,
			},
		},
		Right: &BinaryTreeNode{
			Data: 3,
			Left: &BinaryTreeNode{
				Data: 6,
			},
			Right: &BinaryTreeNode{
				Data: 7,
			},
		},
	}

	tests := []struct {
		data     int
		expected *BinaryTreeNode
	}{
		{data: 1, expected: root},
		{data: 4, expected: root.Left.Left},
		{data: 7, expected: root.Right.Right},
		{data: 8, expected: nil},
	}

	for _, test := range tests {
		result := find(root, test.data)
		suite.Equal(test.expected, result, "find(%d) = %v; expected %v", test.data, result, test.expected)
	}
}

func (suite *TreesTestSuite) TestNonRecursiveFind() {
	root := &BinaryTreeNode{
		Data: 1,
		Left: &BinaryTreeNode{
			Data: 2,
			Left: &BinaryTreeNode{
				Data: 4,
			},
			Right: &BinaryTreeNode{
				Data: 5,
			},
		},
		Right: &BinaryTreeNode{
			Data: 3,
			Left: &BinaryTreeNode{
				Data: 6,
			},
			Right: &BinaryTreeNode{
				Data: 7,
			},
		},
	}

	tests := []struct {
		data     int
		expected *BinaryTreeNode
	}{
		{data: 1, expected: root},
		{data: 4, expected: root.Left.Left},
		{data: 7, expected: root.Right.Right},
		{data: 8, expected: nil},
	}

	for _, test := range tests {
		result := nonRecursiveFind(root, test.data)
		suite.Equal(test.expected, result, "find(%d) = %v; expected %v", test.data, result, test.expected)
	}
}

func (suite *TreesTestSuite) TestSize() {
	tests := []struct {
		name string
		root *BinaryTreeNode
		want int
	}{
		{"empty tree", nil, 0},
		{"single node", &BinaryTreeNode{}, 1},
		{"three nodes", &BinaryTreeNode{
			Left:  &BinaryTreeNode{},
			Right: &BinaryTreeNode{},
		}, 3},
		{"complex tree", &BinaryTreeNode{
			Left: &BinaryTreeNode{
				Left: &BinaryTreeNode{},
			},
			Right: &BinaryTreeNode{
				Right: &BinaryTreeNode{},
			},
		}, 5},
		{"more complex tree", &BinaryTreeNode{
			Left: &BinaryTreeNode{
				Left: &BinaryTreeNode{
					Left: &BinaryTreeNode{},
				},
				Right: &BinaryTreeNode{},
			},
			Right: &BinaryTreeNode{
				Left: &BinaryTreeNode{},
				Right: &BinaryTreeNode{
					Right: &BinaryTreeNode{},
				},
			},
		}, 9},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.want, size(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestNonRecursiveSize() {
	tests := []struct {
		name string
		root *BinaryTreeNode
		want int
	}{
		{"empty tree", nil, 0},
		{"single node", &BinaryTreeNode{}, 1},
		{"three nodes", &BinaryTreeNode{
			Left:  &BinaryTreeNode{},
			Right: &BinaryTreeNode{},
		}, 3},
		{"complex tree", &BinaryTreeNode{
			Left: &BinaryTreeNode{
				Left: &BinaryTreeNode{},
			},
			Right: &BinaryTreeNode{
				Right: &BinaryTreeNode{},
			},
		}, 5},
		{"more complex tree", &BinaryTreeNode{
			Left: &BinaryTreeNode{
				Left: &BinaryTreeNode{
					Left: &BinaryTreeNode{},
				},
				Right: &BinaryTreeNode{},
			},
			Right: &BinaryTreeNode{
				Left: &BinaryTreeNode{},
				Right: &BinaryTreeNode{
					Right: &BinaryTreeNode{},
				},
			},
		}, 9},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.want, nonRecursiveSize(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestLevelOrderBottomUp() {
	tests := []struct {
		name     string
		root     *BinaryTreeNode
		expected [][]int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: [][]int{},
		},
		{
			name: "Single node",
			root: &BinaryTreeNode{Data: 1},
			expected: [][]int{
				{1},
			},
		},
		{
			name: "Multiple levels",
			root: &BinaryTreeNode{
				Data: 1,
				Left: &BinaryTreeNode{
					Data:  2,
					Left:  &BinaryTreeNode{Data: 4},
					Right: &BinaryTreeNode{Data: 5},
				},
				Right: &BinaryTreeNode{
					Data:  3,
					Left:  &BinaryTreeNode{Data: 6},
					Right: &BinaryTreeNode{Data: 7},
				},
			},
			expected: [][]int{
				{7, 6, 5, 4},
				{3, 2},
				{1},
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.expected, LevelOrderBottomUp(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestHeightOfTree() {
	tests := []struct {
		name     string
		root     *BinaryTreeNode
		expected int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node",
			root:     &BinaryTreeNode{Data: 1},
			expected: 1,
		},
		{
			name: "Multiple levels",
			root: &BinaryTreeNode{
				Data: 1,
				Left: &BinaryTreeNode{
					Data: 2,
					Left: &BinaryTreeNode{Data: 4},
					Right: &BinaryTreeNode{
						Data: 5,
						Right: &BinaryTreeNode{
							Data: 6,
						},
					},
				},
				Right: &BinaryTreeNode{
					Data:  3,
					Left:  &BinaryTreeNode{Data: 6},
					Right: &BinaryTreeNode{Data: 7},
				},
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.expected, HeightOfTree(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestNonRecursiveHeightOfTree() {
	tests := []struct {
		name     string
		root     *BinaryTreeNode
		expected int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node",
			root:     &BinaryTreeNode{Data: 1},
			expected: 1,
		},
		{
			name: "Multiple levels",
			root: &BinaryTreeNode{
				Data: 1,
				Left: &BinaryTreeNode{
					Data: 2,
					Left: &BinaryTreeNode{Data: 4},
					Right: &BinaryTreeNode{
						Data: 5,
						Right: &BinaryTreeNode{
							Data: 6,
						},
					},
				},
				Right: &BinaryTreeNode{
					Data:  3,
					Left:  &BinaryTreeNode{Data: 6},
					Right: &BinaryTreeNode{Data: 7},
				},
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.expected, HeightOfTree(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestDeepest() {
	tests := []struct {
		name     string
		root     *BinaryTreeNode
		expected *BinaryTreeNode
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: nil,
		},
		{
			name:     "Single node",
			root:     &BinaryTreeNode{Data: 1},
			expected: &BinaryTreeNode{Data: 1},
		},
		{
			name: "Multiple levels",
			root: &BinaryTreeNode{
				Data: 1,
				Left: &BinaryTreeNode{
					Data:  2,
					Left:  &BinaryTreeNode{Data: 4},
					Right: &BinaryTreeNode{Data: 5},
				},
				Right: &BinaryTreeNode{
					Data: 3,
					Left: &BinaryTreeNode{Data: 6},
					Right: &BinaryTreeNode{Data: 7,
						Right: &BinaryTreeNode{Data: 9},
					},
				},
			},
			expected: &BinaryTreeNode{Data: 9},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.expected, Deepest(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestLeavesCount() {
	tests := []struct {
		name     string
		root     *BinaryTreeNode
		expected int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node",
			root:     &BinaryTreeNode{Data: 1},
			expected: 1,
		},
		{
			name: "Multiple levels",
			root: &BinaryTreeNode{
				Data: 1,
				Left: &BinaryTreeNode{
					Data:  2,
					Left:  &BinaryTreeNode{Data: 4},
					Right: &BinaryTreeNode{Data: 5},
				},
				Right: &BinaryTreeNode{
					Data: 3,
					Left: &BinaryTreeNode{Data: 6},
					Right: &BinaryTreeNode{Data: 7,
						Right: &BinaryTreeNode{Data: 9},
					},
				},
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.expected, LeavesCount(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestFullNodesCount() {
	tests := []struct {
		name     string
		root     *BinaryTreeNode
		expected int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node",
			root:     &BinaryTreeNode{Data: 1},
			expected: 0,
		},
		{
			name: "Multiple levels",
			root: &BinaryTreeNode{
				Data: 1,
				Left: &BinaryTreeNode{
					Data:  2,
					Left:  &BinaryTreeNode{Data: 4},
					Right: &BinaryTreeNode{Data: 5},
				},
				Right: &BinaryTreeNode{
					Data: 3,
					Left: &BinaryTreeNode{Data: 6},
					Right: &BinaryTreeNode{Data: 7,
						Right: &BinaryTreeNode{Data: 9},
					},
				},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.expected, FullNodesCount(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestHalfNodesCount() {
	tests := []struct {
		name     string
		root     *BinaryTreeNode
		expected int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node",
			root:     &BinaryTreeNode{Data: 1},
			expected: 0,
		},
		{
			name: "Multiple levels",
			root: &BinaryTreeNode{
				Data: 1,
				Left: &BinaryTreeNode{
					Data:  2,
					Left:  &BinaryTreeNode{Data: 4},
					Right: &BinaryTreeNode{Data: 5},
				},
				Right: &BinaryTreeNode{
					Data: 3,
					Left: &BinaryTreeNode{Data: 6},
					Right: &BinaryTreeNode{Data: 7,
						Right: &BinaryTreeNode{Data: 9},
					},
				},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.expected, HalfNodesCount(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestMaxSumLevel() {
	tests := []struct {
		name     string
		root     *BinaryTreeNode
		expected int
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node",
			root:     &BinaryTreeNode{Data: 1},
			expected: 1,
		},
		{
			name: "Multiple levels",
			root: &BinaryTreeNode{
				Data: 1,
				Left: &BinaryTreeNode{
					Data:  2,
					Left:  &BinaryTreeNode{Data: 4},
					Right: &BinaryTreeNode{Data: 5},
				},
				Right: &BinaryTreeNode{
					Data: 3,
					Left: &BinaryTreeNode{Data: 6},
					Right: &BinaryTreeNode{Data: 7,
						Right: &BinaryTreeNode{Data: 9},
					},
				},
			},
			expected: 22, // Level with nodes 4, 5, 6, 7 has the maximum sum
		},
		{
			name: "Negative numbers",
			root: &BinaryTreeNode{
				Data: -1,
				Left: &BinaryTreeNode{
					Data:  -2,
					Left:  &BinaryTreeNode{Data: -4},
					Right: &BinaryTreeNode{Data: -5},
				},
				Right: &BinaryTreeNode{
					Data: -3,
					Left: &BinaryTreeNode{Data: -6},
					Right: &BinaryTreeNode{Data: -7,
						Right: &BinaryTreeNode{Data: 10},
					},
				},
			},
			expected: 10, // Level with nodes -4, -5, -6, -7 has the maximum (least negative) sum
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.expected, MaxSumLevel(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestBinaryTreePath() {
	tests := []struct {
		name     string
		root     *BinaryTreeNode
		expected []string
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: []string{},
		},
		{
			name:     "Single node",
			root:     &BinaryTreeNode{Data: 1},
			expected: []string{"1"},
		},
		{
			name: "Multiple levels",
			root: &BinaryTreeNode{
				Data: 1,
				Left: &BinaryTreeNode{
					Data:  2,
					Left:  &BinaryTreeNode{Data: 4},
					Right: &BinaryTreeNode{Data: 5},
				},
				Right: &BinaryTreeNode{
					Data: 3,
					Left: &BinaryTreeNode{Data: 6},
					Right: &BinaryTreeNode{Data: 7,
						Right: &BinaryTreeNode{Data: 9},
					},
				},
			},
			expected: []string{
				"1->2->4",
				"1->2->5",
				"1->3->6",
				"1->3->7->9",
			},
		},
		{
			name: "Negative numbers",
			root: &BinaryTreeNode{
				Data: -1,
				Left: &BinaryTreeNode{
					Data:  -2,
					Left:  &BinaryTreeNode{Data: -4},
					Right: &BinaryTreeNode{Data: -5},
				},
				Right: &BinaryTreeNode{
					Data: -3,
					Left: &BinaryTreeNode{Data: -6},
					Right: &BinaryTreeNode{Data: -7,
						Right: &BinaryTreeNode{Data: -9},
					},
				},
			},
			expected: []string{
				"-1->-2->-4",
				"-1->-2->-5",
				"-1->-3->-6",
				"-1->-3->-7->-9",
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			suite.Equal(tt.expected, BinaryTreePath(tt.root))
		})
	}
}

func (suite *TreesTestSuite) TestVerticalSumOfBinaryTree() {
	testCases := []struct {
		name     string
		root     *BinaryTreeNode
		expected []int
	}{
		{
			name: "multiple levels",
			root: &BinaryTreeNode{
				Data: 1,
				Left: &BinaryTreeNode{
					Data: 2,
					Left: &BinaryTreeNode{
						Data: 4,
					},
					Right: &BinaryTreeNode{
						Data: 5,
					},
				},
				Right: &BinaryTreeNode{
					Data: 3,
					Left: &BinaryTreeNode{
						Data: 6,
					},
					Right: &BinaryTreeNode{
						Data: 7,
					},
				},
			},
			expected: []int{4, 2, 12, 3, 7},
		},
	}

	for _, testCase := range testCases {
		suite.Run(testCase.name, func() {
			suite.Equal(testCase.expected, VerticalSumOfBinaryTree(testCase.root))
		})
	}
}

func (suite *TreesTestSuite) TestBuildTreeFromPreorder() {
	testCases := []struct {
		name     string
		input    string
		expected *BinaryTreeNode
	}{
		{
			name:  "multiple levels",
			input: "ILILL",
			expected: &BinaryTreeNode{
				Data: 1,
				Left: &BinaryTreeNode{
					Data: 0,
				},
				Right: &BinaryTreeNode{
					Data: 1,
					Left: &BinaryTreeNode{
						Data: 0,
					},
					Right: &BinaryTreeNode{
						Data: 0,
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		suite.Equal(testCase.expected, BuildTreeFromPreorder(testCase.input))
	}

}
