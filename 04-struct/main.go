package main

import "fmt"

type Address struct {
	Street string
	Number int
}

type State struct {
	City string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
	State State
}

func (c Client) Inativate() {
	c.Active = false
}

func main() {
	client := Client{
		Name:   "Jordan",
		Age:    18,
		Active: true,
	}

	client.Street = "São Paulo"
	client.Inativate()

	fmt.Println(client)
}
