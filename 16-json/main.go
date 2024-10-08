package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	pureJson := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	err = json.Unmarshal(pureJson, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)
}
