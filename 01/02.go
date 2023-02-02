package main

import "fmt"

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		channel1 <- "Message from channel 1"
	}()

	go func() {
		channel2 <- "Message from channel 2"
	}()

	select {
	case channel1Message := <-channel1:
		fmt.Println(channel1Message)
	case channel2Message := <-channel2:
		fmt.Println(channel2Message)

	}

}
