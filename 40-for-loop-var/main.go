package main

import "fmt"

func main() {
	done := make(chan bool)
	values := []string{"a", "b", "c"}
	for _, v := range values {
		// v := v  - GO 1.22
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for range values {
		<-done
	}
}
