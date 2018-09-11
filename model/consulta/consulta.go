package consulta

//Consulta ...
type Consulta struct {
	Codigo            uint   `json:"codigo"`
	Codigoagendamento int    `json:"codigoagendamento"`
	descricao         string `json:"descricao"`
	exames            string `json:"exames"`
	remedios          string `json:"remedios"`
	data              string `json:"data"`
	hora              string `json:"hora"`
	status            string `json:"status"`
}
