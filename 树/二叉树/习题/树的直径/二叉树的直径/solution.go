package leetcode_543

// https://leetcode.cn/problems/diameter-of-binary-tree/
// 难度：简单

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		d1 := dfs(node.Left)
		d2 := dfs(node.Right)
		diameter = max(diameter, d1+d2)

		return max(d1, d2) + 1
	}
	dfs(root)

	return diameter
}
