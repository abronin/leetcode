package main

import "fmt"

func main() {
	root := &ListNode{Val: 1, Next: nil}
	node := AddNext(root, 2)
	node = AddNext(node, 3)
	node = AddNext(node, 4)
	node = AddNext(node, 5)
	node = AddNext(node, 6)
	fmt.Println(oddEvenList(root))
}

// ListNode define a node of the singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	oddEnd := head
	evenEnd := head.Next
	evenStart := head.Next

	for oddEnd != nil && evenEnd != nil {
		// fmt.Println(head)
		oddEnd, evenEnd = step(oddEnd, evenEnd, evenStart)
	}

	return head
}

func step(odd, even, evenStart *ListNode) (newOdd, newEven *ListNode) {
	if even.Next == nil {
		return nil, nil
	}
	odd.Next = even.Next
	if odd.Next != nil {
		even.Next = odd.Next.Next
		odd.Next.Next = evenStart
	}
	return odd.Next, even.Next
}

func (head *ListNode) String() string {
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

func AddNext(node *ListNode, value int) *ListNode {
	next := &ListNode{Val: value, Next: nil}
	node.Next = next
	return next
}

// 1
// 2
// 1 -> 2 -> 3 -> 4 -> 5 -> nil

// next of 1 should be 3
// next of 2 should be 4
// next of 3 should be 2

// 3
// 4
// 1    3    2    4    5    nil

// next of 3 should be 5
// next of 4 should be nil
// next of 5 should be 4

// 1    3    5    2    4    nil

// 5
// nil
