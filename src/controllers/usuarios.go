package controllers

import (
	"cadastro/src/database"
	"cadastro/src/models"
	"cadastro/src/repositories"
	"cadastro/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// CriarUsuario intemedia as requisições enviadas para criar um usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	// ler o corpo da requisição
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	// converter o corpo da requisição para um usuário
	var usr models.Usuario
	if err := json.Unmarshal(requestBody, &usr); err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	// faz as validações
	if err := usr.Preparar(); err != nil{
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	// abre conexão com o banco de dados
	db, err := database.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	// insere o usuário no banco de dados
	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usr.ID, err = repositorio.Criar(usr)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}	

	// retorna a resposta da requisição
	responses.JSON(w, http.StatusCreated, usr)
}

// BuscarUsarios intermedia as requisições enviadas para buscar todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	
	// Abre conexão com o banco de dados
	db, err := database.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// acessa o repositório de usuários para fazer a busca
	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usuarios, err := repositorio.ListarTodos()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	// retorna a lista dos usuários buscados
	responses.JSON(w, http.StatusOK, usuarios)

}

// BuscarUsuario intermedia as requisições enviadas para buscar um usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	
	parametros := mux.Vars(r)

	db, err := database.Conectar()
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuarios(db)
	usr, err := repositorio.BuscarPorCPF(parametros["usuarioCPF"])
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, usr)


}

// AtualizarUsuario intermedia as requisições enviadas para atualizar um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizando um usuário!"))
}

// DeletarUsuario intermedia as requisições enviadas para excluir um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletando um usuário!"))
}