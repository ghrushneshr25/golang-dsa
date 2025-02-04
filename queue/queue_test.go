package queue

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type QueueTestSuite struct {
	suite.Suite
}

func (suite *QueueTestSuite) Test_MaxSumSlidingWindow() {
	tests := []struct {
		array      []int
		windowSize int
		expected   int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 27},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 2, 19},
		{[]int{1, 2, 3, 4, 5}, 1, 5},
		{[]int{1, 2, 3, 4, 5}, 5, 15},
		{[]int{1, 2, 3, 4, 5}, 0, 0},
		{[]int{}, 3, 0},
	}

	for i, test := range tests {
		suite.Run(fmt.Sprintf("Test %v Max Sum Sliding Window", i), func() {
			result := MaxSumSlidingWindow(test.array, test.windowSize)
			suite.Equalf(test.expected, result, "For array %v and window size %d", test.array, test.windowSize)
		})
	}
}

func (suite *QueueTestSuite) Test_InterLeavingQueue() {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 4, 2, 5, 3, 6}},
		{[]int{1, 2, 3, 4}, []int{1, 3, 2, 4}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{}, []int{}},
	}

	for i, test := range tests {
		suite.Run(fmt.Sprintf("Test %v InterLeaving Queue", i), func() {
			q := list.New()
			for _, val := range test.input {
				q.PushBack(val)
			}

			InterLeavingQueue(q)

			result := make([]int, 0, q.Len())
			for e := q.Front(); e != nil; e = e.Next() {
				result = append(result, e.Value.(int))
			}

			suite.Equalf(test.expected, result, "For input %v, expected %v but got %v", test.input, test.expected, result)
		})
	}
}

func (suite *QueueTestSuite) Test_ReverseFirstKElements() {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		{[]int{}, 3, []int{}},
	}

	for i, test := range tests {
		suite.Run(fmt.Sprintf("test %v", i), func() {
			q := list.New()
			for _, v := range test.input {
				q.PushBack(v)
			}

			ReverseFirstKElements(q, test.k)
			result := []int{}
			for e := q.Front(); e != nil; e = e.Next() {
				result = append(result, e.Value.(int))
			}

			suite.Equalf(test.expected, result, "For input %v and k=%d, expected %v but got %v", test.input, test.k, test.expected, result)
		})
	}
}

func TestQueueTestSuite(t *testing.T) {
	suite.Run(t, new(QueueTestSuite))
}
