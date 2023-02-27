package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Number of cores", runtime.NumCPU())
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(i+1, &wg)
	}
	wg.Wait()
	fmt.Println("Main is done")
}

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task", id, "is done")
}
