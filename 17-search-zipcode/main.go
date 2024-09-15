package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			println("Error request")
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			println("Error reading")
		}
		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			println("Error unmarshal req")
		}
		fmt.Println(data)
		file, err := os.Create("cidade.txt")
		if err != nil {
			println("error creating file")
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("Cep: %s", data.Cep))
	}
}
