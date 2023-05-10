package datastructuresandalgorithms

func BinarySearch(haystack []int, needle int) int {
	lo := 0
	hi := len(haystack)

	for lo < hi {
		mid := lo + (hi-lo)/2 // Floor int rounding
		val := haystack[mid]
		if val == needle {
			return mid
		}
		if val < needle {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return -1
}
