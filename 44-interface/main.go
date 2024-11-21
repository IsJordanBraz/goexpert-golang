package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

type Person interface {
	Inativate()
}

func Disabled(person Person) {
	person.Inativate()
}

func (c Client) Inativate() {
	c.Active = false
}

type Company struct {
	Name string
}

func (c Company) Disabled() {
}

func main() {
	client := Client{
		Name:   "Jordan",
		Age:    18,
		Active: true,
	}

	client.Inativate()
	company := Company{}

	Disabled(company)

	fmt.Println(client)
}
