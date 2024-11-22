package entity

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBrasilApiMessage(t *testing.T) {
	cep := "01001000"
	state := "SP"
	city := "São Paulo"
	neighborhood := "Sé"
	street := "Praça da Sé"
	service := "brasil_api"

	message := NewBrasilApiMessage(cep, state, city, neighborhood, street, service)

	assert.NotNil(t, message)
	assert.Equal(t, "01001000", message.Cep)
	assert.Equal(t, "SP", message.State)
	assert.Equal(t, "São Paulo", message.City)
	assert.Equal(t, "Sé", message.Neighborhood)
	assert.Equal(t, "Praça da Sé", message.Street)
	assert.Equal(t, "brasil_api", message.Service)
}

func TestNewBrasilApiMessageGeneratedByExtension(t *testing.T) {
	type args struct {
		cep          string
		state        string
		city         string
		neighborhood string
		street       string
		service      string
	}
	tests := []struct {
		name string
		args args
		want *BrasilApiMessage
	}{
		{
			name: "TestNewBrasilApiMessage",
			args: args{
				cep:          "01001000",
				state:        "SP",
				city:         "São Paulo",
				neighborhood: "Sé",
				street:       "Praça da Sé",
				service:      "brasil_api",
			},
			want: &BrasilApiMessage{
				Cep:          "01001000",
				State:        "SP",
				City:         "São Paulo",
				Neighborhood: "Sé",
				Street:       "Praça da Sé",
				Service:      "brasil_api",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBrasilApiMessage(tt.args.cep, tt.args.state, tt.args.city, tt.args.neighborhood, tt.args.street, tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBrasilApiMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
