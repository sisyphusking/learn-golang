package main

import "fmt"

func main() {
	fmt.Println(square(demo, 2))
}

func demo(y int) int {
	return y * y
}

func square(f func(int) int, x int) int {
	return f(x * x)
}
