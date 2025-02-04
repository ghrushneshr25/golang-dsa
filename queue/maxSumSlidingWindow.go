/*
Maximum sum in sliding window: Given array A[] with sliding window of size w which is moving from
the very left of the array to the very right. Assume that we can only see the numbers in the window. Each
time the sliding window moves rightwards by one position.
*/

package queue

import (
	"container/list"
)

func MaxSumSlidingWindow(array []int, windowSize int) (maxSum int) {

	if len(array) == 0 || windowSize <= 0 || windowSize > len(array) {
		return
	}

	deque := list.New()
	currentSum := 0

	for i := 0; i < len(array); i++ {
		if deque.Len() > 0 && deque.Front().Value.(int) <= i-windowSize {
			currentSum -= array[deque.Remove(deque.Front()).(int)]
		}

		deque.PushBack(i)
		currentSum += array[i]

		if i >= windowSize-1 {
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}

	return
}
