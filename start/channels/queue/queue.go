package main

import (
	"fmt"
	"time"
)

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // queue中的数据如果没有被接收，通道就不会真正关闭，这个例子展示了我们可以关闭一个非空的channels，但是前提是里面的数据被接收

	time.Sleep(2 * time.Second)
	// for range类似实现了"<-queue"的功能
	for elem := range queue {
		fmt.Println(elem)
	}
}
