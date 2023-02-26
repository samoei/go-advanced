package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})
	go func() {
		work()
		done <- struct{}{}
	}()
	<-done
	fmt.Println("Elapsed: ", time.Since(now))
	fmt.Println("Main is now done and exiting")
	// no join
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done doing work")
}
