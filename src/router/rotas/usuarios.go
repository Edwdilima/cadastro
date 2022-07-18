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
		URI:             "/usuarios/{usuarioId}", // lista um usuário pelo id
		Metodo:          http.MethodGet,
		Funcao:          controllers.BuscarUsuario,
		ReqAutenticacao: false,
	},
	{
		URI:             "/usuarios/{usuarioId}", // atualiza um usuário pelo id
		Metodo:          http.MethodPut,
		Funcao:          controllers.AtualizarUsuario,
		ReqAutenticacao: false,
	},
	{
		URI:             "/usuarios/{usuarioId}", // deleta um usuário pelo id
		Metodo:          http.MethodDelete,
		Funcao:          controllers.DeletarUsuario,
		ReqAutenticacao: false,
	},
}
