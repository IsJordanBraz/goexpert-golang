package main

import "fmt"

func panic1() {
	panic("panic1")
}

func panic2() {
	panic("panic2")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			if r == "panic1" {
				fmt.Println("panic1 Recovered: ")
			} else {
				fmt.Println("Panic2 Recovered: ")
			}
		}
	}()

	panic2()
}
