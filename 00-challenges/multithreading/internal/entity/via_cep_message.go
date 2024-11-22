package entity

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

func NewViaCepMessage(cep, logradouro, complemento, unidade, bairro, localidade, uf, estado, regiao, ibge, gia, ddd, siafi string) *ViaCepMessage {
	return &ViaCepMessage{
		Cep:         cep,
		Logradouro:  logradouro,
		Complemento: complemento,
		Unidade:     unidade,
		Bairro:      bairro,
		Localidade:  localidade,
		Uf:          uf,
		Estado:      estado,
		Regiao:      regiao,
		Ibge:        ibge,
		Gia:         gia,
		Ddd:         ddd,
		Siafi:       siafi,
	}
}
