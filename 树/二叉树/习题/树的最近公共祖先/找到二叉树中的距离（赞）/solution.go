package leetcode_1740

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDistance(root *TreeNode, p int, q int) int {
	// 边界条件
	if root == nil || p == q {
		return 0
	}

	ancestor := lca(root, p, q)

	a, b, c := depth(root, ancestor), depth(root, p), depth(root, q)
	d := b + c - 2*a

	return d
}

// lca 获取公共祖先
func lca(root *TreeNode, p, q int) int {
	if root == nil {
		return -1
	}
	if root.Val == p || root.Val == q {
		return root.Val
	}
	l := lca(root.Left, p, q)
	r := lca(root.Right, p, q)
	if l != -1 && r != -1 {
		return root.Val
	}
	if l != -1 {
		return l
	}
	return r
}

// depth 计算节点深度
func depth(root *TreeNode, target int) int {
	if root == nil {
		return math.MinInt
	}
	if root.Val == target {
		return 0
	}
	return max(depth(root.Left, target), depth(root.Right, target)) + 1
}
