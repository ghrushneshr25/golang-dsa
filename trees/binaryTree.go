package trees

import (
	"container/list"
	"fmt"
	"math/rand"
)

type BinaryTreeNode struct {
	Left  *BinaryTreeNode
	Data  int
	Right *BinaryTreeNode
}

func PreOrderTraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Data)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func PreOrderWalk(root *BinaryTreeNode, ch chan int) {
	if root == nil {
		return
	}

	ch <- root.Data
	PreOrderWalk(root.Left, ch)
	PreOrderWalk(root.Right, ch)
}

func PreOrderWalker(root *BinaryTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		PreOrderWalk(root, ch)
		close(ch)
	}()
	return ch
}

func NewBinaryTree(n, k int) (root *BinaryTreeNode) {
	for _, v := range rand.Perm(n) {
		root = insert(root, (1+v)*k)
	}
	return
}

func insert(root *BinaryTreeNode, v int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{nil, v, nil}
	}

	if v < root.Data {
		root.Left = insert(root.Left, v)
		return root
	}

	root.Right = insert(root.Right, v)
	return root
}

func NonRecursivePreOrderTraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	current := root
	stack := list.New()

	stack.PushBack(current)
	fmt.Println()
	for stack.Len() > 0 {
		temp := stack.Remove(stack.Back())
		current = temp.(*BinaryTreeNode)

		fmt.Printf("%d ", current.Data)
		if current.Right != nil {
			stack.PushBack(current.Right)
		}
		if current.Left != nil {
			stack.PushBack(current.Left)
		}
	}
	fmt.Println("")
}

func InOrderTraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left)
	fmt.Printf("%d ", root.Data)
	InOrderTraversal(root.Right)
}

func NonRecursiveInOrderTraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	temp := root
	stack := list.New()

	for stack.Len() > 0 || temp != nil {
		if temp != nil {
			stack.PushBack(temp)
			temp = temp.Left
		} else {
			obj := stack.Remove(stack.Back())
			temp = obj.(*BinaryTreeNode)
			fmt.Printf("%d ", temp.Data)
			temp = temp.Right
		}
	}
}

func PostOrderTraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	PostOrderTraversal(root.Left)
	PostOrderTraversal(root.Right)
	fmt.Printf("%d ", root.Data)
}

func NonRecursivePostOrderTraversal(root *BinaryTreeNode) {
	var result []int
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		temp := stack.Remove(stack.Back()).(*BinaryTreeNode)

		result = append(result, temp.Data)
		if temp.Left != nil {
			stack.PushBack(temp.Left)
		}

		if temp.Right != nil {
			stack.PushBack(temp.Right)
		}
	}

	n := len(result)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		result[i], result[j] = result[j], result[i]
	}
	fmt.Println(result)
}

func LevelOrderTraversal(root *BinaryTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := list.New()

	for queue.Len() > 0 {
		currentLen := queue.Len()
		level := []int{}
		for i := 0; i < currentLen; i++ {
			node := queue.Remove(queue.Front())
			level = append(level, node.(*BinaryTreeNode).Data)

			if node.(*BinaryTreeNode).Left != nil {
				queue.PushBack(node.(*BinaryTreeNode).Left)
			}
			if node.(*BinaryTreeNode).Right != nil {
				queue.PushBack(node.(*BinaryTreeNode).Right)
			}
		}
		result = append(result, level)
	}
	return result
}
