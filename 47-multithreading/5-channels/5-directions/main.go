package main

import "fmt"

func send(name string, hello chan<- string) {
	hello <- name
}

func read(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go send("Hello", hello)
	read(hello)
}
