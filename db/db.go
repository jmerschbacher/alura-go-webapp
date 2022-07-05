package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func ConexaoDB() *sql.DB {
	conexao := "root:%admin$22@/db_alura_loja"
	db, err := sql.Open("mysql", conexao)

	if err != nil {
		panic(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
