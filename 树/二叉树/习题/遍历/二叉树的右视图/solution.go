package leetcode_199

// https://leetcode.cn/problems/binary-tree-right-side-view/
// 难度：中等
// 题解：
// 先递归右子树，再递归左子树，当某个深度首次到达时，对应的节点就在右视图中

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	views := make([]int, 0)

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(views) {
			views = append(views, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)

	return views
}
