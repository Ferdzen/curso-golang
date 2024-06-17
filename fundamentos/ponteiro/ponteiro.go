package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // pegando o endereço da variável
	*p++   // desreferenciando (pegando o valor e incrementando)

	// Go não tem aritmética de ponteiros
	// p++

	fmt.Println("Valor variável: ", i)
	fmt.Println("Endereço de memória:", p)
}
