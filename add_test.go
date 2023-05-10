// Test for the add function in the main package.

package datastructuresandalgorithms

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Expected 1+2 to equal 3")
	} else {
		t.Log("Success!")
	}
}
