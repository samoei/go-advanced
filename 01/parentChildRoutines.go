package main

import (
	"fmt"
	"time"
)

func longRunningTask(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			startCounter()
		}
	}
}

func startCounter() {
	i := 0
	for {
		i++
		fmt.Println(i)
	}
}

func main() {
	//create a channel to be used to communicate with children and close them when parent terminates
	done := make(chan bool)

	go longRunningTask(done)

	time.Sleep(time.Millisecond * 1)

	close(done)
}
