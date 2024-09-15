package main

import "fmt"

type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Println(c)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	jordan := Cliente{
		Nome:  "Jordan",
		Idade: 30,
		Ativo: true,
	}
	Desativacao(jordan)
}
