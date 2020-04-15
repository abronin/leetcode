package main

import "fmt"

func main() {
	root := &ListNode{Val: 1, Next: nil}
	node := AddNext(root, 2)
	node = AddNext(node, 3)
	node = AddNext(node, 4)
	node = AddNext(node, 5)
	node = AddNext(node, 6)
	k := 1

	fmt.Println(splitListToParts(root, k))
}

// ListNode define a node of the singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	length := ListLen(root)
	quotient, remainder := length/k, length%k
	result := make([]*ListNode, k)
	for i := 0; i < k; i++ {
		step := quotient
		if remainder > 0 {
			step = quotient + 1
			remainder--
		}

		result[i] = root
		for j := 0; j < step-1; j++ {
			root = root.Next
		}
		if root != nil {
			root.Next, root = nil, root.Next
		}
	}
	return result
}

func AddNext(node *ListNode, value int) *ListNode {
	next := &ListNode{Val: value, Next: nil}
	node.Next = next
	return next
}

func ListLen(root *ListNode) int {
	length := 0
	for node := root; node != nil; node = node.Next {
		length++
	}
	return length
}

func (head *ListNode) String() string {
	if head == nil {
		return "nil"
	}

	res := ""

	current := head
	for current != nil {
		if current.Next == nil {
			res = res + fmt.Sprintf("%d", current.Val)
		} else {
			res = res + fmt.Sprintf("%d -> ", current.Val)
		}
		current = current.Next
	}

	return res
}
