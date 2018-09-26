package agendamedico

//AgendaMedico ...
type AgendaMedico struct {
	Codigo       uint   `json:"codigo"`
	NomePaciente string `json:"nomepaciente"`
	Email        string `json:"email"`
	Horario      string `json:"hora"`
	Status       string `json:"status"`
}
