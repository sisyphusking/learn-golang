package main

import (
	"fmt"
	"strconv"
)

func main() {

	var orig string = "666"
	var an int
	var newS string

	//将字符型数组转化为整型
	an, _ = strconv.Atoi(orig)
	fmt.Printf("the integer is : %d\n", an)

	an += 5
	//返回数字所表示的字符串类型的十进制数
	newS = strconv.Itoa(an)
	fmt.Printf("the new string is : %s\n", newS)
}
