package pacientes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"ClinicaBack/config"
	"ClinicaBack/model/agendamento"
	"ClinicaBack/model/agendapaciente"
	"ClinicaBack/model/paciente"
)

//DB ...
var DB = db.Con
var mensagemErro string

var pacientes []paciente.Paciente
var agendamentos []agendamento.Agendamento
var agendapacientes []agendapaciente.AgendaPaciente

// Adicionar ...
func Adicionar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota adicionar paciente ...")
	w.Header().Set("Content-Type", "application/json")

	novoPaciente := paciente.Paciente{}

	err := json.NewDecoder(r.Body).Decode(&novoPaciente)
	mensagemErro = "erro_corpo"
	CheckErro(w, r, mensagemErro, err)

	query := "INSERT INTO paciente (nome, email, senha, data_nascimento, hospital, carteira, role, ativo) VALUES(?,?,?,?,?,?,?,?)"
	stmt, err := DB.Prepare(query)
	mensagemErro = "query_montagem_erro"
	CheckErro(w, r, mensagemErro, err)

	fmt.Println(query)

	_, err = stmt.Exec(novoPaciente.Nome, novoPaciente.Email, novoPaciente.Senha, novoPaciente.DataNascimento, novoPaciente.Hospital, novoPaciente.Carteira, "p", "a")
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	fmt.Println("teste")
	response, _ := json.Marshal(&novoPaciente)
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
	return
}

// Todos ...
func Todos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota buscar todos pacientes ...")
	w.Header().Set("Content-Type", "application/json")

	pacientes = pacientes[:0]

	query := "SELECT codigo, UPPER(nome) , UPPER(email) ,data_nascimento, hospital ,carteira , ativo FROM paciente"
	rows, err := DB.Query(query)
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	for rows.Next() {
		paciente := paciente.Paciente{}
		rows.Scan(&paciente.Codigo, &paciente.Nome, &paciente.Email, &paciente.DataNascimento, &paciente.Hospital,
			&paciente.Carteira, &paciente.Ativo)
		pacientes = append(pacientes, paciente)
	}

	json.NewEncoder(w).Encode(pacientes)

}

// Buscar ...
func Buscar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota buscar paciente ...")
	w.Header().Set("Content-Type", "application/json")

	pacientes = pacientes[:0]
	pacienteBuscar := paciente.Paciente{}

	err := json.NewDecoder(r.Body).Decode(&pacienteBuscar)
	mensagemErro = "erro_corpo"
	CheckErro(w, r, mensagemErro, err)

	query := "SELECT codigo, UPPER(nome) , UPPER(email) ,data_nascimento, hospital,carteira, ativo FROM paciente WHERE nome LIKE ?"
	rows, err := DB.Query(query, pacienteBuscar.Nome)
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	for rows.Next() {
		paciente := paciente.Paciente{}
		rows.Scan(&paciente.Codigo, &paciente.Nome, &paciente.Email, &paciente.DataNascimento, &paciente.Hospital,
			&paciente.Carteira, &paciente.Ativo)
		pacientes = append(pacientes, paciente)
	}

	json.NewEncoder(w).Encode(pacientes)

}

// Alterar ...
func Alterar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota alterar paciente ...")
	w.Header().Set("Content-Type", "application/json")

	paciente := paciente.Paciente{}
	err := json.NewDecoder(r.Body).Decode(&paciente)
	mensagemErro = "erro_corpo"
	CheckErro(w, r, mensagemErro, err)

	// query := "UPDATE paciente SET nome = ? ,email = ?,data_nascimento = ?,hospital = ?,carteira = ? FROM paciente WHERE codigo = ?"
	stmt, err := DB.Prepare("UPDATE paciente SET nome = ? , email = ? , data_nascimento = ? , hospital = ? , carteira = ?, ativo = ? WHERE codigo = ?")
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	stmt.Exec(paciente.Nome, paciente.Email, paciente.DataNascimento, paciente.Hospital, paciente.Carteira, paciente.Ativo, paciente.Codigo)
	json.NewEncoder(w).Encode(paciente)

}

