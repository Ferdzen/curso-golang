package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Importa o driver PostgreSQL
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("insert into usuarios(nome) values($1)")

	stmt.Exec("Ana Vitoria")
	stmt.Exec("Nicolas")
	_, err = stmt.Exec("Carlos Henrique")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

}
