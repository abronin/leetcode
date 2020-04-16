package main

import "testing"

func TestEmpty(t *testing.T) {
	name := ""
	typed := ""
	expected := true
	given := isLongPressedName(name, typed)

	if given != expected {
		t.Error(
			"Expected:", expected,
			"Given", given,
			"Name", name,
			"Typed", typed,
		)
	}
}

func TestEmptyButTyped(t *testing.T) {
	name := ""
	typed := "asd"
	expected := true
	given := isLongPressedName(name, typed)

	if given != expected {
		t.Error(
			"Expected:", expected,
			"Given", given,
			"Name", name,
			"Typed", typed,
		)
	}
}

func TestAlex(t *testing.T) {
	name := "alex"
	typed := "aaleex"
	expected := true
	given := isLongPressedName(name, typed)

	if given != expected {
		t.Error(
			"Expected:", expected,
			"Given", given,
			"Name", name,
			"Typed", typed,
		)
	}
}

func TestSaeed(t *testing.T) {
	name := "saeed"
	typed := "ssaaedd"
	expected := false
	given := isLongPressedName(name, typed)

	if given != expected {
		t.Error(
			"Expected:", expected,
			"Given", given,
			"Name", name,
			"Typed", typed,
		)
	}
}

func TestLeelee(t *testing.T) {
	name := "leelee"
	typed := "lleeelee"
	expected := true
	given := isLongPressedName(name, typed)

	if given != expected {
		t.Error(
			"Expected:", expected,
			"Given", given,
			"Name", name,
			"Typed", typed,
		)
	}
}

func TestSame(t *testing.T) {
	name := "laiden"
	typed := "laiden"
	expected := true
	given := isLongPressedName(name, typed)

	if given != expected {
		t.Error(
			"Expected:", expected,
			"Given", given,
			"Name", name,
			"Typed", typed,
		)
	}
}
