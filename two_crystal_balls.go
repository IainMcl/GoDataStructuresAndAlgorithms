package datastructuresandalgorithms

import (
	"math"
)

func TwoCrystalBalls(arr []bool) int {
	// Walk up in setps of sqrt(2) untill the first ball breaks
	// Then linear search from the previous known point of safety
	step := int(math.Sqrt(float64(len(arr)))) // int will do the floor rounding

	for i := 0; i < len(arr); i += step {
		if arr[i] {
			// First ball has broken now linear search from the previous step
			j := i - step
            if j <= 0{
                j = 0
            }
			for ; j < i; j++ {
				if arr[j] {
                    if j == 0{
                        return 0
                    }
					return j - 1
				}
			}
		}
	}
	return -1
}
