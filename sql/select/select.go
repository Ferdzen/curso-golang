package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Realiza consulta no banco
	rows, err := db.Query("select id, nome from usuarios where id > $1", 5)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Percorrendo variavel que recebeu select
	for rows.Next() {
		var u usuario
		err := rows.Scan(&u.id, &u.nome)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u)
	}

	// Verificar erros após a iteração
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
