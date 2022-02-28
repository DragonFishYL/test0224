package main

import (
	"testing"
)

func TestNode(t *testing.T) {
	sortSlice := []int{55, 3, 10, 12, 66, 7, 11, 33, 26}
	var (
		tree Tree
		node Node
	)
	for _, v := range sortSlice {
		n := node.SetValue(v)
		if tree.root == nil {
			tree.root = &n
		} else {
			tree.CreateTree(tree.root, n)
		}
	}
	tree.InGenerate(*tree.root)
}
