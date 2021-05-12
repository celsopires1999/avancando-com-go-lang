package main

import (
	"fmt"
	"math/rand"
	"time"
)

var result int

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		z := result
		z++
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		result = z
		fmt.Println(name, "->", i, "Partial result: ", result)
	}
}

func main() {

	go runProcess("P1", 20)
	go runProcess("P2", 20)

	var s string
	fmt.Scanln(&s)

	fmt.Println("Final result: ", result)

}
