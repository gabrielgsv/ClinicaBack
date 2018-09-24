package medico

import (
	"ClinicaBack/config"
	"ClinicaBack/model/agendamedico"
	"ClinicaBack/model/agendamento"
	"ClinicaBack/model/horariosagenda"
	"ClinicaBack/model/medico"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//DB ...
var DB = db.Con
var mensagemErro string

var medicos []medico.Medico
var agendamentos []agendamento.Agendamento
var horarios []horariosagenda.Horariosagenda
var agendamedicos []agendamedico.AgendaMedico

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

// BuscarHorariosDisponiveis ...
// func BuscarHorariosDisponiveis(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Chamando rota buscar horarios disponíveis do médico ...")
// 	w.Header().Set("Content-Type", "application/json")

// 	agendamentos = agendamentos[:0]
// 	novoAgendamento := agendamento.Agendamento{}
// 	agendamento := agendamento.Agendamento{}

// 	err := json.NewDecoder(r.Body).Decode(&novoAgendamento)
// 	mensagemErro = "erro_corpo"
// 	CheckErro(w, r, mensagemErro, err)

// 	query := "SELECT data,hora FROM agendamento WHERE codigomedico = ? AND data = ? AND hora BETWEEN ? AND ?"
// 	rows, err := DB.Query(query, novoAgendamento.Codigomedico, novoAgendamento.Data, novoAgendamento.HoraInicio, novoAgendamento.HoraFim)
// 	mensagemErro = "query_exec_erro"
// 	CheckErro(w, r, mensagemErro, err)

// 	for rows.Next() {
// 		rows.Scan(&agendamento.Data, &agendamento.HoraFim)
// 		agendamentos = append(agendamentos, agendamento)
// 	}

// 	if len(agendamentos) > 0 {
// 		w.WriteHeader(400)
// 	} else {
// 		w.WriteHeader(http.StatusOK)
// 	}
// }

// BuscarHorariosDisponiveis ...
func BuscarHorariosDisponiveis(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota buscar horarios disponíveis do médico ...")
	w.Header().Set("Content-Type", "application/json")

	horarios = horarios[:0]
	horario := horariosagenda.Horariosagenda{}

	data := mux.Vars(r)["data"]
	codigomedico := mux.Vars(r)["codigomedico"]

	query := "SELECT codigo,hora FROM agendamento WHERE codigomedico = ? AND data = ? AND status != 'a'"
	rows, err := DB.Query(query, codigomedico, data)
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	if rows != nil {
		for rows.Next() {
			rows.Scan(&horario.Codigo, &horario.Horario)
			horarios = append(horarios, horario)
		}

		jsonResult, err := json.Marshal(horarios)
		if err != nil {
			fmt.Fprintln(w, "Erro ao gerar o json.")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResult)
	}
}

func BuscarHorariosIndisponiveis(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota buscar horarios indisponíveis do médico ...")
	w.Header().Set("Content-Type", "application/json")

	horarios = horarios[:0]
	horario := horariosagenda.Horariosagenda{}

	data := mux.Vars(r)["data"]
	codigomedico := mux.Vars(r)["codigomedico"]

	query := "SELECT codigo,hora FROM agendamento WHERE codigomedico = ? AND data = ? AND status = 'a'"
	rows, err := DB.Query(query, codigomedico, data)
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	if rows != nil {
		for rows.Next() {
			rows.Scan(&horario.Codigo, &horario.Horario)
			horarios = append(horarios, horario)
		}

		jsonResult, err := json.Marshal(horarios)
		if err != nil {
			fmt.Fprintln(w, "Erro ao gerar o json.")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResult)
	}
}

// Agenda ...
func Agenda(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota agenda medico ...")
	w.Header().Set("Content-Type", "application/json")

	data := mux.Vars(r)["data"]
	codigomedico := mux.Vars(r)["codigomedico"]

	agendas := agendamedicos[:0]

	query := "SELECT agendamento.codigo,nome,email,carteira FROM agendamento " +
		"INNER JOIN paciente " +
		"ON paciente.codigo = agendamento.codigopaciente " +
		"WHERE codigomedico = ? AND data = ?"
	rows, err := DB.Query(query, codigomedico, data)
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	for rows.Next() {
		agenda := agendamedico.AgendaMedico{}
		rows.Scan(&agenda.Codigo, &agenda.NomePaciente, &agenda.Email, &agenda.Carteira)
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
	codigomedico := mux.Vars(r)["codigomedico"]

	agendas := medico.Medico{}

	query := "SELECT COUNT(*) FROM agendamento WHERE codigomedico = ? AND data = ?"

	row := DB.QueryRow(query, codigomedico, dataAtual)

	row.Scan(&agendas.AgendamentosDia)

	json.NewEncoder(w).Encode(agendas)
}

// TotalAgendamento ...
func TotalAgendamento(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota total agendas ...")
	w.Header().Set("Content-Type", "application/json")

	codigomedico := mux.Vars(r)["codigomedico"]

	agendas := medico.Medico{}

	query := "SELECT COUNT(*) FROM agendamento WHERE codigomedico = ?"

	row := DB.QueryRow(query, codigomedico)

	row.Scan(&agendas.TotalAgedamentos)

	json.NewEncoder(w).Encode(agendas)
}

// AgendaHome ...
func AgendaHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota agenda home médico ...")
	w.Header().Set("Content-Type", "application/json")

	currentTime := time.Now()
	dataAtual := currentTime.Format("2006-01-02")
	codigomedico := mux.Vars(r)["codigomedico"]

	agendas := agendamedicos[:0]

	query := "SELECT agendamento.codigo,nome,email,carteira FROM agendamento " +
		"INNER JOIN paciente " +
		"ON paciente.codigo = agendamento.codigopaciente " +
		"WHERE codigomedico = ? AND data = ?"
	rows, err := DB.Query(query, codigomedico, dataAtual)
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	for rows.Next() {
		agenda := agendamedico.AgendaMedico{}
		rows.Scan(&agenda.Codigo, &agenda.NomePaciente, &agenda.Email, &agenda.Carteira)
		agendas = append(agendas, agenda)
	}

	json.NewEncoder(w).Encode(agendas)

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
