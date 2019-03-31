package main

import "fmt"

func main() {
	//指定第四个元素为1
	slice := []int{3: 1}
	fmt.Println(slice)

	//新的切片和原切片共用的是一个底层数组，所以当修改的时候，底层数组的值就会被改变，所以原切片的值也改变了。当然对于基于数组的切片也一样的。
	slice1 := []int{1, 2, 3, 4, 5}
	newSlice := slice1[1:3]

	// 对于底层数组容量是k的切片slice[i:j]来说
	// 长度：j-i
	// 容量: k-i
	newSlice[0] = 10

	fmt.Println(slice1)
	fmt.Println(newSlice)
	//计算长度和容量
	fmt.Printf("newSlice长度:%d,容量:%d", len(newSlice), cap(newSlice))

}
