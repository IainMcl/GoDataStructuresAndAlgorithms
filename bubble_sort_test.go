package datastructuresandalgorithms

import (
	"fmt"
	"testing"
)

func compareSlice(slice1, slice2 *[]int) bool {
	if len(*slice1) != len(*slice2) {
		return false
	}

	for i, v := range *slice1 {
		if v != (*slice2)[i] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	arr := []int{4, 5, 6, 2, 3, 7, 65, 90, 0}
	sorted := []int{0, 2, 3, 4, 5, 6, 7, 65, 90}

	fmt.Println(arr)
	BubbleSort(&arr)
	fmt.Println(arr)

	if !compareSlice(&arr, &sorted) {
		t.Error("Failed to sort")
	}
}
