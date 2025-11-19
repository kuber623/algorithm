package leetcode_104

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 难度：简单

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
