package entity

type BrasilApiMessage struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func NewBrasilApiMessage(cep, state, city, neighborhood, street, service string) *BrasilApiMessage {
	return &BrasilApiMessage{
		Cep:          cep,
		State:        state,
		City:         city,
		Neighborhood: neighborhood,
		Street:       street,
		Service:      service,
	}
}
