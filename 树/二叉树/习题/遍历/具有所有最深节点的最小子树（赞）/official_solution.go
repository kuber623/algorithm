package leetcode_865

func subtreeWithAllDeepestOfficial(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) (int, *TreeNode)
	dfs = func(node *TreeNode) (int, *TreeNode) {
		if node == nil {
			return 0, nil
		}

		d1, lca1 := dfs(node.Left)
		d2, lca2 := dfs(node.Right)

		switch {
		case d1 > d2:
			return d1 + 1, lca1
		case d1 < d2:
			return d2 + 1, lca2
		default:
			return d1 + 1, node
		}
	}

	_, lca := dfs(root)
	return lca
}
