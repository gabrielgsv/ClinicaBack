package agendapaciente

//AgendaPaciente ...
type AgendaPaciente struct {
	Codigo         uint   `json:"codigo"`
	NomeMedico     string `json:"nomemedico"`
	Especializacao string `json:"especializacao"`
	Hora           int    `json:"hora"`
	Status         string `json:"status"`
}
