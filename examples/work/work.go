package work

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// 新建一个工作池，等待传入p.main，一个任务处理完后再开始另外一个
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	//单独的WaitGroup
	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		//使用了go func()就不会阻塞
		go func() {
			for w := range p.work {
				//传入一个实现了Worker接口的实例
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

//提交工作到工作池
// work是无缓冲通道，调用者必须等到工作池里的某个goroutine接收到这个值才会返回
func (p *Pool) Run(w Worker) {
	p.work <- w
}

//shutdown等待所有goroutine停止工作
func (p *Pool) Shutdown() {
	//因为用了for range , 所以这里一定要有close
	close(p.work)
	//等待所有的goroutine完成， 也就是跳出了for range循环，执行p.wg.Done()后
	p.wg.Wait()
}
