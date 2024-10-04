package main

import (
	"fmt"
	"time"
)

func sayHi(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Hello"
}

func main() {
	ch := make(chan string)

	go sayHi(ch)

	message := <-ch

	fmt.Println(message)
}
