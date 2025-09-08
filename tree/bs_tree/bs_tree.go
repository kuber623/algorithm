package bs_tree

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (t *Tree) Find(value int) *TreeNode {
	return t.findNode(t.Root, value)
}

func (t *Tree) findNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}

	switch {
	case value < node.Value:
		return t.findNode(node.Left, value)
	case value > node.Value:
		return t.findNode(node.Right, value)
	default:
		return node
	}
}

func (t *Tree) Insert(value int) {
	t.Root = t.insertNode(t.Root, value)
}

func (t *Tree) insertNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}

	switch {
	case value < node.Value:
		node.Left = t.insertNode(node.Left, value)
	case value > node.Value:
		node.Right = t.insertNode(node.Right, value)
	}

	return node
}

func (t *Tree) Delete(value int) {
	t.deleteNode(t.Root, value)
}

func (t *Tree) deleteNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}

	switch {
	case value < node.Value:
		node.Left = t.deleteNode(node.Left, value)
	case value > node.Value:
		node.Right = t.deleteNode(node.Right, value)
	default:
		// 找到目标节点

		// 情况一：目标节点为叶子节点
		// 解法一：直接删除目标节点
		if node.Left == nil && node.Right == nil {
			return nil
		}

		// 情况二：目标节点只有一个子节点
		// 解法二：提拔该子节点提拔
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// 情况三：目标节点有两个子节点
		// 解法三：找到当前节点的后继节点（即右子树中最小的节点），交换当前节点与后继节点的值后，删除后继节点
		next := node.Right
		for next.Left != nil {
			next = next.Left
		}
		node.Value = next.Value
		node.Right = t.deleteNode(node.Right, next.Value)
	}

	return node
}
