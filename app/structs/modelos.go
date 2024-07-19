package structs

type Usuario struct {
	Nome      string     `json:"nome"`
	Idade     int        `json:"idade"`
	Enderecos []Endereco `json:"Enderecos"`
}

type Endereco struct {
	NomeEndereco string `json:"nomeEndereco"`
}

type ProximaEtapaPayload struct {
	IdDynamo string `json:"idDynamo"`
	Nome     string `json:"nome"`
	Sucesso  bool   `json:"sucesso"`
}
