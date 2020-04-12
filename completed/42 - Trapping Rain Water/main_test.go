package main

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	height := []int{}
	if trap(height) != 0 {
		t.Errorf("Empty borders means zero volume")
	}
}

func TestZeroPeak(t *testing.T) {
	height := []int{0}
	if trap(height) != 0 {
		t.Errorf("Zero peak means zero volume")
	}
}

func TestPeak(t *testing.T) {
	height := []int{1}
	if trap(height) != 0 {
		t.Errorf("One peak means zero volume")
	}
}

func TestPlato(t *testing.T) {
	height := []int{1, 1, 1}
	if trap(height) != 0 {
		t.Errorf("Plato means zero volume")
	}
}

func TestNested(t *testing.T) {
	height := []int{3, 2, 1, 1, 2, 3}
	res := trap(height)
	if res != 6 {
		t.Errorf("3, 2, 1, 1, 2, 3 should five volume = 6, %d given", res)
	}
}

func TestStarterIsBiggerThanCloser(t *testing.T) {
	height := []int{1, 10, 0, 5, 0, 1, 0}
	res := trap(height)
	if res != 6 {
		t.Errorf("1, 10, 0, 1, 0 should five volume = 6, %d given", res)
	}
}

func TestStairwayToHeaven(t *testing.T) {
	height := []int{0, 1, 2, 3, 0, 0}
	res := trap(height)
	if res != 0 {
		t.Errorf("Stairway to heaven should five zero volume, %d given", res)
	}
}

func TestStairwayFromHeaven(t *testing.T) {
	height := []int{0, 3, 2, 1, 0, 0}
	res := trap(height)
	if res != 0 {
		t.Errorf("Stairway from heaven should five zero volume, %d given", res)
	}
}
