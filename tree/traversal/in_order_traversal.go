package traversal

// InOrderTraversal 二叉树中序遍历（非递归实现）
// 中序遍历顺序为 左 -> 根 -> 右
func InOrderTraversal(tree *Tree) []int {
	ret := make([]int, 0)

	if tree == nil || tree.Root == nil {
		return ret
	}

	cur, stack := tree.Root, make([]*TreeNode, 0)
	for cur != nil || len(stack) > 0 {
		// 持续将左侧节点压栈
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}

		// 将最左节点出栈
		cur, stack = stack[len(stack)-1], stack[:len(stack)-1]

		// 访问当前节点
		ret = append(ret, cur.Value)

		// 转向当前节点右子树
		cur = cur.Right
	}

	return ret
}
