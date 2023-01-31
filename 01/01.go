package main

import (
	"fmt"
	"time"
)

func anotherFunc(str string) {
	fmt.Println(str)
}

func main() {
	go anotherFunc("1")
	go anotherFunc("2")
	go anotherFunc("3")
	go anotherFunc("4")

	time.Sleep(time.Second * 1)

	fmt.Println("Main is done")
}
