package medico

import (
	"ClinicaBack/config"
	"ClinicaBack/model/medico"
	"encoding/json"
	"fmt"
	"net/http"
)

//DB ...
var DB = db.Con
var mensagemErro string

var medicos []medico.Medico

// Adicionar ...
func Adicionar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota adicionar médico ...")
	w.Header().Set("Content-Type", "application/json")

	novoMedico := medico.Medico{}

	err := json.NewDecoder(r.Body).Decode(&novoMedico)
	mensagemErro = "erro_corpo"
	CheckErro(w, r, mensagemErro, err)

	query := "INSERT INTO medico (nome, email, senha, data_nascimento, especializacao, hospital, crm, role, ativo) VALUES(?,?,?,?,?,?,?,?,?)"
	stmt, err := DB.Prepare(query)
	mensagemErro = "query_montagem_erro"
	CheckErro(w, r, mensagemErro, err)
	fmt.Println(query)

	_, err = stmt.Exec(novoMedico.Nome, novoMedico.Email, novoMedico.Senha, &novoMedico.DataNascimento, &novoMedico.Especializacao, &novoMedico.Hospital, &novoMedico.Crm, "m", "a")
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	response, _ := json.Marshal(&novoMedico)
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	return
}

// Todos ...
func Todos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota buscar todos médico ...")
	w.Header().Set("Content-Type", "application/json")

	medicos = medicos[:0]

	query := "SELECT codigo,nome,email,data_nascimento,especializacao,hospital,crm, ativo FROM medico"
	rows, err := DB.Query(query)
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	for rows.Next() {
		medico := medico.Medico{}
		rows.Scan(&medico.Codigo, &medico.Nome, &medico.Email, &medico.DataNascimento, &medico.Especializacao,
			&medico.Hospital, &medico.Crm, &medico.Ativo)
		medicos = append(medicos, medico)
	}

	json.NewEncoder(w).Encode(medicos)
}

// Buscar ...
func Buscar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota buscar médico ...")
	w.Header().Set("Content-Type", "application/json")

	medicos = medicos[:0]
	medicoBuscar := medico.Medico{}

	err := json.NewDecoder(r.Body).Decode(&medicoBuscar)
	mensagemErro = "erro_corpo"
	CheckErro(w, r, mensagemErro, err)

	query := "SELECT codigo, UPPER(nome) , UPPER(email) ,data_nascimento, especializacao, hospital, crm, ativo FROM medico WHERE nome LIKE ?"
	rows, err := DB.Query(query, medicoBuscar.Nome)
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	for rows.Next() {
		medico := medico.Medico{}
		rows.Scan(&medico.Codigo, &medico.Nome, &medico.Email, &medico.DataNascimento, &medico.Especializacao, &medico.Hospital,
			&medico.Crm, &medico.Ativo)
		medicos = append(medicos, medico)
	}

	json.NewEncoder(w).Encode(medicos)

}

// BuscarEspecializacao ...
func BuscarEspecializacao(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota buscar por especialização médico ...")
	w.Header().Set("Content-Type", "application/json")

	medicos = medicos[:0]
	medicoBuscar := medico.Medico{}

	err := json.NewDecoder(r.Body).Decode(&medicoBuscar)
	mensagemErro = "erro_corpo"
	CheckErro(w, r, mensagemErro, err)

	query := "SELECT codigo, UPPER(nome) , UPPER(email) ,data_nascimento, especializacao, hospital,crm, ativo FROM medico WHERE especializacao = ?"
	rows, err := DB.Query(query, medicoBuscar.Especializacao)
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	for rows.Next() {
		medico := medico.Medico{}
		rows.Scan(&medico.Codigo, &medico.Nome, &medico.Email, &medico.DataNascimento, &medico.Especializacao, &medico.Hospital,
			&medico.Crm, &medico.Ativo)
		medicos = append(medicos, medico)
	}

	json.NewEncoder(w).Encode(medicos)

}

// Alterar ...
func Alterar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota alterar medico ...")
	w.Header().Set("Content-Type", "application/json")

	medico := medico.Medico{}
	err := json.NewDecoder(r.Body).Decode(&medico)
	mensagemErro = "erro_corpo"
	CheckErro(w, r, mensagemErro, err)

	stmt, err := DB.Prepare("UPDATE medico SET nome = ? , email = ? , especializacao = ? , data_nascimento = ? , hospital = ? , crm = ?, ativo = ? WHERE codigo = ?")
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	stmt.Exec(medico.Nome, medico.Email, medico.Especializacao, medico.DataNascimento, medico.Hospital, medico.Crm, medico.Ativo, medico.Codigo)
	json.NewEncoder(w).Encode(medico)

}

// CheckErro ...
func CheckErro(w http.ResponseWriter, r *http.Request, msg string, err error) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response, _ := json.Marshal(map[string]interface{}{
			msg: err.Error(),
		})
		w.Write(response)
		return
	}
}
