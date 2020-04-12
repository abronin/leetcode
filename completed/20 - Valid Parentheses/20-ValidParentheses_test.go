package main

import (
	"testing"
)

func TestEmptyIsValid(t *testing.T) {
	if !isValid("") {
		t.Errorf("Empty string should be a valid parentheses")
	}
}

func TestSimpleValid(t *testing.T) {
	if !isValid("()[]{}") {
		t.Errorf("()[]{} should be a valid parentheses")
	}
}

func TestNestedValid(t *testing.T) {
	if !isValid("([{}]){}") {
		t.Errorf("([{}]){} should be a valid parentheses")
	}
}

func TestClosingFirstIsInvalid(t *testing.T) {
	if isValid("]") {
		t.Errorf("Valid parentheses can not start from opener")
	}
}

func TestIncompleteInvalid(t *testing.T) {
	if isValid("([]{}") {
		t.Errorf("([]{} should be an invalid parentheses")
	}
}

func TestOrderInvalid(t *testing.T) {
	if isValid("([{})]{}") {
		t.Errorf("([{})]{} should be an invalid parentheses")
	}
}
