package main

import (
	"cadastro/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", handlers.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", handlers.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{cpf}", handlers.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{cpf}", handlers.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{cpf}", handlers.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}