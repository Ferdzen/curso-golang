package main

import "fmt"

// Não tem operador ternaário em Go
func obterResultado(nota float64) string {
	//return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "aprovado"
	}
	return "reprovado"

}

func main() {

	fmt.Println(obterResultado(5.0))
}
