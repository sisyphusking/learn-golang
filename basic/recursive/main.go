package main

import "fmt"

func main() {
	rec(0)
	fmt.Println(fact(7))
}

// 递归1
func rec(i int) {
	if i == 10 {
		return
	}
	rec(i + 1)
	fmt.Printf("%d ", i)
}

//递归2
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
