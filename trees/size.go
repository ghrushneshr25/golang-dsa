package trees

import "container/list"

func size(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	return size(root.Left) + 1 + size(root.Right)
}

func nonRecursiveSize(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	var result int
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		qLen := queue.Len()
		for i := 0; i < qLen; i++ {
			node := queue.Remove(queue.Front()).(*BinaryTreeNode)
			result++
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return result
}
