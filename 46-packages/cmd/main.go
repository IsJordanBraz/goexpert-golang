package main

import (
	"fmt"

	"github.com/IsJordanBraz/goexpert-golang/46-packages/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
