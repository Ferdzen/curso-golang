package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1

	return // retorno limpo, como há dois tipos de retorno e eles são nomeados, não há necessidade de escrever o que será retornado
}

func main() {
	x, y := trocar(2, 3)

	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
