package main

import "fmt"

//使用二分查找法
func main() {
	var a = []int{0, 2, 1, 0}
	fmt.Println(peakIndexInMountainArray(a))
}

func peakIndexInMountainArray(A []int) int {
	var medium int
	left, right := 0, len(A)-2
	for left < right {
		medium = (right + left) / 2
		if A[medium] > A[medium+1] {
			right = medium
		} else {
			left = medium + 1
		}
	}
	return left
}
