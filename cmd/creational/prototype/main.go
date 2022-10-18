package main

import (
	"fmt"
	pt "gof/src/creational/prototype"
)

func main() {
	bt := pt.NewBinaryTree(10)
	l := pt.NewBinaryTree(9)
	r := pt.NewBinaryTree(8)
	l.SetLeft(pt.NewBinaryTree(7))
	l.SetRight(pt.NewBinaryTree(6))
	r.SetLeft(pt.NewBinaryTree(5))
	r.SetRight(pt.NewBinaryTree(4))
	bt.SetLeft(l)
	bt.SetRight(r)

	fmt.Println(fmt.Sprintf("Original: %s", bt.Print()))
	fmt.Println(fmt.Sprintf("Clone: %s", bt.Clone().Print()))
}
