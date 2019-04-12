package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}

func selfDividingNumbers(left int, right int) []int {
	//func to check for is a self Dividing number
	var sd = func(v int) bool {
		//将整数转成字符串
		test := strconv.Itoa(v)
		for _, k := range test {
			//依次取出每个数字字符串，然后转成对应的整型
			n, _ := strconv.Atoi(string(k))
			if k == '0' || v%n != 0 {
				return false
			}
		}
		return true
	}
	var r = []int{}

	//loop over the bounds
	for i := left; i <= right; i++ {
		if sd(i) {
			r = append(r, i)
		}

	}
	return r
}
