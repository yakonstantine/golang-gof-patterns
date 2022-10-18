package prototype

type TreeNode interface {
	Value() int
	Left() TreeNode
	Right() TreeNode
	SetLeft(node TreeNode)
	SetRight(node TreeNode)
	Print() string
	Clone() TreeNode
}
