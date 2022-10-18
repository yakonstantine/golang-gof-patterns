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
	return &BinaryTree{
		value: val,
	}
}

func (bt *BinaryTree) Value() int {
	return bt.value
}

func (bt *BinaryTree) Print() string {
	var s string
	stack := make([]*BinaryTree, 0)
	curr := bt
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			if curr.left == nil {
				curr = nil
				continue
			}
			curr = curr.left.(*BinaryTree)
		}

		l := len(stack)
		curr = stack[l-1]
		stack = stack[:l-1]

		s += fmt.Sprintf("%d ", curr.value)

		if curr.right == nil {
			curr = nil
			continue
		}
		curr = curr.right.(*BinaryTree)
	}

	return s
}

func (bt *BinaryTree) Clone() TreeNode {
	cl := NewBinaryTree(bt.value)
	if bt.left != nil {
		cl.left = bt.left.Clone()
	}
	if bt.right != nil {
		cl.right = bt.right.Clone()
	}
	return cl
}

func (bt *BinaryTree) SetLeft(node TreeNode) {
	bt.left = node
}

func (bt *BinaryTree) SetRight(node TreeNode) {
	bt.right = node
}
