package agendamento

//Agendamento ...
type Agendamento struct {
	Codigo         uint    `json:"codigo"`
	Codigopaciente int     `json:"codigopaciente"`
	Codigomedico   int     `json:"codigomedico"`
	Data           string  `json:"data"`
	HoraInicio     int     `json:"horainicio"`
	HoraFim        float64 `json:"horafim"`
	Motivo         string  `json:"motivo"`
	Alergias       string  `json:"alergias"`
	Status         string  `json:"status"`
}
