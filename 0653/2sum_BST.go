package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	BSTree := createTree()
	ans := findTarget(BSTree, 9)
	fmt.Println("Ans: ", ans)
}

func createTree() *TreeNode {
	num1 := &TreeNode{2, nil, nil}
	num2 := &TreeNode{3, nil, nil}
	num3 := &TreeNode{4, nil, nil}
	num4 := &TreeNode{5, nil, nil}
	num5 := &TreeNode{6, nil, nil}
	num6 := &TreeNode{7, nil, nil}
	num4.Left = num2
	num4.Right = num5
	num2.Left = num1
	num2.Right = num3
	num5.Right = num6
	root := num4
	return root
}

func findTarget(root *TreeNode, k int) bool {
	sumMap := make(map[int]bool)
	return check(root, k, sumMap)
}

func check(root *TreeNode, k int, sumMap map[int]bool) bool {
	if root == nil {
		return false
	}
	isExist := sumMap[root.Val]
	if isExist {
		return true
	}
	sumMap[k-root.Val] = true
	return check(root.Left, k, sumMap) || check(root.Right, k, sumMap)
}
