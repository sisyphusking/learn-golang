package main

import (
	"learn-golang/examples/work/work"
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
	//同步的方式
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
