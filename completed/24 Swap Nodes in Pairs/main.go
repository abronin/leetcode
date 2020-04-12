package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// var ln1 *ListNode = nil
	ln := ListNode{Val: 1, Next: nil}
	ln.PushValue(2)
	ln.PushValue(3)
	ln.PushValue(4)
	ln.PushValue(5)
	ln.PushValue(6)
	fmt.Println(swapPairs(&ln))
}

func swapPairs(head *ListNode) *ListNode {
	newHead := head

	current := head
	var next *ListNode = nil
	if current != nil {
		next = current.Next
	}

	if current != nil && next != nil {
		newHead = swapPair(current)
		newHead.Next.Next = swapPairs(newHead.Next.Next)
	}

	return newHead
}

func swapPair(head *ListNode) *ListNode {
	res := head.Next
	tail := head.Next.Next
	head.Next.Next = head
	head.Next = tail
	return res
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
