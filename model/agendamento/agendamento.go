package agendamento

//Agendamento ...
type Agendamento struct {
	Codigo         uint   `json:"codigo"`
	Codigopaciente int    `json:"codigopaciente"`
	Codigomedico   int    `json:"codigomedico"`
	Data           string `json:"data"`
	Horario        string `json:"horario"`
	Motivo         string `json:"motivo"`
	Alergias       string `json:"alergias"`
	Status         string `json:"status"`
}
