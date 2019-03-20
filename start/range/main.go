package main

import "fmt"

func main() {

	data := [3]int{10, 20, 30}
	fmt.Printf("&data: %p\n", &data)

	// x 是data的复制，而不是data本身
	for i, x := range data {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		// data始终没变，只是range中循环取的是data的复制值
		fmt.Printf("&data1: %p\n", &data)
		fmt.Printf("x: %d, data: %d\n", x, data[i])
	}

	fmt.Println("-----------")
	fmt.Println("data[:]:", data[:])
	for i, x := range data[:] {
		if i == 0 {
			data[0] += 100
			data[1] += 200
			data[2] += 300
		}
		// x第一次循环的时候是从data的复制值中取的，和执行了加法操作的data不一样，所以这里是110
		fmt.Printf("x: %d, data: %d\n", x, data[i])
	}

}
