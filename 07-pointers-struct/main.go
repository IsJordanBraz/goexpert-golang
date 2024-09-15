package main

import "fmt"

type Cliente struct {
	nome string
}

func (c *Cliente) Andou() {
	c.nome = "Jordan Braz"
	fmt.Printf("O cliente %v Andou\n", c.nome)
}

func main() {
	jordan := Cliente{
		nome: "Jordan",
	}
	jordan.Andou()
	fmt.Printf("o valor da struct com nome %v", jordan.nome)
}
