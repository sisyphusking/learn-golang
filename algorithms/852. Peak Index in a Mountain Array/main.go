package main

import "fmt"

func main() {
	var a = []int{0, 2, 1, 0}
	fmt.Println(peakIndexInMountainArray(a))
}

func peakIndexInMountainArray(A []int) int {
	// 0 < i < A.length - 1
	for i := range A {
		if i != len(A)-1 && A[i] > A[i+1] {
			return i
		}
	}
	return 0
}
