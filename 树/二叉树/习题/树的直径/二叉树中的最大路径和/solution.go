package leetcode_124

import "math"

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/
// 难度：困难

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		sum1 := max(dfs(node.Left), 0)
		sum2 := max(dfs(node.Right), 0)

		maxSum = max(maxSum, node.Val+sum1+sum2)

		return node.Val + max(sum1, sum2)
	}
	dfs(root)

	return maxSum
}
