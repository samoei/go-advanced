package main

import "fmt"

func main() {
	//create a buffered channel
	charsChannel := make(chan string, 3)

	//characters to store in the channel
	chars := []string{"a", "b", "c"}

	//insert the characters into the channel
	for _, c := range chars {
		select {
		case charsChannel <- c:
		}
	}

	//close the channel
	close(charsChannel)

	//loop over the channel and print the results

	for r := range charsChannel {
		fmt.Println(r)
	}

}
