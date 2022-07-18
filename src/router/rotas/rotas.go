package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Função Rota representa todas as rotas a serem inicializadas
type Rota struct {
	URI             string
	Metodo          string
	Funcao          func(http.ResponseWriter, *http.Request)
	ReqAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router{
	rotas := rotasUsuarios

	for _, rota := range rotas{
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
