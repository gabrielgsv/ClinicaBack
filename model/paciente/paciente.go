package paciente

//Paciente ...
type Paciente struct {
	Codigo         uint   `json:"codigo"`
	Nome           string `json:"nome"`
	Email          string `json:"email"`
	Senha          string `json:"senha"`
	DataNascimento string `json:"data_nascimento"`
	Hospital       string `json:"hospital"`
	Carteira       string `json:"carteira"`
	Role           string `json:"role"`
}
