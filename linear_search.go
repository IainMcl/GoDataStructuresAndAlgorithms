package datastructuresandalgorithms

func LinearSearch(haystack []int, needle int) int{
    for i, val := range(haystack){
        if val == needle{
            return i
        }
    }
    return -1
}
