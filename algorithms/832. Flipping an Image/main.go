package main

import "fmt"

func main() {

	A := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}

	for _, a := range A {
		//先前后交换位置
		a = reverse(a)
		for i, b := range a {
			//修改0或者1
			//这里不能用b:=inverse(b)， 因为b属于值引用；
			a[i] = inverse(b)
		}

	}
	fmt.Println(A)

}

func reverse(B []int) []int {

	start := 0
	end := len(B) - 1
	for start < end {
		B[start], B[end] = B[end], B[start]
		start++
		end--
	}
	return B
}

func inverse(c int) int {
	if c == 1 {
		c = 0
	} else {
		c = 1
	}
	return c
}
