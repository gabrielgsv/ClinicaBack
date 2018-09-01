package auth

import (
	"Projeto_Clinica/back/config"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/request"

	"github.com/dgrijalva/jwt-go"
)

// Credenciais ...
type Credenciais struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

// UsuarioBuscado ...
type UsuarioBuscado struct {
	Codigo int    `json:"codigo"`
	Nome   string `json:"nome"`
	Email  string `json:"email"`
	Senha  string `json:"senha"`
	Role   string `json:"role"`
}

// Usuario ...
type Usuario struct {
	Codigo int    `json:"codigo"`
	Nome   string `json:"nome"`
	Role   string `json:"role"`
	Token  string `json:"token"`
}

// Claim ...
type Claim struct {
	Usuario Usuario `json:"usuario"`
	jwt.StandardClaims
}

// ResponseToken ...
type ResponseToken struct {
	Token string `json:"token"`
}

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

//DB ...
var DB = db.Con
var mensagemErro string
var credencias Credenciais
var usuario UsuarioBuscado
var usuarioToken Usuario

func init() {
	privateBytes, err := ioutil.ReadFile("../private.rsa")
	if err != nil {
		log.Fatal(err, "Erro ao abrir a chave privada !")
	}

	publicBytes, err := ioutil.ReadFile("../public.rsa.pub")
	if err != nil {
		log.Fatal("Erro ao abrir a chave pública !")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("Não foi possível ter parte da chave privada !")
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("Não foi possível ter parte da chave pública !")
	}

}

// GenerateJWT ...
func GenerateJWT(user Usuario) string {
	claims :=
		Claim{
			Usuario: user,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
				Issuer:    "Token de autenticação",
			},
		}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("Não foi possível confirmar o token ..")
	}
	return result
}

// Logar ...
func Logar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chamando rota logar ...")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	usuarioToken.Token = ""
	usuario := UsuarioBuscado{}

	err := json.NewDecoder(r.Body).Decode(&credencias)
	mensagemErro = "erro_corpo"
	CheckErro(w, r, mensagemErro, err)

	query := "SELECT codigo, nome, email, senha, role FROM paciente WHERE email = ? " +
		"AND " +
		"senha = ? " +
		"UNION " +
		"SELECT codigo, nome, email, senha, role FROM medico WHERE email = ? " +
		"AND " +
		"senha = ? "

	row := DB.QueryRow(query, credencias.Email, credencias.Senha, credencias.Email, credencias.Senha)

	row.Scan(&usuario.Codigo, &usuario.Nome, &usuario.Email, &usuario.Senha, &usuario.Role)

	if usuario.Email == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	usuarioToken.Codigo = usuario.Codigo
	usuarioToken.Nome = usuario.Nome
	usuarioToken.Role = usuario.Role
	token := GenerateJWT(usuarioToken)
	usuarioToken.Token = token
	// result := ResponseToken{token}
	jsonResult, err := json.Marshal(usuarioToken.Token)
	if err != nil {
		fmt.Fprintln(w, "Erro ao gerar o json.")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResult)
}

// ValidarToken ...
func ValidarToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Validando token ...")
	w.Header().Set("Content-Type", "application/json")
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				fmt.Println("O token expirou ..")
				w.WriteHeader(http.StatusUnauthorized)
				response, _ := json.Marshal(0)
				fmt.Println(response)
				w.Write(response)
				return
			case jwt.ValidationErrorSignatureInvalid:
				fmt.Println(w, "Os tokens não se coincidem ..")
				w.WriteHeader(http.StatusUnauthorized)
				return
			default:
				fmt.Println(w, "O token não é válido 1")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
	}
	fmt.Println("Aprovando tooken para ser passado ...")
	if token.Valid {
		fmt.Println("Tooken válido ...")
		w.WriteHeader(http.StatusAccepted)
		return
	} else {
		fmt.Println("Tooken inválido...")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func RecuperarToken(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(usuarioToken)
	w.WriteHeader(http.StatusAccepted)
	w.Write(response)
	return
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
