package main

func imprimirResultado(nota float64) {

	if nota > 7.0 {
		println("Aluno(a) aprovado(a).")
	} else {
		println("Aluno(a) reprovado(a).")
	}
}

func main() {
	imprimirResultado(8.0)
	//imprimirResultado(5.0)
	//imprimirResultado(7.0)
}
