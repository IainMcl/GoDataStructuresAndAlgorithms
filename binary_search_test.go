package datastructuresandalgorithms

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	index := BinarySearch(arr, 5)
	if index != 4 {
		t.Error("BinarySearch failed. Got", index, "Expected 4")
	}
}

func TestBinarySearch2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	index := BinarySearch(arr, 10)
	if index != -1 {
		t.Error("BinarySearch failed. Got", index, "Expected -1")
	}
}
