package main

import (
	"fmt"
	"time"
)

func main() { //switch executa case que estiver TRUE
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() > 12:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}
}
