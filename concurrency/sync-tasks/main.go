package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	task1()
	task2()
	task3()
	task4()
	task5()
	fmt.Println("Elapsed: ", time.Since(now))
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 1")
}

func task2() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Task 2")
}

func task3() {
	fmt.Println("Task 3")
}

func task4() {
	time.Sleep(10 * time.Second)
	fmt.Println("Task 4")
}

func task5() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Task 5")
}
