package leetcode_951

// https://leetcode.cn/problems/flip-equivalent-binary-trees/
// 难度：中等

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	// 比对当前节点是否相等
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	// 分别在未翻转和翻转的情况下比对子节点是否相等
	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}
