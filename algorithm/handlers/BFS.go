package handlers

/*
BFS：深度优先
⼆叉树的最⼩⾼度

*/

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func MinDepth(root *TreeNode) int {
	queue := make([]*TreeNode, 0) // 保存当前要遍历的一层
	queue = append(queue, root)
	depth := 1
	for len(queue) > 0 {
		newQuene := make([]*TreeNode, 0)
		for _, cur := range queue {
			// 判断是否到终点
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			// 遍历当前节点的下一步的所有选择，写入到newQuene中
			if cur.Left != nil {
				newQuene = append(newQuene, cur.Left)
			}
			if cur.Right != nil {
				newQuene = append(newQuene, cur.Right)
			}
		}
		depth++
		queue = newQuene
	}
	return depth
}
