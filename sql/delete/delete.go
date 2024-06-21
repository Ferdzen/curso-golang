package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, _ := db.Prepare("delete from usuarios where id = $1 ")

	stmt.Exec(18)
	stmt.Exec(19)
	stmt.Exec(20)
}
