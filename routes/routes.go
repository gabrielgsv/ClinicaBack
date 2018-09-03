package routes

import (
	"ClinicaBack/config"
	"ClinicaBack/controller/auth"
	"ClinicaBack/controller/medicos"
	"ClinicaBack/controller/pacientes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	// "github.com/gorilla/handlers"
	"github.com/rs/cors"
)

var portaAplicacao string

// HandleFunc ...
func HandleFunc() {
	rotas := mux.NewRouter().StrictSlash(true)

	db.TestarConn()

	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "Bearer", "content-type", "Origin", "Accept", "X-Requested-With", "Authorization"},
		Debug:            true,
	})

	handler := c.Handler(rotas)

	portaAplicacao = ":" + os.Getenv("PORT")

	fmt.Println("Aplicação ON: porta => ", portaAplicacao)

	rotas.HandleFunc("/api/login", auth.Logar).Methods("POST")
	rotas.HandleFunc("/api/recuperartoken", auth.RecuperarToken).Methods("GET")
	rotas.HandleFunc("/api/validartoken", auth.ValidarToken).Methods("GET")

	rotas.HandleFunc("/api/paciente", pacientes.Adicionar).Methods("POST")
	rotas.HandleFunc("/api/paciente", pacientes.Todos).Methods("GET")
	rotas.HandleFunc("/api/alterarpaciente", pacientes.Alterar).Methods("POST")
	rotas.HandleFunc("/api/buscarpaciente", pacientes.Buscar).Methods("POST")

	rotas.HandleFunc("/api/medico", medico.Adicionar).Methods("POST")
	rotas.HandleFunc("/api/medico", medico.Todos).Methods("GET")
	rotas.HandleFunc("/api/alterarmedico", medico.Alterar).Methods("POST")
	rotas.HandleFunc("/api/buscarmedico", medico.Buscar).Methods("POST")
	rotas.HandleFunc("/api/especializacao", medico.BuscarEspecializacao).Methods("POST")

	log.Fatal(http.ListenAndServe(portaAplicacao, handler))
}
