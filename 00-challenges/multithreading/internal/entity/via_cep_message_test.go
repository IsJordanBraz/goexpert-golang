package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViaCepMessageAuto(t *testing.T) {
	cep := "01001000"
	logradouro := "Praça da Sé"
	complemento := ""
	unidade := ""
	bairro := "Sé"
	localidade := "São Paulo"
	uf := "SP"
	estado := "São Paulo"
	regiao := ""
	ibge := ""
	gia := ""
	ddd := ""
	siafi := ""

	message := NewViaCepMessage(cep, logradouro, complemento, unidade, bairro, localidade, uf, estado, regiao, ibge, gia, ddd, siafi)

	assert.NotNil(t, message)
	assert.Equal(t, "Praça da Sé", message.Logradouro)
	assert.Equal(t, "Sé", message.Bairro)
	assert.Equal(t, "São Paulo", message.Localidade)
	assert.Equal(t, "São Paulo", message.Estado)
	assert.Equal(t, "SP", message.Uf)
	assert.Equal(t, "01001000", message.Cep)
}
