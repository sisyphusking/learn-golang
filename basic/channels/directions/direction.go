package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg

}

//单向通道，pings只能发送数据，pongs只能接受数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "send msg....")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
