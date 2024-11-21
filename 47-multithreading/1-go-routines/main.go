package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go task("A")
	go task("B")
	time.Sleep(5 * time.Second)
}
