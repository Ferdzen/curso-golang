package main

import "fmt"

func main() {
	s := make([]int, 10)

	s[9] = 12

	fmt.Println(s)

	// make(tipo, tamanho, capacidade máx) = cria o slice com tamanho e capacidade máxima
	s = make([]int, 10, 20)

	// len() = tamanho do slice
	// cap() = capacidade máxima do slice
	fmt.Println(s, len(s), cap(s))

	// slice em capacidade máxima
	//s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	//fmt.Println(s, len(s), cap(s))

	// quando adicionado elemento em slice já cheio, ele se auto ajusta para receber o elemento
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
