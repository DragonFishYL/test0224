package main

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func (n *Node) SetValue(value int) *Node {
	n.Value = value
	return n
}

var (
	root      Node
	sortSlice = make([]int, 0)
)

func (n *Node) InsertNode(node *Node, newNode Node) {
	if node.Value > newNode.Value {
		fmt.Println("left")
		if node.Left == nil {
			fmt.Println("-------------------left")
			node.Left = node
		} else {
			fmt.Println("left-------------------")
			n.InsertNode(node.Left, newNode)
		}
	} else {
		fmt.Println("right")
		if node.Right == nil {
			fmt.Println("-------------------right")
			node.Right = node
		} else {
			fmt.Println("right-------------------")
			n.InsertNode(node.Right, newNode)
		}
	}
}

func (n *Node) GenerNode() {
	var nilNode Node
	for _, v := range sortSlice {
		node := n.SetValue(v)
		if root == nilNode {
			root = *node
		} else {
			n.InsertNode(&root, *node)
		}
	}
}

func (n *Node) InGenera() {
	var nilNode Node
	if root != nilNode {
		n.InItera(root)
	}
}

func (n *Node) InItera(node Node) {
	if node.Left != nil {
		n.InItera(*node.Left)
	}
	fmt.Println(node.Value)
	if node.Right != nil {
		n.InItera(*node.Right)
	}
}
