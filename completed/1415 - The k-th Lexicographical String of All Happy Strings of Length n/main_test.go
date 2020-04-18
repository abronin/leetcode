package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	n := 1
	k := 3

	expected := "c"
	given := getHappyString(n, k)

	if given != expected {
		t.Errorf("%s given but %s expected for %d %d", given, expected, n, k)
	}
}

func TestSecond(t *testing.T) {
	n := 1
	k := 4

	expected := ""
	given := getHappyString(n, k)

	if given != expected {
		t.Errorf("%s given but %s expected for %d %d", given, expected, n, k)
	}
}

func TestThird(t *testing.T) {
	n := 4
	k := 4

	expected := "abcb"
	given := getHappyString(n, k)

	if given != expected {
		t.Errorf("%s given but %s expected for %d %d", given, expected, n, k)
	}
}

func TestFifth(t *testing.T) {
	n := 10
	k := 100

	expected := "abacbabacb"
	given := getHappyString(n, k)

	if given != expected {
		t.Errorf("%s given but %s expected for %d %d", given, expected, n, k)
	}
}
