package database

import (
	"cadastro/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o mysql
)

// Conectar abre a conexão com o banco de dados e a retorna
func Conectar() (*sql.DB, error){
	db, err := sql.Open("mysql", config.StringConexaoBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
	
}