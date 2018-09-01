package painel

import (
	"Projeto_Clinica/back/config"
	"encoding/json"
	"fmt"
	"net/http"
)

//DB ...
var DB = db.Con
var mensagemErro string

// Inicio ...
func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota inicio ...")
	w.Header().Set("Content-Type", "application/json")

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
