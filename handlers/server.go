package handlers

import (
	"cadastro/database"
	"cadastro/validators"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID            int    `json:"id"`
	CPF           string `json:"cpf"`
	Nome          string `json:"nome"`
	Endereco      string `json:"endereco"`
	Telefone      string `json:"telefone"`
	DataNascimeto string `json:"dataNascimento"`
}

// função para criar ususários
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// lendo o cortpo da requisição
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	// convertendo json para um objeto
	var usuario usuario
	if err = json.Unmarshal(requestBody, &usuario); err != nil {
		w.Write([]byte("Falha ao converter o JSON para um objeto!"))
		return
	}

	ver, _ := validators.VerificarCPF(usuario.CPF)
	if !ver {
		w.Write([]byte("CPF inválido!"))
	}

	// abrindo conexão com o banco de dados
	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("insert into usuarios(id, cpf, nome, endereco, telefone, dataNascimento) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement! " + err.Error()))
		return
	}
	defer statement.Close()

	insercao, err := statement.Exec(usuario.ID, usuario.CPF, usuario.Nome, usuario.Endereco, usuario.Telefone, usuario.DataNascimeto)
	if err != nil {
		w.Write([]byte("Erro ao inserir o usuário!" + err.Error()))
		return
	}

	idInserido, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário criado com sucesso! ID: %d", idInserido))) // pesquisar idSerido do sql.Result

}

// BuscarUsuarios retorna todos os usuários cadastrados no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	// abrindo conexão com o banco de dados
	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	// lendo as linhas da tabema usuario no banco de dados
	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Erro ao executar o select!"))
		return
	}
	defer linhas.Close()

	// convertendo o corpo da requisição para um slice de usuários
	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if err := linhas.Scan(&usuario.ID, &usuario.CPF, &usuario.Nome, &usuario.Endereco, &usuario.Telefone, &usuario.DataNascimeto); err != nil {
			w.Write([]byte("Erro ao escanera usuario!"))
			return
		}

		usuarios = append(usuarios, usuario)

	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao converter o JSON para um objeto!"))
		return
	}

}

// BuscarUsuarioPorCpf retorna um usuário pelo CPF
func BuscarUsuarioPorCpf(w http.ResponseWriter, r *http.Request) {

	// verificando parâmetro da requisição
	parametros := mux.Vars(r)

	// abrindo conexão com o banco de dados
	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	linha, err := db.Query("SELECT * FROM usuarios WHERE cpf = ?", parametros["cpf"])
	if err != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}

	var usr usuario
	if linha.Next() {
		if err := linha.Scan(&usr.ID, &usr.CPF, &usr.Nome, &usr.Endereco, &usr.Telefone, &usr.DataNascimeto); err != nil {
			w.Write([]byte("Erro ao escanear usuario!"))
			return
		}
	}

	if usr.CPF == " " {
		w.WriteHeader(http.StatusNotFound)
	}

	if err := json.NewEncoder(w).Encode(usr); err != nil {
		w.Write([]byte("Erro ao converter o JSON para um objeto!"))
		return
	}

}

// BuscarUsuarioPorID retorna um usuário pelo ID
func BuscarUsuarioPorID(w http.ResponseWriter, r *http.Request) {
	// verificando parâmetro da requisição
	parametros := mux.Vars(r)
	ID, err := strconv.ParseInt(parametros["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Erro ao converter o ID para int!"))
		return
	}

	// abrindo conexão com o banco de dados
	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	linha, err := db.Query("SELECT * FROM usuarios WHERE id = ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}

	var usr usuario
	if linha.Next() {
		if err := linha.Scan(&usr.ID, &usr.CPF, &usr.Nome, &usr.Endereco, &usr.Telefone, &usr.DataNascimeto); err != nil {
			w.Write([]byte("Erro ao escanear usuario!"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usr); err != nil {
		w.Write([]byte("Erro ao converter o JSON para um objeto!"))
		return
	}
}

// AtualizarUsuario altera um usuário pelo CPF
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	// verificando parametro
	parametros := mux.Vars(r)

	CPF, err := strconv.ParseInt(parametros["cpf"], 10, 64)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro!" + err.Error()))
		return
	}

	// lendo o corpo da requisição
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usr usuario
	if err := json.Unmarshal(bodyRequest, &usr); err != nil {
		w.Write([]byte("Falha ao converter o JSON para um objeto!"))
		return
	}

	// abrindo conexão com o banco de dados
	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update usuarios set id = ?, nome = ?, endereco = ?, telefone = ?, dataNascimento = ? where cpf = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement! " + err.Error()))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(usr.ID, usr.Nome, usr.Endereco, usr.Telefone, usr.DataNascimeto, CPF); err != nil {
		w.Write([]byte("Erro ao atualizar o usuário!" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// DeletarUsuario deleta um usuário pelo CPF
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	// verificando parâmetro da requisição
	parametros := mux.Vars(r)
	ver, CPF := validators.VerificarCPF(parametros["cpf"])
	if !ver {
		w.Write([]byte("CPF inválido!"))
		return
	}

	// abrindo conexão com o banco de dados
	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where cpf = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement! " + err.Error()))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(CPF); err != nil {
		w.Write([]byte("Erro ao deletar o usuário!" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
