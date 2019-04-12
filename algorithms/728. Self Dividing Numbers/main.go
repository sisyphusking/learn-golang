package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}

/**
a self-dividing number is not allowed to contain the digit zero，不允许包含0
**/
func selfDividingNumbers(left int, right int) []int {
	var res []int
	for num := left; num <= right; num++ {
		tmp := num
		for tmp > 0 {
			//前面一个条件判断过滤掉包含0的整数
			//第二个条件是用原来的数依次除以位上的数
			if tmp%10 == 0 || num%(tmp%10) != 0 {
				break
			}
			tmp = int(tmp / 10)
		}

		if tmp == 0 {
			res = append(res, num)
		}
	}
	return res

}
