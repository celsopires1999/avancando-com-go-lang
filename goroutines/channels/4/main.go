package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		channel <- 10
	}()

	fmt.Println(<-channel)
}
