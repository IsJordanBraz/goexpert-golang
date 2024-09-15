package main

import "fmt"

func main() {
	salarios := map[string]int{"Jordan": 1000, "Nathalia": 2000}
	jordan := salarios["Jordan"]
	fmt.Print("value for map jordan is: ", jordan)
	delete(salarios, "Jordan")

	sal := make(map[string]int)
	print("\nsal: ", sal)
	sal["jordan"] = 10

	for nome, salario := range sal {
		fmt.Printf("\n o salario de %s é %d\n", nome, salario)
	}

	for _, salario := range sal {
		fmt.Printf("\n o salario é %d\n", salario)
	}
}
