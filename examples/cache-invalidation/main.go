package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/abusizhishen/do-once-while-concurrent/src"
)

// do-once-while-concurrent：https://github.com/abusizhishen/do-once-while-concurrent
//事实上不止于防止缓存穿透, do-once-while-concurrent 更准确的定位是重复资源过滤 ,
//在某讲座业务中，使用 do-once-while-concurrent 来避免同一时刻同一用户id 重复解析 、列表页 重复检索 、排序 等，减少了资源竞争，提高了整体的qps和稳定性。

//参考文章：https://mp.weixin.qq.com/s/s8i2wQXHaP9EDJ0uXIR54A
/*
do-once-while-concurrent中有三个主要方法,
- Req 方法 对具有同一资源标识的请求进行拦截
- Wait 方法 等待数据
- Release 方法 广播信号，数据已就位
*/

/*
我们的实际项目中有 两级缓存 ，一级 本地缓存 ，一级 redis ，如果都查询不到才会 读取mysql 或 调用中台接口 ,
本次只模拟 本地缓存失效 时, do-once-while-concurrent 对防止 缓存穿透 的处理(实际叫 重复资源过滤 更合理)
1.缓存失效时， 所有请求该缓存的请求会先调用 Req方法 对具有相同标签的重复请求进行拦截
2.只有第一个请求会 获取锁 ，执行读取redis操作
3.所有其他的线程 获取锁 失败，调用 Wait 方法, 等待第一个线程 执行结束
4.第一个线程读取到用户信息,写入本地缓存,通过 close(chan) 事件来 广播消息
5.其他线程收到消息，结束 等待 ，读取本地缓存，返回用户信息
*/

func main() {
	//并发do something
	for i := 0; i < 5; i++ {
		go doSomeThing()
	}

	//避免程序直接退出
	time.Sleep(time.Second * 5)
}

var once src.DoOnce

//模拟获取用户信息
func doSomeThing() {
	var userId = 12345
	var user, err = getUserInfo(userId)
	fmt.Println(user, err)
}

//example for usage
// 演示获取用户详情的过程，先从本地缓存读取用户,如果本地缓存不存在,就从redis读取
var keyUser = "user_%d"

func getUserInfo(userId int) (user UserInfo, err error) {
	user, err = userCache.GetUser(userId)
	if err == nil {
		return
	}

	log.Println(err)
	var requestTag = fmt.Sprintf(keyUser, userId)
	if !once.Req(requestTag) {
		log.Println("没抢到锁，等待抢到锁的线程执行结束。。。")
		once.Wait(requestTag)
		log.Println("等待结束:", requestTag)
		return userCache.GetUser(userId)
	}

	//得到资源后释放锁
	defer once.Release(requestTag)
	log.Println(requestTag, "获得锁，let's Go")

	//为演示效果，sleep
	time.Sleep(time.Second * 3)

	//redis读取用户信息
	log.Println("redis读取用户信息:", userId)
	user, err = getUserInfoFromRedis(userId)
	if err != nil {
		return
	}

	//用户写入缓存
	log.Println("用户写入缓存:", userId)
	userCache.setUser(user)
	return
}

//用户信息缓存
type UserCache struct {
	Users map[int]UserInfo
	sync.RWMutex
}

type UserInfo struct {
	Id   int
	Name string
	Age  int
}

var userCache UserCache
var errUserNotFound = errors.New("user not found in cache")

func (c *UserCache) GetUser(id int) (user UserInfo, err error) {
	c.RLock()
	defer c.RUnlock()
	var ok bool
	user, ok = userCache.Users[id]
	if ok {
		return
	}

	return user, errUserNotFound
}

func (c *UserCache) setUser(user UserInfo) {
	c.Lock()
	defer c.Unlock()
	if c.Users == nil {
		c.Users = make(map[int]UserInfo)
	}

	c.Users[user.Id] = user
	return
}

func getUserInfoFromRedis(id int) (user UserInfo, err error) {
	user = UserInfo{
		Id:   12345,
		Name: "abusizhishen",
		Age:  18,
	}
	return
}
