package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 1, 4, 2, 1, 3}
	fmt.Println(heightChecker(a))
}
func heightChecker(heights []int) int {

	// newInt:= heights
	// sort.Ints(newInt)
	newInt := make([]int, len(heights))
	copy(newInt, heights)
	sort.Ints(newInt)

	count := 0
	for i := 0; i < len(heights); i++ {
		if newInt[i] != heights[i] {
			count++
		}
	}
	return count
}
