package main

import (
	"testing"
)

func TestOne(t *testing.T) {
	input := []int{1}
	expected := 2
	res := firstMissingPositive(input)
	if res != expected {
		t.Errorf("%d given but %d is expected for %v", res, expected, input)
	}
}

func TestTwo(t *testing.T) {
	input := []int{2}
	expected := 1
	res := firstMissingPositive(input)
	if res != expected {
		t.Errorf("%d given but %d is expected for %v", res, expected, input)
	}
}

func TestEmpty(t *testing.T) {
	input := []int{}
	expected := 1
	res := firstMissingPositive(input)
	if res != expected {
		t.Errorf("%d given but %d is expected for %v", res, expected, input)
	}
}

func Test1(t *testing.T) {
	input := []int{1, 2, 0}
	expected := 3
	res := firstMissingPositive(input)
	if res != expected {
		t.Errorf("%d given but %d is expected for %v", res, expected, input)
	}
}

func Test2(t *testing.T) {
	input := []int{3, 4, -1, 1}
	expected := 2
	res := firstMissingPositive(input)
	if res != expected {
		t.Errorf("%d given but %d is expected for %v", res, expected, input)
	}
}

func Test3(t *testing.T) {
	input := []int{7, 8, 9, 11, 12}
	expected := 1
	res := firstMissingPositive(input)
	if res != expected {
		t.Errorf("%d given but %d is expected for %v", res, expected, input)
	}
}