// AdicionarConsulta ...
func AdicionarConsulta(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota adicionar nova consulta ...")
	w.Header().Set("Content-Type", "application/json")

	novoAgendamento := agendamento.Agendamento{}

	err := json.NewDecoder(r.Body).Decode(&novoAgendamento)

	mensagemErro = "erro_corpo"
	CheckErro(w, r, mensagemErro, err)

	fmt.Println(novoAgendamento.Codigomedico)
	fmt.Println(novoAgendamento.Codigopaciente)
	fmt.Println(novoAgendamento.Data)
	fmt.Println(novoAgendamento.Horario)
	fmt.Println(novoAgendamento.Motivo)
	fmt.Println(novoAgendamento.Alergias)

	stmt, err := DB.Prepare("INSERT INTO agendamento (codigopaciente, codigomedico, data, hora, motivo, alergias, status) VALUES(?,?,?,?,?,?,?)")
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	stmt.Exec(novoAgendamento.Codigopaciente, novoAgendamento.Codigomedico, novoAgendamento.Data,
		novoAgendamento.Horario, novoAgendamento.Motivo, novoAgendamento.Alergias, "a")
	json.NewEncoder(w).Encode(novoAgendamento)

}

// Agenda ...
func Agenda(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota agenda paciente ...")
	w.Header().Set("Content-Type", "application/json")

	data := mux.Vars(r)["data"]
	codigopaciente := mux.Vars(r)["codigopaciente"]

	agendas := agendapacientes[:0]

	query := "SELECT agendamento.codigo,nome,especializacao,hora,status FROM agendamento " +
		"INNER JOIN medico " +
		"ON medico.codigo = agendamento.codigomedico " +
		"WHERE codigopaciente = ? AND data = ?"
	rows, err := DB.Query(query, codigopaciente, data)
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	for rows.Next() {
		agenda := agendapaciente.AgendaPaciente{}
		rows.Scan(&agenda.Codigo, &agenda.NomeMedico, &agenda.Especializacao, &agenda.Hora, &agenda.Status)
		agendas = append(agendas, agenda)
	}

	json.NewEncoder(w).Encode(agendas)

}

// TotalAgendasDia ...
func TotalAgendasDia(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota total agendas do dia ...")
	w.Header().Set("Content-Type", "application/json")

	currentTime := time.Now()
	dataAtual := currentTime.Format("2006-01-02")
	codigopaciente := mux.Vars(r)["codigopaciente"]

	agendas := paciente.Paciente{}

	query := "SELECT COUNT(*) FROM agendamento WHERE codigopaciente = ? AND data = ?"

	row := DB.QueryRow(query, codigopaciente, dataAtual)

	row.Scan(&agendas.AgendamentosDia)

	json.NewEncoder(w).Encode(agendas)
}

// TotalAgendamento ...
func TotalAgendamento(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota total agendas ...")
	w.Header().Set("Content-Type", "application/json")

	codigopaciente := mux.Vars(r)["codigopaciente"]

	agendas := paciente.Paciente{}

	query := "SELECT COUNT(*) FROM agendamento WHERE codigopaciente = ?"

	row := DB.QueryRow(query, codigopaciente)

	row.Scan(&agendas.TotalAgedamentos)

	json.NewEncoder(w).Encode(agendas)
}

// AgendaHome ...
func AgendaHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota agenda home paciente ...")
	w.Header().Set("Content-Type", "application/json")

	currentTime := time.Now()
	dataAtual := currentTime.Format("2006-01-02")
	codigopaciente := mux.Vars(r)["codigopaciente"]

	agendas := agendapacientes[:0]

	query := "SELECT agendamento.codigo,nome,especializacao,hora,status FROM agendamento " +
		"INNER JOIN medico " +
		"ON medico.codigo = agendamento.codigomedico " +
		"WHERE codigopaciente = ? AND data = ?"
	rows, err := DB.Query(query, codigopaciente, dataAtual)
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	for rows.Next() {
		agenda := agendapaciente.AgendaPaciente{}
		rows.Scan(&agenda.Codigo, &agenda.NomeMedico, &agenda.Especializacao, &agenda.Hora, &agenda.Status)
		agendas = append(agendas, agenda)
	}

	json.NewEncoder(w).Encode(agendas)

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
