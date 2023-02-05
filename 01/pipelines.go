package main

import "fmt"

func ingress(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func convertToSquares(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for r := range in {
			out <- r * r
		}
		close(out)
	}()
	return out
}

func main() {
	//input
	input := []int{1, 2, 3, 4, 5, 6}

	//stage 1
	dataChannel := ingress(input)

	//stage 2
	finalChannel := convertToSquares(dataChannel)

	//output

	for r := range finalChannel {
		fmt.Println(r)
	}
}
