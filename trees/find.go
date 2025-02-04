package trees

import "container/list"

func find(root *BinaryTreeNode, data int) *BinaryTreeNode {
	if root == nil {
		return root
	}
	if data == root.Data {
		return root
	}
	temp := find(root.Left, data)

	if temp != nil {
		return temp
	}
	return find(root.Right, data)
}

func nonRecursiveFind(root *BinaryTreeNode, data int) *BinaryTreeNode {
	if root == nil {
		return root
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		qLen := queue.Len()
		for i := 0; i < qLen; i++ {
			node := queue.Remove(queue.Front()).(*BinaryTreeNode)
			if node.Data == data {
				return node
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return nil
}
