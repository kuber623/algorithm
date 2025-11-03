package leetcode_333

// https://leetcode.cn/problems/largest-bst-subtree
// 难度：中等
// 题解：深度优先遍历 + 关键信息复用

var maxSize int

func largestBSTSubtree(root *TreeNode) int {
	maxSize = 0
	dfs(root)
	return maxSize
}

func dfs(root *TreeNode) (bool, int, int, int) {
	if root == nil {
		return true, 0, 0, 0
	}

	// 遍历左子树
	lok, lsize, lmin, lmax := dfs(root.Left)
	// 遍历右子树
	rok, rsize, rmin, rmax := dfs(root.Right)
	// 如果左右子树不为二叉搜索树，则该树不为二叉搜索树
	if !lok || !rok {
		return false, 0, 0, 0
	}
	// 如果左子树最大值比树根大，或右子树最小值比树根小，则该树不为二叉搜索树
	if (lsize > 0 && lmax >= root.Val) || (rsize > 0 && rmin <= root.Val) {
		return false, 0, 0, 0
	}

	size := lsize + rsize + 1
	minval, maxval := lmin, rmax
	// 如果不存在左子树，该树的最小值为树根
	if lsize == 0 {
		minval = root.Val
	}
	// 如果不存在右子树，该树的最大值为树根
	if rsize == 0 {
		maxval = root.Val
	}

	// 更新最大值
	maxSize = max(maxSize, size)

	return true, size, minval, maxval
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
