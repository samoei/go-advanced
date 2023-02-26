package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}()
	wg.Wait()
	fmt.Println("Elapsed: ", time.Since(now))
	fmt.Println("Main is now done and exiting")
	// no join
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done doing work")
}
