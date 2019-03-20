package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}() //这样申明之后必须带上“（）”表示引用

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}
