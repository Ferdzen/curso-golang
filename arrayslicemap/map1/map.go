package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[3124659] = "Jonas"
	aprovados[98456231] = "Mario"
	aprovados[46597231] = "Luigi"
	aprovados[45623178] = "Dino"

	fmt.Println(aprovados)

	//percorrendo chave e valor
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[46597231]) // imprime o valor dando a chave
	delete(aprovados, 46597231)      // deleta o valor buscando a chave
	fmt.Println(aprovados[98456231])
	fmt.Println(aprovados[46597231]) // tenta imprimir o valor que foi deletado e nao retona nada

}
