package main

import "fmt"

// Chindex store a rune and it's index from the string
type Chindex struct {
	Val   rune
	Index int
}

// ListNode of the List
type ListNode struct {
	Val  Chindex
	Prev *ListNode
	Next *ListNode
}

// List having Head and Tail
type List struct {
	Head *ListNode
	Tail *ListNode
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("", "ACF"))
	fmt.Println(minWindow("AA", "AA"))
}

func minWindow(s string, t string) string {
	sr, tr := []rune(s), []rune(t)
	result := ""
	reference, counter := make(map[rune]int), make(map[rune]int)
	list := NewList()

	for _, r := range tr {
		reference[r]++
	}

	wLength := len(s) + 1
	for si, r := range sr {
		if _, ok := reference[r]; ok {
			list.Push(NewChindex(r, si))
			counter[r]++
			for isFull(reference, counter) {
				if wLength > list.Tail.Val.Index-list.Head.Val.Index {
					wLength = list.Tail.Val.Index - list.Head.Val.Index + 1
					result = s[list.Head.Val.Index : list.Tail.Val.Index+1]
				}
				first := list.Shift()
				counter[first.Val.Val]--
			}
		}
	}

	return result
}

func isFull(reference, counter map[rune]int) bool {
	for key, refCount := range reference {
		if count, ok := counter[key]; !ok || count < refCount {
			return false
		}
	}
	return true
}

// NewList create a new dlist
func NewList() *List {
	return &List{Head: nil, Tail: nil}
}

// NotEmpty check for emptiness of the list
func (list *List) NotEmpty() bool {
	return list.Head != nil
}

// Clear the list
func (list *List) Clear() {
	list.Head = nil
	list.Tail = nil
}

// Unshift a node from the list's Head
func (list *List) Unshift(r Chindex) {
	node := newListNode(r, nil, list.Head)

	if list.NotEmpty() {
		list.Head.Prev = node
	} else {
		list.Tail = node
	}
	list.Head = node
}

// Shift a node from the list's Head
func (list *List) Shift() *ListNode {
	node := list.Head
	if list.NotEmpty() {
		if list.Head.Next == nil {
			list.Clear()
		} else {
			list.Head = list.Head.Next
		}
	}
	return node
}

// Push a node to the list's Tail
func (list *List) Push(r Chindex) {
	node := newListNode(r, list.Tail, nil)

	if list.NotEmpty() {
		list.Tail.Next = node
	} else {
		list.Head = node
	}

	list.Tail = node
}

// Pull a node from the list's Tail
func (list *List) Pull() *ListNode {
	node := list.Tail
	if list.NotEmpty() {
		if list.Tail.Prev == nil {
			list.Clear()
		} else {
			list.Tail = list.Tail.Prev
		}
	}
	return node
}

func newListNode(r Chindex, prev, next *ListNode) *ListNode {
	node := &ListNode{Val: r, Prev: prev, Next: next}
	return node
}

// NewChindex creates new Chindex struct
func NewChindex(value rune, index int) Chindex {
	return Chindex{Val: value, Index: index}
}

func (r Chindex) String() string {
	return fmt.Sprintf("%c/%d", r.Val, r.Index)
}

func (node *ListNode) String() string {
	return fmt.Sprintf("%s", node.Val)
}

func (list *List) String() string {
	if !list.NotEmpty() {
		return "List is empty!"
	}

	res := ""
	for current := list.Head; current != nil; current = current.Next {
		res += fmt.Sprintf("%s ", current)
	}
	return res
}
