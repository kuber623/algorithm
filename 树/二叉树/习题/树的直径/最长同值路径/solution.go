package leetcode_687

// https://leetcode.cn/problems/longest-univalue-path/
// 难度：中等

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	maxPath := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left, right := node.Left, node.Right
		d1 := dfs(left)
		if left == nil || left.Val != node.Val {
			d1 = 0
		}
		d2 := dfs(right)
		if right == nil || right.Val != node.Val {
			d2 = 0
		}
		maxPath = max(maxPath, d1+d2)

		return max(d1, d2) + 1
	}
	dfs(root)

	return maxPath
}
