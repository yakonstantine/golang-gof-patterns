package prototype

import (
	"fmt"
)

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func NewBinaryTree(val int) *BinaryTree {
	return &BinaryTree{
		value: val,
	}
}

func (bt *BinaryTree) Value() int {
	return bt.value
}

func (bt *BinaryTree) Left() TreeNode {
	if bt == nil || bt.left == nil {
		return nil
	}
	return bt.left
}

func (bt *BinaryTree) SetLeft(node TreeNode) {
	if node == nil {
		return
	}
	bt.left = node.(*BinaryTree)
}

func (bt *BinaryTree) Right() TreeNode {
	if bt == nil || bt.right == nil {
		return nil
	}
	return bt.right
}

func (bt *BinaryTree) SetRight(node TreeNode) {
	if node == nil {
		return
	}
	bt.right = node.(*BinaryTree)
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
