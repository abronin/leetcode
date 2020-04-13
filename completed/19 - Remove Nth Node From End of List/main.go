package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ln := ListNode{Val: 1, Next: nil}
	// ln.PushValue(2)
	// ln.PushValue(3)
	// ln.PushValue(4)
	fmt.Println(removeNthFromEnd(&ln, 1))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	prev, target, current := head, head, head
	var accumulated int

	for accumulated = 0; current != nil; accumulated++ {
		if accumulated >= n {
			target = target.Next
		}

		if accumulated >= n+1 {
			prev = prev.Next
		}

		current = current.Next
	}

	prev.Next = target.Next

	if accumulated == n {
		return head.Next
	}

	return head
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

func (head *ListNode) Push(newcomer *ListNode) {
	current := head
	last := head
	for current != nil {
		last = current
		current = current.Next
	}
	last.Next = newcomer
}

func (head *ListNode) PushValue(value int) {
	newcomer := ListNode{Val: value, Next: nil}
	head.Push(&newcomer)
}
