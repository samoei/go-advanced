package main

import "fmt"

func main() {
	// create channels
	channel1 := make(chan string)
	channel2 := make(chan string)

	//insert messages to the channels
	go func() {
		channel1 <- "Message from channel 1"
	}()

	go func() {
		channel2 <- "Message from channel 2"
	}()

	//use select to get messages out of the channels
	select {
	case channel1Message := <-channel1:
		fmt.Println(channel1Message)
	case channel2Message := <-channel2:
		fmt.Println(channel2Message)

	}

}
