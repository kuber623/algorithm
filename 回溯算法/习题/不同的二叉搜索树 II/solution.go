package leetcode_95

// https://leetcode.cn/problems/unique-binary-search-trees-ii
// 难度：中等
// 题解：回溯算法

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return backtrack(0, n-1)
}

func backtrack(start, end int) []*TreeNode {
	if start == end {
		return []*TreeNode{{Val: start + 1}}
	}

	trees := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		switch i {
		case start:
			rightTrees := backtrack(start+1, end)
			for _, right := range rightTrees {
				trees = append(trees, &TreeNode{Val: i + 1, Right: right})
			}
		case end:
			leftTrees := backtrack(start, end-1)
			for _, left := range leftTrees {
				trees = append(trees, &TreeNode{Val: i + 1, Left: left})
			}
		default:
			leftTrees, rightTrees := backtrack(start, i-1), backtrack(i+1, end)
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					trees = append(trees, &TreeNode{Val: i + 1, Left: left, Right: right})
				}
			}

		}
	}
	return trees
}
