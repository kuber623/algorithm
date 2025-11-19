package leetcode_101

// https://leetcode.cn/problems/symmetric-tree/
// 难度：简单

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(p, q *TreeNode) bool
	dfs = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil || p.Val != q.Val {
			return false
		}
		return dfs(p.Right, q.Left) && dfs(p.Left, q.Right)
	}

	return dfs(root.Left, root.Right)
}
