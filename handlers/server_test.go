package handlers

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

type Usuario struct{
	Results []struct {
		CPF string `json:"cpf"`
		Nome string `json:"nome"`
		Endereco string `json:"endereco"`
		Telefone int64 `json:"telefone"`
	} `json:"results"`
	Status string `json:"status"`
}

// TestCadastroUsuario testa a função de criar usuário
func TestCadastroUsuario(t *testing.T) {
	
	resp, err := http.Post("http://localhost:5000/usuarios",
							 "application/json", 
							 bytes.NewBuffer([]byte(`{"cpf":"307.037.130-27","nome":"Testinho","endereco":"Endereço Para Teste","telefone":986263622}`)))

	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Erro ao ler o corpo da resposta: %v", err)
	}

	if string(body) != "Usuário criado com sucesso! ID: 0"{
		t.Errorf("Erro ao criar usuário: %v", string(body))
	}
}

// TestBuscarUsuarios testa a função de buscar usuários
func TestBuscarUsuarios(t *testing.T) {

	resp, err := http.Get("http://localhost:5000/usuarios")

	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}

	if resp.StatusCode != http.StatusOK{
		t.Errorf("Erro ao buscar usuários: %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Erro ao ler o corpo da resposta: %v", err)
	}

	log.Printf("%s", body)


}

// TestBuscarUsuario testa a função de buscar usuário pelo cpf
func TestBuscarUsuario(t *testing.T) {
	resp, err := http.Get("http://localhost:5000/usuarios/123456789")
	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Erro ao ler o corpo da resposta: %v", err)
	}
	log.Println(string(body))
	
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Erro ao buscar usuário: %v", resp.StatusCode)
	}

}

// TestAtualizarUsuario testa a função de atualizar usuário
func TestAtualizarUsuario(t *testing.T){
	req, err := http.NewRequest("PUT",
								"http://localhost:5000/usuarios/123456789",
								bytes.NewBuffer([]byte(`{"cpf":"123456789","nome":"Testinho","endereco":"Endereço Para Teste","telefone":986263622}`)))

	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	defer req.Body.Close()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Erro ao atualizar usuário: %v", resp.StatusCode)
	}

}

// TestDeletarUsuario testa a função de deletar usuário
func TestDeletarUsuario(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:5000/usuarios/123456789", nil)
	if err != nil {
		t.Errorf("Erro ao criar requisição: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("Erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Erro ao deletar usuário: %v", resp.StatusCode)
	}
}
