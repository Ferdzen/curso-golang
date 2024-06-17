package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)

	fmt.Println(resultado)

	resultado2 := fatorial(6)

	fmt.Println(resultado2)

	// uma solução melhor seria... usar inteiros sem sinais. uint!
}
