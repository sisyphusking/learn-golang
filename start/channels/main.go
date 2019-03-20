package main

// 每个子目录中只能存在一个package
// 引入包的时候，go会使用子目录名作为包的路径，而你在代码中真正使用时，却要使用你package的名称。
import (
	"gogogo/channels/channel"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("...开始执行任务...")

	timeout := 3 * time.Second
	r := channel.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case channel.ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case channel.ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}
	log.Println("...任务执行结束...")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("正在执行任务%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
