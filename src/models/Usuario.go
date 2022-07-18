package models

import (
	"cadastro/src/models/validators"
	"errors"
	"strings"
)

// Usuario é um modelo para um usuário
type Usuario struct {
	ID             uint64 `json:"id,omitempty"`
	Nome           string `json:"nome,omitempty"`
	CPF            string `json:"cpf,omitempty"`
	Endereco       string `json:"endereco,omitempty"`
	Telefone       string `json:"telefone,omitempty"`
	DataNascimento string `json:"dataNascimento,omitempty"`
}

// Preparar irá chamar os métodos para validar e formatar o nome do usuário recebido
func (usr *Usuario) Preparar() error{
	if err := usr.Validar(); err != nil{
		return err
	}

	usr.formatar()
	return nil

}

// Validar irá validar os dados do usuário
func (usr *Usuario) Validar() error{

	if usr.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	validarCPF, _:= validators.VerificarCPF(usr.CPF)

	if !validarCPF {
		return errors.New("cpf inválido")
	}

	_, validarTelefone := validators.NumeroValido(usr.Telefone)

	if !validarTelefone {
		return errors.New("telefone inválido")
	}

	return nil
}

// Formatar irá formatar o nome do usuário para remover espaços extras criados no nome
func (usr *Usuario) formatar(){
	usr.Nome = strings.TrimSpace(usr.Nome)
}