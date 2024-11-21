package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x+1)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	workers := 10000

	for i := 0; i < workers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100000; i++ {
		data <- i
	}
}
