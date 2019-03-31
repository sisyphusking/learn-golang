package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Println("worker", id, "started job", j) //job要等到有发送者才能运行
		time.Sleep(time.Second)                     //模拟要处理的任务
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2 //有缓冲的通道，不一定阻塞
	}
}

func main() {

	jobs := make(chan int, 100)    //发送任务
	results := make(chan int, 100) // 统计结果

	//分配工人，每个工人的任务详情，以及声明总的任务数jobs，完成后反馈结果results
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results) // 初始化的时候这个地方是阻塞的，因为没有传入jobs
	}

	// 下发任务
	for j := 1; j <= 5; j++ {
		fmt.Println("send job", j)
		jobs <- j
	}

	close(jobs) //直到通道中的数据被接收后，才开始关闭，这种主要适用于有缓冲的channel

	//这里循环的次数和上面要保持一致，如果是大于5次，那么最后会提示goroutine已经被asleep，会报错。
	for a := 1; a <= 5; a++ {
		//防止在goroutinew完成前，main函数终止；
		//如果注释掉下面这一行，那么结果都不会打印出worker，因为程序已经退出了
		//非常巧妙地设计
		<-results
		fmt.Println("closed")

		//除非在下面手动加上线程等待
		// time.Sleep(3 * time.Second)

	}
	fmt.Println("end...")
}
