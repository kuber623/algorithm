package leetcode_236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	// 从左子树寻找公共祖先
	lca1 := lowestCommonAncestor(root.Left, p, q)
	// 从右子树寻找公共祖先
	lca2 := lowestCommonAncestor(root.Right, p, q)
	if lca1 != nil && lca2 != nil {
		return root
	}
	if lca1 != nil {
		return lca1
	}
	return lca2
}
