package leetcode_110

import "math"

// https://leetcode.cn/problems/balanced-binary-tree/
// 难度：简单

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(*TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}

		// 判断左子树是否为平衡二叉树
		h1, ok := dfs(node.Left)
		if !ok {
			return 0, false
		}
		// 判断右子树是否为平衡二叉树
		h2, ok := dfs(node.Right)
		if !ok {
			return 0, false
		}
		// 比对左右子树的高度差是否超过 1
		ok = math.Abs(float64(h2-h1)) <= 1
		if !ok {
			return 0, false
		}

		return max(h1, h2) + 1, true
	}

	_, ok := dfs(root)
	return ok
}
