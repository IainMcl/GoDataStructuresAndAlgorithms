package datastructuresandalgorithms

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if LinearSearch(arr, 3) != 2 {
		t.Error("LinearSearch failed")
	}
}

func TestLinearSearch2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	if LinearSearch(arr, 6) != -1 {
		t.Error("LinearSearch failed")
	}
}
