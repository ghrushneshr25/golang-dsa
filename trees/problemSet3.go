package trees

import (
	"container/list"
	"math"
	"strconv"
)

func LevelOrderBottomUp(root *BinaryTreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result = make([][]int, 0)

	queue := list.New()
	queue.PushBack(root)

	stack := list.New()

	for queue.Len() > 0 {
		qLen := queue.Len()
		var level []int

		for i := 0; i < qLen; i++ {
			node := queue.Remove(queue.Front()).(*BinaryTreeNode)
			level = append(level, node.Data)

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
		}

		stack.PushBack(level)
	}

	for stack.Len() > 0 {
		result = append(result, stack.Remove(stack.Back()).([]int))
	}

	return result
}

func HeightOfTree(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := HeightOfTree(root.Left)
	rightHeight := HeightOfTree(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func NonRecursiveHeightOfTree(root *BinaryTreeNode) (count int) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		qLen := queue.Len()

		for i := 0; i < qLen; i++ {
			node := queue.Remove(queue.Front()).(*BinaryTreeNode)

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
		}
		count++
	}
	return
}

func Deepest(root *BinaryTreeNode) (node *BinaryTreeNode) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		qLen := queue.Len()
		for i := 0; i < qLen; i++ {
			node = queue.Remove(queue.Front()).(*BinaryTreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return
}

func LeavesCount(root *BinaryTreeNode) (count int) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		qLen := queue.Len()
		for i := 0; i < qLen; i++ {
			node := queue.Remove(queue.Front()).(*BinaryTreeNode)
			if node.Right == nil && node.Left == nil {
				count++
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
		}
	}
	return
}

func FullNodesCount(root *BinaryTreeNode) (count int) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		qLen := queue.Len()
		for i := 0; i < qLen; i++ {
			node := queue.Remove(queue.Front()).(*BinaryTreeNode)
			if node.Right != nil && node.Left != nil {
				count++
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
		}
	}
	return
}

func HalfNodesCount(root *BinaryTreeNode) (count int) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		qLen := queue.Len()
		for i := 0; i < qLen; i++ {
			node := queue.Remove(queue.Front()).(*BinaryTreeNode)
			if (node.Right != nil && node.Left == nil) || (node.Right == nil && node.Left != nil) {
				count++
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
		}
	}
	return
}

func CompareStructure(root1, root2 *BinaryTreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	return CompareStructure(root1.Left, root2.Left) && CompareStructure(root1.Right, root2.Right)
}

func DiameterOfBinaryTree(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	var Diameter func(root *BinaryTreeNode, diameter *int) int
	Diameter = func(root *BinaryTreeNode, diameter *int) int {
		if root == nil {
			return 0
		}
		leftDepth := Diameter(root.Left, diameter)
		rightDepth := Diameter(root.Right, diameter)

		if leftDepth+rightDepth > *diameter {
			*diameter = leftDepth + rightDepth
		}
		return max(leftDepth, rightDepth) + 1
	}

	var diameter int
	Diameter(root, &diameter)
	return diameter
}

func MaxSumLevel(root *BinaryTreeNode) (maxSum int) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	maxSum = root.Data
	for queue.Len() > 0 {
		qLen := queue.Len()
		currentSum := 0
		for i := 0; i < qLen; i++ {
			node := queue.Remove(queue.Front()).(*BinaryTreeNode)
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			currentSum += node.Data
		}
		maxSum = max(maxSum, currentSum)
	}
	return
}

func BinaryTreePath(root *BinaryTreeNode) []string {
	result := []string{}

	var paths func(root *BinaryTreeNode, path string, result *[]string)
	paths = func(root *BinaryTreeNode, path string, result *[]string) {
		if root == nil {
			return
		}
		if path == "" {
			path = strconv.Itoa(root.Data)
		} else {
			path = path + "->" + strconv.Itoa(root.Data)
		}

		if root.Left == nil && root.Right == nil {
			*result = append(*result, path)
		}
		paths(root.Left, path, result)
		paths(root.Right, path, result)
	}

	paths(root, "", &result)
	return result
}

func HasPathSum(root *BinaryTreeNode, sum int) bool {
	allSums := make([]int, 0)

	var getAllSums func(root *BinaryTreeNode, allSums *[]int, currSum int)
	getAllSums = func(root *BinaryTreeNode, allSums *[]int, currSum int) {
		if root != nil {
			currSum += root.Data
			if root.Left == nil && root.Right == nil {
				*allSums = append(*allSums, currSum)
				return
			}
			getAllSums(root.Left, allSums, currSum)
			getAllSums(root.Right, allSums, currSum)
		}
	}

	getAllSums(root, &allSums, 0)
	for _, currentSum := range allSums {
		if sum == currentSum {
			return true
		}
	}

	return false
}

