package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("pool has been closed ")

func New(fn func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}
	return &Pool{
		//有缓冲通道，如果是无缓冲通道就达不到效果
		resources: make(chan io.Closer, size),
		factory:   fn,
	}, nil
}

//从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire: ", "shared resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	//没有空闲资源，提供一个新的
	//执行用户提供的工厂函数，并且创建并返回一个新资源
	default:
		log.Println("Acquire: ", "new resource")
		return p.factory()
	}
}

//释放一个使用后的资源到池中
func (p *Pool) Release(r io.Closer) {
	//加锁
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		//如果池关闭了，销毁这个资源
		r.Close()
		return
	}

	select {
	//资源放入池中
	case p.resources <- r:
		log.Println("Release: ", "In Queue")
	//如果资源池队列已经满了
	default:
		log.Println("Release: ", "Closing")
		r.Close()
	}
}

//release和close需要都加锁，非常经典，适合于两个协程操作同一个变量
//第一，可以保护读取closed标志的行为，保证同一时刻不会有其他 goroutine 调用 Close 方法写同一个标志。
//第二，我们不想往一个已经关闭的通道里发送数据，因为那样会引起崩溃。如果 closed 标志是 true，我们 就知道 resources 通道已经被关闭。

func (p *Pool) Close() {
	//加锁
	p.m.Lock()
	defer p.m.Unlock()

	//如果池已经被关闭，什么都不做
	if p.closed {
		return
	}

	//将池关闭
	p.closed = true
	//在清空通道里的资源之前，将通道关闭
	//这里一定要close，否则后面range 会死锁
	// resource是有缓冲通道，只有在通道为空或者满了的时候才阻塞，如果这里不手动关闭，resource为空，for循环会一直阻塞
	close(p.resources)
	//将缓存的资源清空
	for r := range p.resources {
		r.Close()
	}
}
