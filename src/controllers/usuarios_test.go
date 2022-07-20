package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
)

var user_Id = 3

type Usuario struct {
	ID             uint64 `json:"id,omitempty"`
	Nome           string `json:"nome,omitempty"`
	CPF            string `json:"cpf,omitempty"`
	Endereco       string `json:"endereco,omitempty"`
	Telefone       string `json:"telefone,omitempty"`
	DataNascimento string `json:"dataNascimento,omitempty"`
}

//   TestCadastrarUsuario testa o cadastro de um usuário
func TestCadastrarUsuario(t *testing.T) {

	resp, err := http.Post("http://localhost:5000/usuarios", "application/json",
		bytes.NewBuffer([]byte(`{"nome":"Fulano testado","cpf":"713.181.860-72","endereco":"Rua Qualquer de Teste", 
										"telefone":"(11)9 4002-8922","dataNascimento":"10/05/2010"}`)))

	if err != nil {
		t.Errorf("Falha ao criar a requisição: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Erro ao ler o corpo da resposta %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		log.Println(string(body))
		t.Error("O objeto não foi criado")
	}

}

// TestCadastrarUsuarioCpfVazio testa a criação de um usuário com cpf vazio
func TestCadastrarUsuarioCpfVazio(t *testing.T) {

}

// TestCadastrarUsuarioCpfVazio testa a criação de um usuário com cpf invalido
func TestCadastrarUsuarioCpfInvalido(t *testing.T) {

}

// TestBuscarUsuarios testa a busca de todos os usuários
func TestBuscarUsuarios(t *testing.T) {
	resp, err := http.Get("http://localhost:5000/usuarios")
	if err != nil {
		t.Errorf("Falha ao criar a requisição: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Falha ao ler o corpo da resposta: %v", err)
	}

	usr := []Usuario{}
	err = json.Unmarshal(body, &usr)
	if err != nil {
		t.Errorf("Falha ao converter o corpo da resposta: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Erro ao buscar usuários: %v", resp.StatusCode)
	}

	// verificando id do primeiro usuário
	if usr[0].ID != 1 {
		t.Errorf("Usuário com id diferente: %v", usr[0].ID)
	}

	// verificando o nome do primeiro usuário
	if usr[0].Nome != "Bob o Bobo" {
		t.Errorf("Usuário com nome diferente: %v", usr[0].Nome)
	}

	// verificando o cpf do primeiro usuário
	if usr[0].CPF != "789.560.430-94" {
		t.Errorf("Usuário com cpf diferente: %v", usr[0].CPF)
	}

	// verificando id do segundo usuário
	if usr[1].ID != 3 {
		t.Errorf("Usuário com id diferente: %v", usr[1].ID)
	}

	// verificando o nome do segundo usuário
	if usr[1].Nome != "Fulano tentado" {
		t.Errorf("Usuário com nome diferente: %v", usr[1].Nome)
	}

	// verificando o cpf do segundo usuário
	if usr[1].CPF != "713.181.860-72" {
		t.Errorf("Usuário com cpf diferente: %v", usr[1].CPF)
	}

	log.Println(usr)

}

// TestBuscarUsuario testa a pesquisa de um usuário passando o id como parâmetro
func TestBuscarUsuarioPorId(t *testing.T) {

	resp, err := http.Get("http://localhost:5000/usuarios/3")
	if err != nil {
		t.Errorf("Falha ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Erro ao buscar usuário: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	if err != nil {
		t.Errorf("Falha ao ler o corpo da resposta: %v", err)
	}

	usr := Usuario{}
	err = json.Unmarshal(body, &usr)
	if err != nil {
		t.Errorf("Falha ao converter o json para struct: %v", err)
	}

	nomeEsperado := "Fulano tentado"
	cpfEsperado := "713.181.860-72"
	// verificando se o nome do usuário buscado é o mesmo do esperado
	if usr.Nome != nomeEsperado {
		t.Errorf("Usuário com nome diferente: %v, esperado %v", usr.Nome, nomeEsperado)
	}

	// verificando se o cpf do usuário buscado é o mesmo do esperado
	if usr.CPF != cpfEsperado {
		t.Errorf("Usuário com cpf diferente: %v, esperado %v", usr.CPF, cpfEsperado)
	}

}

// TestBuscarUsuarioPorNome testa a pesquisa de um usuário passando o id inexistente como parâmetro
func TestBuscarUsuarioPorIdInexistente(t *testing.T) {

	resp, err := http.Get("http://localhost:5000/usuarios/2")
	if err != nil {
		t.Errorf("Falha ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Erro ao buscar usuário: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	if err != nil {
		t.Errorf("Falha ao ler o corpo da resposta: %v", err)
	}

	usr := Usuario{}
	err = json.Unmarshal(body, &usr)
	if err != nil {
		t.Errorf("Falha ao converter o json para struct: %v", err)
	}

	// comparando os atributos do objeto buscado
	if usr.Nome != "" {
		t.Errorf("Usuário com nome diferente: %v", usr.Nome)
	}

	if usr.CPF != "" {
		t.Errorf("Usuário com cpf diferente: %v", usr.CPF)
	}

	if usr.Endereco != "" {
		t.Errorf("Usuário com endereço diferente: %v", usr.Endereco)
	}

	if usr.Telefone != "" {
		t.Errorf("Usuário com telefone diferente: %v", usr.Telefone)
	}

	if usr.DataNascimento != "" {
		t.Errorf("Usuário com data de nascimento diferente: %v", usr.DataNascimento)
	}

}

// TestExcluirUsuario irá excluir um usuário com o id inserido
func TestExcluirUsuario(t *testing.T) {

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:5000/usuarios/"+strconv.Itoa(user_Id), nil)
	if err != nil {
		t.Error("**********************************\n")
		t.Errorf("Erro ao fazer a requisição %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error("**********************************\n")
		t.Errorf("Erro ao obter a resposta %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("**********************************\n")
		t.Errorf("Erro ao ler o corpo da resposta %v", err)
	}

	if resp.StatusCode == http.StatusNoContent {
		t.Errorf("Status diferente do esperado")
	}

	if string(body) != "{\"status\":\"success\"}" {
		t.Error(string(body))
		t.Error("Erro ao deletar usuário")
	}
}
