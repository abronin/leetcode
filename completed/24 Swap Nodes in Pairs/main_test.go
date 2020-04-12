package main

import (
	"testing"
)

func TestNil(t *testing.T) {
	var ln1 *ListNode = nil

	res := swapPairs(ln1)

	if res != nil {
		t.Errorf("When you swap a nil list you should get a nil")
	}
}

func TestOne(t *testing.T) {
	ln := ListNode{Val: 1, Next: nil}

	res := swapPairs(&ln)

	if res.Val != 1 || res.Next != nil {
		t.Errorf("When you swap a one-length list you should get it back")
	}
}

func TestEven(t *testing.T) {
	ln := ListNode{Val: 1, Next: nil}
	ln.PushValue(2)
	ln.PushValue(3)
	ln.PushValue(4)

	res := swapPairs(&ln)

	if res.Next.Next.Val != 4 || res.Next.Next.Next.Val != 3 {
		t.Errorf("1-2-3-4 should become 2-1-4-3")
	}
}

func TestOdd(t *testing.T) {
	ln := ListNode{Val: 1, Next: nil}
	ln.PushValue(2)
	ln.PushValue(3)

	res := swapPairs(&ln)

	if res.Val != 2 || res.Next.Val != 1 || res.Next.Next.Val != 3 || res.Next.Next.Next != nil {
		t.Errorf("1-2-3 should become 2-1-3")
	}
}
