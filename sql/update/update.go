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

	stmt, _ := db.Prepare("update usuarios set nome = $1 where id = $2")

	stmt.Exec("Alan Ferreira", 1)
	stmt.Exec("Maethe Perroni", 3)
	stmt.Exec("Silvio Santos", 4)
}
