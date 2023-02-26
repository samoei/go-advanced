package main

import (
	"fmt"
	"time"
)

func main() {
	go work() //fork
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main is now done and exiting")
	// no join
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done doing work")
}
