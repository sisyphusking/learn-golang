package main

import "fmt"

func printer(c chan int) {
	//无线循环
	for {
		//从channel中获取一个数据
		data := <-c
		//如果数据为0就跳出循环
		if data == 0 {
			break
		}
		fmt.Println(data)
	}
	//通知main已经结束了循环
	c <- 0

}

func main() {
	c := make(chan int)
	go printer(c)

	for i := 1; i <= 10; i++ {
		//将数据通过channel投送给printer
		c <- i
	}
	//通知printer结束循环，跳出循环体
	c <- 0

	//等待printer结束
	<-c

}
