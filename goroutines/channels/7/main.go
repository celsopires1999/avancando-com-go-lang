package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	x := funnel(generateMsg("Porco"), generateMsg("Verdão"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
		var s string
		fmt.Scanln(&s)
	}
	fmt.Println("Finished...")
}

func generateMsg(s string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 0; ; i++ {
			fmt.Println("Gerado:", s, "contador:", i)
			t := time.Now()
			channel <- fmt.Sprintf("String %s - Value %d - Segundo %d", s, i, t.Second())
			time.Sleep(time.Duration(rand.Intn(255)) * time.Millisecond)
		}
	}()
	return channel
}

func funnel(channel1, channel2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- <-channel1
		}
	}()
	go func() {
		for {
			channel <- <-channel2
		}
	}()
	return channel
}
