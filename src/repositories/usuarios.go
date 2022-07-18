package repositories

import (
	"cadastro/src/models"
	"database/sql"
)

// usuarios representa um repositório de usuarios
type usuarios struct{
	db *sql.DB
}



// NovoRepositorioDeUsuarios cria um novo repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios{
	return &usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio usuarios) Criar(usuario models.Usuario) (uint64, error){
	statement, err := repositorio.db.Prepare("insert into usuarios (id, nome, cpf, endereco, telefone, dataNascimento) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.ID, usuario.Nome, usuario.CPF, usuario.Endereco, usuario.Telefone, usuario.DataNascimento)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// ListarTodos lista todos os usuários salvos no banco de dados
func (repositorio usuarios) ListarTodos() ([]models.Usuario, error){

	linhas, err := repositorio.db.Query("select * from usuarios")
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usrs []models.Usuario

	for linhas.Next(){
		var usuario	models.Usuario

		if err = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.CPF, &usuario.Endereco, &usuario.Telefone, &usuario.DataNascimento); err != nil {
			return nil, err
		}

		usrs = append(usrs, usuario)
	}

	return usrs, nil
}

// lista o usuário pelo cpf
func (repositorio usuarios) BuscarPorCPF(cpf string) (models.Usuario, error){
	linhas, err := repositorio.db.Query("select * from usuarios where cpf = ?", cpf)
	if err != nil {
		return models.Usuario{}, err
	}
	defer linhas.Close()

	var usr models.Usuario

	if linhas.Next(){
		if err := linhas.Scan(&usr.ID, &usr.Nome, &usr.CPF, &usr.Endereco, &usr.Telefone, &usr.DataNascimento); err != nil {
			return models.Usuario{}, err
		}
	}

	return usr, nil
}