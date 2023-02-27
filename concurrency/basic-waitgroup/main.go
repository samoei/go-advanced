package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // Declare
	wg.Add(3)             // how many threads should it wait for
	go func() {
		defer wg.Done() // tell the wait group that this thread is done
		fmt.Println("Thread 1")
	}()
	go func() {
		defer wg.Done() // tell the wait group that this thread is done
		fmt.Println("Thread 2")
	}()
	go func() {
		defer wg.Done() // tell the wait group that this thread is done
		fmt.Println("Thread 3")
	}()
	wg.Wait() // Wait untill all threads are done
	fmt.Println("Main is done waiting and now it can exit")
}
