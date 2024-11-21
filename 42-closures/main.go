package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sumAll(1, 2) * 2
	}()
	fmt.Println(total)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, numero := range numbers {
		total += numero
	}
	return total
}
