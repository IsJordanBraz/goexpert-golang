package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {
	var p fastjson.Parser
	jsonData := `{"nome": "jordan", "age": "25", "colors": [255, 100, 200] }`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("nome=%s\n", v.GetStringBytes("nome"))

	colors := v.GetArray("colors")
	for i, value := range colors {
		fmt.Printf("Index %d: %s\n", i, value)
	}
}
