package medico

//Medico ...
type Medico struct {
	Codigo           uint   `json:"codigo"`
	Nome             string `json:"nome"`
	Email            string `json:"email"`
	Senha            string `json:"senha"`
	DataNascimento   string `json:"data_nascimento"`
	Especializacao   string `json:"especializacao"`
	Hospital         string `json:"hospital"`
	Crm              string `json:"crm"`
	Role             string `json:"role"`
	Ativo            string `json:"ativo"`
	TotalAgedamentos int    `json:"totalagendamentos"`
	AgendamentosDia  int    `json:"agendamentosdia"`
}
