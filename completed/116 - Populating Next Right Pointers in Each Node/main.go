package main

import "fmt"

func main() {
	root := NewNode(1, nil, nil)
	// root.Left = NewNode(21, nil, nil)
	// root.Right = NewNode(22, nil, nil)
	// root.Left.Left = NewNode(31, nil, nil)
	// root.Left.Right = NewNode(32, nil, nil)
	// root.Right.Left = NewNode(33, nil, nil)
	// root.Right.Right = NewNode(34, nil, nil)

	// root.Left.Left.Left = NewNode(41, nil, nil)
	// root.Left.Left.Right = NewNode(42, nil, nil)
	// root.Left.Right.Left = NewNode(43, nil, nil)
	// root.Left.Right.Right = NewNode(44, nil, nil)
	// root.Right.Left.Left = NewNode(45, nil, nil)
	// root.Right.Left.Right = NewNode(46, nil, nil)
	// root.Right.Right.Left = NewNode(47, nil, nil)
	// root.Right.Right.Right = NewNode(48, nil, nil)

	root = connect(root)
	fmt.Println(root)
}

// Node definition
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	if !IsLeaf(root) {
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
		connect(root.Left)
		connect(root.Right)
	}
	return root
}

// NewNode create new node
func NewNode(value int, left, right *Node) *Node {
	return &Node{Val: value, Left: left, Right: right, Next: nil}
}

// IsLeaf check for the node
func IsLeaf(node *Node) bool {
	return node.Left == nil
}

func (node *Node) String() string {
	res := fmt.Sprintf("%d ", node.Val)
	if !IsLeaf(node) {
		res += fmt.Sprintf("%s ", node.Left)
		res += fmt.Sprintf("%s ", node.Right)
	}
	return res
}
