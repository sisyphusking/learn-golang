package main

import "fmt"
import "log"

func main() {

	defer func() {
		log.Println(recover())
	}()

	test()

}

func test() {
	defer fmt.Println("test1")
	defer fmt.Println("test2")

	// panic会立即终端当前函数流程，执行defer ，而在main中defer调用了recover() 方法，捕获了panic提交的错误信息
	panic("this is a test")

	// 连续调用panic，只有一个会被recover()捕获
	panic("this is second test")
}
