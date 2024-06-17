package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3
	i += 3
	i -= 3
	i /= 2
	i *= 2
	i %= 2
	fmt.Println("Teste: ", i)

	x, y := 1, 2

	fmt.Println("Teste 2 - atribuindo valores para duas variaveis ao mesmo tempo: ", x, y)

	x, y = y, x

	fmt.Println("Teste 3 - inversão de valores: ", x, y)
}
