package rotas

import (
	"cadastro/src/controllers"
	"net/http"
)

// rotasUsuarios representam todas as rotas dentro do sistema
var rotasUsuarios = []Rota{
	{
		URI:             "/usuarios",
		Metodo:          http.MethodPost, // cadastra um usuário
		Funcao:          controllers.CriarUsuario,
		ReqAutenticacao: false,
	},
	{
		URI:             "/usuarios",
		Metodo:          http.MethodGet, // lista todos os usuários
		Funcao:          controllers.BuscarUsuarios,
		ReqAutenticacao: false,
	},
	{
		URI:             "/usuarios/{usuarioCPF}", // lista um usuário pelo cpf
		Metodo:          http.MethodGet,
		Funcao:          controllers.BuscarUsuario,
		ReqAutenticacao: false,
	},
	{
		URI:             "/usuarios/{usuarioCPF}", // atualiza um usuário pelo cpf
		Metodo:          http.MethodPut,
		Funcao:          controllers.AtualizarUsuario,
		ReqAutenticacao: false,
	},
	{
		URI:             "/usuarios/{usuarioCPF}", // deleta um usuário pelo cpf
		Metodo:          http.MethodDelete,
		Funcao:          controllers.DeletarUsuario,
		ReqAutenticacao: false,
	},
}
