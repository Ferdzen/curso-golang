package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Importa o driver PostgreSQL
)

func main() {
	// Configuração da conexão com PostgreSQL
	connStr := "postgres://postgres:root@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Preparar a instrução SQL com placeholders do PostgreSQL ($1, $2, etc.)
	stmt, err := db.Prepare("INSERT INTO usuarios(nome) VALUES($1)")
	if err != nil {
		log.Fatal(err)
	}

	// Executar a instrução preparada
	_, err = stmt.Exec("Pedro")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec("Maria")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec("Caio")
	if err != nil {
		log.Fatal(err)
	}
	// PostgreSQL não suporta LastInsertId()
	// Mostra o ID do ultimo usuario inserido
	var lastInsertId int
	err = db.QueryRow("SELECT LASTVAL()").Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lastInsertId)

	// imprime quantidade de linhas inseridas
	linhas, err := res.RowsAffected()
	println(linhas)
}
