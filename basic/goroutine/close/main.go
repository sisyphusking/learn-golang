package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ready := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			println(id, ":ready.") // 运动员准备就绪
			<-ready                // 等待发令,有数据传输或者通道关闭了，那么就停止阻塞
			println(id, ":running...")
		}(i)
	}

	time.Sleep(time.Second)
	println("Ready?Go!")

	close(ready) // 砰！

	wg.Wait()
}
