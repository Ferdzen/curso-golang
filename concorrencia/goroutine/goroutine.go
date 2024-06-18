package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// Forma serial de chamar função
	//fale("Maria", "Pq vc nao fala comigo?", 3)
	//fale("João", "Só posso falar depois de vc!", 1)

	// Goroutine será executada no mesmo tempo de execução da função Main. Neste exemplo não foi executado nenhuma vez as funções pois o método main termina de executar antes de 1 segundo.
	//go fale("Maria", "Ei....", 500)
	//go fale("João", "Opa....", 500)

	// Mas se pedir para esperar um pouquinho...
	//time.Sleep(time.Second * 7)

	// Função chamada para ser executada de forma independente.
	go fale("Maria", "Entendi!!!", 10)

	// Função chamada normalmente.
	fale("João", "Parabéns!", 5)
}
