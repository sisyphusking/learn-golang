package main

import (
	"fmt"
	"math/bits"
)

func hammingDistance(x int, y int) int {

	//异或运算
	return bits.OnesCount(uint(x ^ y))
}

func main() {
	a := 2
	b := 3
	fmt.Println(hammingDistance(a, b))
}
