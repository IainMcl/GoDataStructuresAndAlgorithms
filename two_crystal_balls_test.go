package datastructuresandalgorithms

import (
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	arr := []bool{false, false, false, false, true, true, true}
	max := TwoCrystalBalls(arr)
	if max != 3 {
		t.Error("Expected to have a minimum at 3")
	}
}

func TestTwoCrystalBallsAllFalse(t *testing.T) {
	arr := []bool{false, false, false, false, false, false, false, false, false, false, false, false}
	max := TwoCrystalBalls(arr)
	if max != -1 {
		t.Error("Expected no result to be found -1")
	}
}

func TestTwoCrystalBallsAllTrue(t *testing.T) {
	arr := []bool{true, true, true, true, true, true, true, true, true, true, true, true}
	max := TwoCrystalBalls(arr)
	if max != 0 {
		t.Error("Expected all to work 0")
	}

}
