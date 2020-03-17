package main

import "sync"

// call is an in-flight or completed Do call
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}

/*
最想抢到锁的goroutine 执行到g.m[key] = c后立即释放锁，然后调动fn()函数，fn()是无缓冲通道，会一直阻塞，这时候g.m已经释放了锁，
其他的goroutine有能有一个抢到锁，执行到c, ok := g.m[key]; ok时，下一步就是释放锁，然后c.wg.Wait()就一直阻塞等第一个抢占到锁的goroutine结束，
这个时候锁已经释放了，剩下的交给其他goroutine来继续运行，但是也是阻塞到c.wg.Wait()这里。
一直到所有的goroutine都被阻塞。

然后fn()中的c有消息传递，第一个goroutine获取到了fn的返回，传给c, 这时候wg.Done() , c.wg.Wait()这里不在阻塞，且c中已经有了函数的返回结果。
所以其他goroutine都能在return c.val, c.err这里获取到结果。从而保证了fn()只被执行了一次。

最后就是第一个goroutine继续往下执行(不一定是最后执行)，重新上锁并且删除key  然后获取到了返回。
*/
// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	//当第一个goroutine执行到这里的时候，其他的都在c.wg.Wait()这里运行下去，他们已经获取到了c.val, c.err, c, ok := g.m[key]
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
