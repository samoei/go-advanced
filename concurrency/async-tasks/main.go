package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	now := time.Now()
	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)
	go task5(done)
	<-done
	<-done
	<-done
	<-done
	<-done
	fmt.Println("Elapsed: ", time.Since(now))
}

func task1(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 1")
	done <- struct{}{}
}

func task2(done chan struct{}) {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Task 2")
	done <- struct{}{}
}

func task3(done chan struct{}) {
	fmt.Println("Task 3")
	done <- struct{}{}
}

func task4(done chan struct{}) {
	time.Sleep(10 * time.Second)
	fmt.Println("Task 4")
	done <- struct{}{}
}

func task5(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 5")
	done <- struct{}{}
}
