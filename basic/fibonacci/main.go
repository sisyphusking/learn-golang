package main

import (
	"fmt"
	"time"
)

func main() {
	//方法一
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	fmt.Println("***********")
	//方法二
	ch := make(chan int)
	n := 10
	go fibonacci2(n, ch)
	for i := range ch {
		fmt.Println(i)
	}
	time.Sleep(1 * time.Second)
}

// 闭包，会保留x,y的值
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func fibonacci2(n int, ch chan<- int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	//注意，这里需要手动关闭通道
	close(ch)
}
