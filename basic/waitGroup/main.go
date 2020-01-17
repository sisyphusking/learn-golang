package main

import (
	"sync"
	"time"
)

type data struct {
	sync.Mutex
}

//当Mutex作为匿名字段时，相关方法必须实现为pointer-receiver，否则会因复制导致锁机制失效。
//这里的接收者一定要是指针类型，要不然锁机制会失效，换成data，锁会失败
//重点，这里使用*data，说明在调用的时候传递的是指针，所以锁机制会生效
func (d *data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var d data

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}
