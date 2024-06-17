package main

import "fmt"

func init() {
	//por convenção, mesmo não sendo chamado, a função init vai ser executada antes do Main. Se houver funções init em outros arquivos do mesmo pacote, todas elas vão executar antes do main.
	fmt.Println("Inicializando....")
}

func main() {
	fmt.Println("Main.....")
}
