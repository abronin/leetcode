package main

import (
	"testing"
)

func TestLevel(t *testing.T) {
	input := "level"
	expected := "l"
	given := longestPrefix(input)

	if given != expected {
		t.Errorf("%s given but %s expected for %s", given, expected, input)
	}
}

func TestAbabab(t *testing.T) {
	input := "ababab"
	expected := "abab"
	given := longestPrefix(input)

	if given != expected {
		t.Errorf("%s given but %s expected for %s", given, expected, input)
	}
}

func TestLeet(t *testing.T) {
	input := "leetcodeleet"
	expected := "leet"
	given := longestPrefix(input)

	if given != expected {
		t.Errorf("%s given but %s expected for %s", given, expected, input)
	}
}

func TestA(t *testing.T) {
	input := "a"
	expected := ""
	given := longestPrefix(input)

	if given != expected {
		t.Errorf("%s given but %s expected for %s", given, expected, input)
	}
}

func TestAaaaa(t *testing.T) {
	input := "aaaaa"
	expected := "aaaa"
	given := longestPrefix(input)

	if given != expected {
		t.Errorf("%s given but %s expected for %s", given, expected, input)
	}
}
