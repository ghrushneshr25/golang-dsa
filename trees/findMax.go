package trees

import (
	"container/list"
	"math"
)

/*
	Problem-1 Give an algorithm for finding maximum element in binary tree
*/

func findMax(root *BinaryTreeNode) int {
	max := math.MinInt32
	if root != nil {
		rootVal := root.Data
		leftMax := findMax(root.Left)
		rightMax := findMax(root.Right)

		if leftMax > rightMax {
			max = leftMax
		} else {
			max = rightMax
		}
		if rootVal > max {
			max = rootVal
		}
	}
	return max
}

/*
	Problem-2 Give an algorithm for finding the maximum element in binary tree without recursion.
*/

func nonRecursiveFindMax(root *BinaryTreeNode) int {
	max := math.MinInt32
	if root != nil {
		queue := list.New()
		queue.PushBack(root)
		for queue.Len() != 0 {
			qlen := queue.Len()
			for i := 0; i < qlen; i++ {
				node := queue.Remove(queue.Front()).(*BinaryTreeNode)
				if node.Data > max {
					max = node.Data
				}
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			}
		}
	}
	return max
}
