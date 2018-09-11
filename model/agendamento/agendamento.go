package agendamento

//Agendamento ...
type Agendamento struct {
	Codigo         uint   `json:"codigo"`
	Codigopaciente int    `json:"codigopaciente"`
	Codigomedico   int    `json:"codigomedico"`
	Data           string `json:"data"`
	HoraInicio     string `json:"horainicio"`
	HoraFim        string `json:"horafim"`
	Motivo         string `json:"motivo"`
	Alergias       string `json:"alergias"`
	Status         string `json:"status"`
}
