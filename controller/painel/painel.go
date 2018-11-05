package painel

import (
	"ClinicaBack/config"
	"ClinicaBack/model/agendamento"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//DB ...
var DB = db.Con
var mensagemErro string

var agendamentos []agendamento.Agendamento

// CancelarConsulta ...
func CancelarConsulta(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota cancelar consulta ...")
	w.Header().Set("Content-Type", "application/json")

	codigoagendamento := mux.Vars(r)["codigoagendamento"]

	stmt, err := DB.Prepare("UPDATE agendamento SET status = 'c' WHERE codigo = ?")
	mensagemErro = "query_exec_erro"
	CheckErro(w, r, mensagemErro, err)

	stmt.Exec(codigoagendamento)
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
