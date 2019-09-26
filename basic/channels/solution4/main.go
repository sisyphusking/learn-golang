package main

import (
	"fmt"
	"time"
)

func B(quit chan<- string) {
	fmt.Println("B crraied out")
	quit <- "B"
}

func A(quit chan<- string, finished chan<- bool) {
	// 模拟耗时任务
	time.Sleep(time.Second * 1)
	fmt.Println("A crraied out")
	finished <- true
	quit <- "A"
}

func C(quit chan<- string, finished <-chan bool) {
	// 在A没有执行完之前，finished获取不到数据，会阻塞
	<-finished
	fmt.Println("C crraied out")
	quit <- "C"
}

func main() {

	quit := make(chan string)
	defer close(quit)
	finished := make(chan bool)
	defer close(finished)

	go A(quit, finished)
	go B(quit)
	go C(quit, finished)

	//这里需要多次输出quit, 因为上面向channle里输入了三次
	fmt.Println(<-quit)
	fmt.Println(<-quit)
	fmt.Println(<-quit)

	//没close的通道不能用for range
	//从quit中读数据不能使用for-range语法，不然程序会出现死锁
	//原因很简单，程序中quit通道没有被close，A、B、C运行完了，
	//但是Go的主协程还在for循环中继续等待，所以阻塞了，所有Go协程都阻塞了，进入了死锁状态
	// for res := range quit {
	// 	fmt.Println(res)
	// }

	//具体可以看下面的test例子， 有加上close，所以可以使用range
}

func test() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()

	for x := range ch {
		fmt.Println(x)
	}
}
