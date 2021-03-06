package main

import (
	"learn-golang/examples/work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"jason",
	"ted",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {

	//非常经典的用法
	//类似于同步的方式，创建一个工作池，一次只能接收一条数据
	p := work.New(2)
	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{name: name}
			//创建多个goroutine运行
			go func() {
				//将任务提交到pool中
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	//等待所有goroutine完成
	p.Shutdown()
}
