package main

import (
	"fmt"
	"learn-golang/examples/pubsub"
	"strings"
	"time"
)

/*
在发布订阅模型中，每条消息都会传送给多个订阅者。发布者通常不会知道、也不关心哪一个订阅者正在接收主题消息。
订阅者和发布者可以在运行时动态添加，是一种松散的耦合关系，这使得系统的复杂性可以随时间的推移而增长。
在现实生活中，像天气预报之类的应用就可以应用这个并发模式
*/

func main() {

	//创建一个发布者
	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	//所有的都订阅， 返回一个通道
	all := p.Subscribe()

	//topic过滤器，只订阅包含golang的推送
	//返回一个通道
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	//发布两条消息
	//不包含golang字符串的会被golang订阅者过滤掉
	p.Publish("hello,  world!")
	p.Publish("hello, golang!")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
}
