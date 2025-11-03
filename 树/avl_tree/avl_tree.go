package avl_tree

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
	Height      int
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
		return &TreeNode{Value: value, Height: 1}
	}

	switch {
	case value < node.Value:
		node.Left = t.insertNode(node.Left, value)
	case value > node.Value:
		node.Right = t.insertNode(node.Right, value)
	default:
		return node
	}

	t.updateHeight(node)

	bf := t.balanceFactor(node)

	// 情况一：左左失衡
	// 解法一：右旋
	if bf > 1 && value < node.Left.Value {
		return t.rightRotate(node)
	}

	// 情况二：右右失衡
	// 解法二：左旋
	if bf < -1 && value > node.Right.Value {
		return t.leftRotate(node)
	}

	// 情况三：左右失衡
	// 解法三：先右旋左子树，再按左左失衡处理
	if bf > 1 && value > node.Left.Value {
		node.Left = t.leftRotate(node.Left)
		return t.rightRotate(node)
	}

	// 情况四：右左失衡
	// 解法三：先左旋右子树，再按右右失衡处理
	if bf < -1 && value < node.Right.Value {
		node.Right = t.rightRotate(node.Right)
		return t.leftRotate(node)
	}

	return node
}

func (t *Tree) Delete(value int) {
	t.Root = t.deleteNode(t.Root, value)
}

func (t *Tree) deleteNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}

	// 执行二叉搜索树节点删除
	switch {
	case value < node.Value:
		node.Left = t.deleteNode(node.Left, value)
	case value > node.Value:
		node.Right = t.deleteNode(node.Right, value)
	default:
		if node.Left == nil && node.Right == nil {
			return nil
		}

		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		next := node.Right
		for next.Left != nil {
			next = next.Left
		}
		node.Value = next.Value
		node.Right = t.deleteNode(node.Right, next.Value)
	}

	t.updateHeight(node)

	bf := t.balanceFactor(node)

	// 情况一：左左失衡
	// 解法一：右旋
	if bf > 1 && value > node.Right.Value {
		return t.rightRotate(node)
	}

	// 情况二：右右失衡
	// 解法二：左旋
	if bf < -1 && value < node.Left.Value {
		return t.leftRotate(node)
	}

	// 情况三：左右失衡
	// 解法三：先右旋左子树，再按左左失衡处理
	if bf > 1 && value < node.Right.Value {
		node.Left = t.leftRotate(node.Left)
		return t.rightRotate(node)
	}

	// 情况四：右左失衡
	// 解法四：先左旋右子树，再按右右失衡处理
	if bf < -1 && value > node.Left.Value {
		node.Right = t.rightRotate(node.Right)
		return t.leftRotate(node)
	}

	return node
}

// rightRotate 右旋
func (t *Tree) rightRotate(node *TreeNode) *TreeNode {
	left := node.Left
	left.Right, node.Left = node, left.Right
	t.updateHeight(node)
	t.updateHeight(left)
	return left
}

// leftRotate 左旋
func (t *Tree) leftRotate(node *TreeNode) *TreeNode {
	right := node.Right
	right.Left, node.Right = node, right.Left
	t.updateHeight(node)
	t.updateHeight(right)
	return right
}

func (t *Tree) balanceFactor(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return t.height(node.Left) - t.height(node.Right)
}

func (t *Tree) height(node *TreeNode) int {
	if node == nil {
		return -1
	}
	return node.Height
}

func (t *Tree) updateHeight(node *TreeNode) {
	if node == nil {
		return
	}

	node.Height = 1 + max(t.height(node.Left), t.height(node.Right))
}
