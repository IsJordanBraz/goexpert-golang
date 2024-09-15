package main

import (
	"fmt"

	mathutils "github.com/IsJordanBraz/auth-service/math-utils"
)

func main() {
	sum := mathutils.Soma(10, 20)
	fmt.Println(sum)
}
