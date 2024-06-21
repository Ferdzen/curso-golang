package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Importa o driver PostgreSQL
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}
	return result
}

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:root@localhost/postgres?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verifica se a conexão está disponível
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão estabelecida com sucesso!")

	// Execução das operações SQL
	exec(db, `create database cursogo`)
	exec(db, `drop table if exists usuarios`)
	exec(db, `create table usuarios(
                id serial primary key,
                nome varchar(80)
    )`)

	fmt.Println("Tabelas criadas com sucesso!")
}
