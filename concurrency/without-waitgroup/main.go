package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Number of cores", runtime.NumCPU())
	for i := 0; i < 10; i++ {
		go work(i + 1)
	}
	time.Sleep(time.Second)
	fmt.Println("Main is done")
}

func work(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task", id, "is done")
}
