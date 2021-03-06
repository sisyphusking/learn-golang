// This sample program demonstrates how to use a buffered
// channel to main on multiple tasks with a predefined number
// of goroutines.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // Number of goroutines to use.
	taskLoad         = 10 // Amount of main to process.
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

// init is called to initialize the package by the
// Go runtime prior to any other code being executed.
func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().Unix())
}

// main is the entry point for all Go programs.
func main() {
	// Create a buffered channel to manage the task load.
	tasks := make(chan string, taskLoad)

	// Launch goroutines to handle the main.
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// Add a bunch of main to get done.
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// Close the channel so the goroutines will quit
	// when all the main is done.
	//关闭通道的代码非常重要。当通道关闭后，goroutine 依旧可以从通道接收数据， 但是不能再向通道里发送数据。
	//能够从已经关闭的通道接收数据这一点非常重要，因为这允许通 道关闭后依旧能取出其中缓冲的全部值，而不会有数据丢失。
	//从一个已经关闭且没有数据的通道 里获取数据，总会立刻返回，并返回一个通道类型的零值。如果在获取通道时还加入了可选的标 志，就能得到通道的状态信息。
	close(tasks)

	// Wait for all the main to get done.
	wg.Wait()
}

// worker is launched as a goroutine to process main from
// the buffered channel.
func worker(tasks chan string, worker int) {
	// Report that we just returned.
	defer wg.Done()

	for {
		// Wait for main to be assigned.
		task, ok := <-tasks
		if !ok {
			// This means the channel is empty and closed.
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// Display we are starting the main.
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// Randomly wait to simulate main time.
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// Display we finished the main.
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
