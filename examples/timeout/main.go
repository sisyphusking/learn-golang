package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
实现超时控制不能只是select-case, 因为也要考虑超时的时候返回，也要将正在执行的goroutine取消掉。
如果各个goroutine之间有依赖(获取的锁未释放)的话，还会导致连锁反应。
最好是使用context

下面的方案不合理的地方：go func() 没有动态传入变量
*/

var (
	mutex sync.Mutex
	id    int
)

func dosomething(ctx context.Context, val int) {
	mutex.Lock()
	defer mutex.Unlock()
	select {
	case <-ctx.Done():
		fmt.Println(time.Now(), "op timeout", val)
		return
	default:
		//....
		time.Sleep(time.Second)
		id = val
	}

}

func main() {

	for i := 0; i < 3; i++ {
		done := make(chan bool)
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(i)*time.Second)
			defer cancel()
			dosomething(ctx, i)
			done <- true
		}()
		select {
		case res := <-done:
			fmt.Println(time.Now(), res, id)
		case <-time.After(time.Duration(i) * time.Second):
			fmt.Println(time.Now(), "timeout ", i)
		}
	}
}
