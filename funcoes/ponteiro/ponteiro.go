package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja, ter acesso ao valor no qual o ponteiro aponta
	*n++
} // deixa de ser uma função pura, isto é, altera valores que estão fora do seu escopo. É ruim, melhor evitar. Pode ocasionar bugs.

func main() {
	n := 1

	inc1(n) // passagem por valor

	fmt.Println(n)

	// revisão: & é usado para obter o endereço da variável
	inc2(&n) // por referência

	fmt.Println(n)
}
