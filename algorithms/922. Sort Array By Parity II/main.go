package main

import (
	"fmt"
)

func main() {
	a := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(a))
}

func sortArrayByParityII(A []int) []int {
	//用双指针来控制
	evenIndex := 0
	oddIndex := 1

	ret := make([]int, len(A))
	for _, elem := range A {
		if elem%2 == 0 {
			ret[evenIndex] = elem
			evenIndex += 2
		} else {
			ret[oddIndex] = elem
			oddIndex += 2
		}

	}
	return ret
}

func sortArrayByParityII1(A []int) []int {
	length := len(A)

	for even, odd := 0, 1; even < length && odd < length; {
		for even < length && A[even]%2 == 0 {
			even += 2
		}
		if even < length && odd < length {
			A[even], A[odd] = A[odd], A[even]
			odd += 2
		}
	}

	return A
}
