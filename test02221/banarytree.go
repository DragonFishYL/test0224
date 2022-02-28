package main

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

type Tree struct {
	root *Node
}

func (n *Node) SetValue(value int) Node {
	n.Value = value
	return *n
}

func (tree *Tree) CreateTree(node *Node, newNode Node) {
	if node.Value > newNode.Value {
		if node.Left == nil {
			node.Left = &newNode
		} else {
			tree.CreateTree(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = &newNode
		} else {
			tree.CreateTree(node.Right, newNode)
		}
	}
}

func (tree *Tree) InGenerate(node Node) {
	if node.Left != nil {
		tree.InGenerate(*node.Left)
	}
	fmt.Printf("%+v\n", node.Value)
	if node.Right != nil {
		tree.InGenerate(*node.Right)
	}
}
