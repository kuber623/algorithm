package leetcode_255

import "math"

// https://leetcode.cn/problems/verify-preorder-sequence-in-binary-search-tree
// 难度：中等
// 题解：
// 先序遍历的顺序是：中 -> 左 -> 右，由于遍历的是二叉搜索树，因此左子树所有元素的值均小于父节点，右子树所有元素的值均大于父节点
// 因此我们可以使用单调栈计算当前元素下一个更大元素的位置，得到当前节点右孩子的位置，如果右子树所有值均大于父节点，则为二叉搜索树
// 我们可以基于单调栈使用递归的方式对每个子树进行先序遍历验证

var order, nextGreater []int

func verifyPreorder(preorder []int) bool {
	order = preorder

	// 使用单调栈获取当前元素下一个更大元素的索引
	nextGreater = make([]int, len(preorder))
	monostack := make([]int, 0, len(preorder))
	for i := len(preorder) - 1; i >= 0; i-- {
		for len(monostack) > 0 && preorder[i] > preorder[monostack[len(monostack)-1]] {
			monostack = monostack[:len(monostack)-1]
		}

		if len(monostack) > 0 {
			nextGreater[i] = monostack[len(monostack)-1]
		} else {
			nextGreater[i] = -1
		}

		monostack = append(monostack, i)
	}

	return recursionVerify(0, len(preorder)-1, math.MinInt, math.MaxInt)
}

func recursionVerify(left, right int, down, up int) bool {
	if left > right {
		return true
	}

	for i := left; i <= right; i++ {
		if order[i] < down || order[i] > up {
			return false
		}
	}

	node := left
	if nextGreater[node] == -1 {
		return recursionVerify(left+1, right, down, order[node])
	}
	return recursionVerify(left+1, nextGreater[node]-1, down, order[node]) &&
		recursionVerify(nextGreater[node], right, order[node], up)
}
