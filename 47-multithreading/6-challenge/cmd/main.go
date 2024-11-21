package main

import (
	"net/http"
	"time"
)

type BrasilApiMessage struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaCepMessage struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func getViaCepMessage(cep string) {

}

func getViaCepMessage(cep string) {

}

func main() {
	ch1 := make(chan BrasilApiMessage)
	ch2 := make(chan ViaCepMessage)

	cep := "01511010"

	go func() {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		msg1 := Message{id: 1, Msg: "RabbitMQ"}
		ch1 <- msg1
		time.Sleep(time.Second)
	}()

	// Kafka
	go func() {
		msg2 := Message{id: 2, Msg: "Kafka"}
		ch2 <- msg2
		time.Sleep(time.Second * 2)
	}()
}
