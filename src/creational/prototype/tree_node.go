package prototype

type TreeNode interface {
	Value() int
	Print() string
	Clone() TreeNode
	SetLeft(node TreeNode)
	SetRight(node TreeNode)
}
