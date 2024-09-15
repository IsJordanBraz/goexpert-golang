package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(51, 2))
	fmt.Println(sumAll(51, 2, 3, 1, 1, 1, 100))
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}

func sumAll(numbers ...int) int {
	total := 0
	for _, numero := range numbers {
		total += numero
	}
	return total
}
