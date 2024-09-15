package main

import "fmt"

const a = "Hello"

type ID int

var (
	b bool = true
	c int
	d string
	e float64
	f ID
)

func main() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	x := "x"
	println(x)
	fmt.Printf("O tipo é %T", f)
	println(f)
	fmt.Printf("type is %T", e)
	fmt.Printf("type is %v", e)
	println(myArray[0])
	fmt.Println(len(myArray))

	for i, v := range myArray {
		fmt.Printf("Indice is %d and value is %d\n", i, v)
	}
}
