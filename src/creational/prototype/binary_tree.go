package prototype

import (
	"fmt"
)

type BinaryTree struct {
	value int
	left  TreeNode
	right TreeNode
}

func NewBinaryTree(val int) *BinaryTree {
	return &BinaryTree{value: val}
}

func (bt *BinaryTree) Value() int {
	return bt.value
}

func (bt *BinaryTree) Left() TreeNode {
	return bt.left
}

func (bt *BinaryTree) SetLeft(node TreeNode) {
	bt.left = node
}

func (bt *BinaryTree) Right() TreeNode {
	return bt.right
}

func (bt *BinaryTree) SetRight(node TreeNode) {
	bt.right = node
}

func (bt *BinaryTree) Print() string {
	var s string
	stack := make([]TreeNode, 0)
	var curr TreeNode = bt
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left()
		}

		l := len(stack)
		curr = stack[l-1]
		stack = stack[:l-1]

		s += fmt.Sprintf("%d ", curr.Value())

		curr = curr.Right()
	}

	return s
}

func (bt *BinaryTree) Clone() TreeNode {
	cl := NewBinaryTree(bt.value)
	if bt.left != nil {
		cl.left = bt.left.Clone().(*BinaryTree)
	}
	if bt.right != nil {
		cl.right = bt.right.Clone().(*BinaryTree)
	}
	return cl
}
