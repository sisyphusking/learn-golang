package main

import (
	"learn-golang/basic/runner"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("start main")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(1)
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout")
			os.Exit(2)
		}
	}
	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("process task %d", id)
		time.Sleep(1 * time.Second)
	}
}
