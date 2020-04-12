package main

import (
	"testing"
)

func TestEmptyInput(t *testing.T) {
	n := []int{}
	if canJump(n) {
		t.Errorf("Should not be able to jump through the empty line")
	}
}

func TestZeroStart(t *testing.T) {
	n := []int{0, 1, 1}
	if canJump(n) {
		t.Errorf("Should not be able to jump through the line starts with zero")
	}
}

func TestZeroLine(t *testing.T) {
	n := []int{0}
	if !canJump(n) {
		t.Errorf("Should be able to jump through the line contains of zero")
	}
}

func TestHugeJump(t *testing.T) {
	n := []int{1, 1000, 0, 0, 0}
	if !canJump(n) {
		t.Errorf("Should be able to jump through the 1, 1000, 0, 0, 0")
	}
}

func TestCanJump(t *testing.T) {
	n := []int{2, 3, 1, 1, 4}
	if !canJump(n) {
		t.Errorf("Should be able to jump through the 2, 3, 1, 1, 4")
	}
}

func TestCantJump(t *testing.T) {
	n := []int{3, 2, 1, 0, 4}
	if canJump(n) {
		t.Errorf("Should not ba able to jump through the 2, 3, 1, 1, 4")
	}
}
