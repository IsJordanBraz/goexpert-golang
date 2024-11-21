package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string) // empty

	// Thread 2
	go func() {
		channel <- "Hello World" //full
	}()

	msg := <-channel //empty
	fmt.Println(msg)
}
