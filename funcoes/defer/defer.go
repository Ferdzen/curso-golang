package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!") // segura a execução dessa linha, mesmo ela estando por primeiro. Executa ela apenas antes do return.

	if numero > 5000 {
		fmt.Println("Valor alto....")
		return 5000
	} else {
		fmt.Println("valor baixo...")

		return numero
	}
}

func main() {
	fmt.Println(obterValorAprovado(12000))
	fmt.Println(obterValorAprovado(3000))
}
