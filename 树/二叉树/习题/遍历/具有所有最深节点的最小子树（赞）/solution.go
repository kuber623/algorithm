package leetcode_865

// https://leetcode.cn/problems/smallest-subtree-with-all-the-deepest-nodes/
// 难度：中等

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	maxDepth := -1                           // 记录最大深度
	deepest := make([]*TreeNode, 0)          // 记录位于最大深度的节点
	parents := make(map[*TreeNode]*TreeNode) // 记录每个节点的父节点的值映射

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		left, right := node.Left, node.Right

		// 因为最深节点只能是树的叶子节点，所以只需要关注叶子节点的深度
		if depth >= maxDepth && left == nil && right == nil {
			if depth > maxDepth {
				deepest = []*TreeNode{node}
				maxDepth = depth
			} else {
				deepest = append(deepest, node)
			}
		}

		parents[left], parents[right] = node, node
		dfs(left, depth+1)
		dfs(right, depth+1)
	}

	dfs(root, 0)

	// 如果只有一个最大深度节点，直接返回该节点即可
	if len(deepest) == 1 {
		return deepest[0]
	}

	// 如果存在多个最大深度节点，则返回其最大公共祖先
	return lca(deepest, parents)
}

func lca(nodes []*TreeNode, parents map[*TreeNode]*TreeNode) *TreeNode {
	var x *TreeNode
	same := true
	for i := 0; i < len(nodes); i++ {
		if i == 0 {
			x = nodes[i]
		} else {
			if x != nodes[i] {
				same = false
				break
			}
		}
	}
	if same {
		return x
	}

	for i := 0; i < len(nodes); i++ {
		nodes[i] = parents[nodes[i]]
	}
	return lca(nodes, parents)
}