func InvertTree(root *BinaryTreeNode) *BinaryTreeNode {
	if root != nil {
		root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	}
	return root
}

func CheckIfMirror(root1, root2 *BinaryTreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Data != root2.Data {
		return false
	}

	return CheckIfMirror(root1.Left, root2.Right) && CheckIfMirror(root1.Right, root2.Left)
}

func LCA(root *BinaryTreeNode, alpha, beta int) *BinaryTreeNode {
	if root == nil {
		return nil
	}

	if root.Data == alpha || root.Data == beta {
		return root
	}

	left := LCA(root.Left, alpha, beta)
	right := LCA(root.Right, alpha, beta)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func BuildBinaryTree(preOrder []int, inOrder []int) *BinaryTreeNode {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}

	find := func(A []int, target int) int {
		for i, x := range A {
			if x == target {
				return i
			}
		}
		return -1
	}

	inOrderInfix := find(inOrder, preOrder[0])
	left := BuildBinaryTree(preOrder[1:inOrderInfix+1], inOrder[:inOrderInfix])
	right := BuildBinaryTree(preOrder[inOrderInfix+1:], inOrder[inOrderInfix+1:])

	return &BinaryTreeNode{
		Data:  preOrder[0],
		Left:  left,
		Right: right,
	}
}

func ZigZagTraversal(root *BinaryTreeNode) (res [][]int) {
	if root == nil {
		return [][]int{}
	}

	queue := list.New()
	queue.PushBack(root)

	leftToRight := false
	for {
		qLen := queue.Len()
		if qLen == 0 {
			break
		}
		var levelData []int
		for i := qLen - 1; i >= 0; i-- {
			node := queue.Remove(queue.Front()).(*BinaryTreeNode)
			levelData = append(levelData, node.Data)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		if leftToRight {
			for i := 0; i < len(levelData)/2; i++ {
				levelData[i], levelData[len(levelData)-i-1] = levelData[len(levelData)-i-1], levelData[i]
			}
		}
		res = append(res, levelData)
		leftToRight = !leftToRight
	}

	return
}

func ZigZagTraversalApproach2(root *BinaryTreeNode) (res [][]int) {
	if root == nil {
		return [][]int{}
	}

	queue := list.New()
	queue.PushBack(root)

	level := 0

	for queue.Len() > 0 {
		size := queue.Len()
		res = append(res, []int{})
		for i := 0; i < size; i++ {
			current := queue.Remove(queue.Front()).(*BinaryTreeNode)
			if level%2 == 0 {
				res[level] = append(res[level], current.Data)
			} else {
				res[level] = append([]int{current.Data}, res[level]...)
			}

			if current.Left != nil {
				queue.PushBack(current.Left)
			}

			if current.Right != nil {
				queue.PushBack(current.Right)
			}
		}
		level++
	}

	return
}

func VerticalSumOfBinaryTree(root *BinaryTreeNode) []int {

	hashMap := map[int]int{}
	var recursiveFunction func(root *BinaryTreeNode, distance int, hashMap map[int]int)

	recursiveFunction = func(root *BinaryTreeNode, distance int, hashMap map[int]int) {
		if root == nil {
			return
		}
		if val, ok := hashMap[distance]; !ok {
			hashMap[distance] = root.Data
		} else {
			hashMap[distance] = root.Data + val
		}
		recursiveFunction(root.Left, distance-1, hashMap)
		recursiveFunction(root.Right, distance+1, hashMap)
	}

	recursiveFunction(root, 0, hashMap)
	min, max := math.MaxInt32, math.MinInt32
	for key, _ := range hashMap {
		if key < min {
			min = key
		}
		if key > max {
			max = key
		}
	}
	result := []int{}
	for i := min; i <= max; i++ {
		result = append(result, hashMap[i])
	}
	return result
}

func BuildTreeFromPreorder(preOrder string) *BinaryTreeNode {
	var builder func(preOrder string, i *int) *BinaryTreeNode
	getType := func(b byte) int {
		if b == 'L' {
			return 0
		}
		return 1
	}
	builder = func(preOrder string, i *int) (newNode *BinaryTreeNode) {
		newNode = &BinaryTreeNode{
			Data: getType(preOrder[*i]),
		}
		if len(preOrder) == *i-1 {
			return nil
		}
		if preOrder[*i] == 'L' {
			return newNode
		}
		*i = *i + 1
		newNode.Left = builder(preOrder, i)
		*i = *i + 1
		newNode.Right = builder(preOrder, i)
		return newNode
	}

	i := 0
	return builder(preOrder, &i)
}
