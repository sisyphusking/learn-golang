package main

import (
	"io"
	"learn-golang/examples/pool"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25
	pooledResources = 5
)

//给每个连接分配一个唯一的id
var idCounter int32

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("close: connection", dbConn.ID)
	return nil
}

//工厂函数, 每次调用都会返回一个dbConn的实例
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: new connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("shut down")
	//关闭连接池
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	//从池里获得一个连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	//最后将资源放回到池里， 而不是关闭
	defer p.Release(conn)

	//模拟业务操作耗时
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	//  conn.(*dbConnection)类型判断，指定类型
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
