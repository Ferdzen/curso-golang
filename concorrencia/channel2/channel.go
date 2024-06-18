package main

import (
	"fmt"
	"time"
)

// Channel (canal) - forma de comunicação entre goroutines. Channel é um tipo, assim como float e int.

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c // recebendo dados do canal

	fmt.Println(a, b) // imprime os dados recebidos
	fmt.Println(<-c)  // imprime o último resultado
}
