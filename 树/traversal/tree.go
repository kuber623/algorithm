package traversal

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}
