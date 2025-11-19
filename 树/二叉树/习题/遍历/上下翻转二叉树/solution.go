package leetcode_156

// https://leetcode.cn/problems/binary-tree-upside-down/
// 难度：中等

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	ret := upsideDownBinaryTree(root.Left)
	left, right := root.Left, root.Right
	left.Left, left.Right = right, root
	root.Left, root.Right = nil, nil
	return ret
}
