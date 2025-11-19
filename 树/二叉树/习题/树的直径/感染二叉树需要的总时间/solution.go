package leetcode_2385

import "math"

// https://leetcode.cn/problems/amount-of-time-for-binary-tree-to-be-infected/
// 难度：中等
// 题解：感染时间 = MAX{}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	time := math.MinInt

	var dfs func(node *TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, false
		}

		if node.Val == start {
			time = max(time, maxDepth(node)-1)
			return 1, true
		}

		d1, ok1 := dfs(node.Left)
		d2, ok2 := dfs(node.Right)
		if !ok1 && !ok2 {
			return max(d1, d2) + 1, false
		}

		time = max(time, d1+d2)

		if ok1 {
			return d1 + 1, true
		}
		return d2 + 1, true
	}
	dfs(root)

	return time
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
