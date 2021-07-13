package letcode

/*
合并二叉树
https://leetcode-cn.com/problems/merge-two-binary-trees/
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	newNode := &TreeNode{Val:root1.Val+root2.Val,Left:nil,Right:nil}
	newNode.Left = mergeTrees(root1.Left, root2.Left)
	newNode.Right = mergeTrees(root1.Right, root2.Right)

	return newNode
}

