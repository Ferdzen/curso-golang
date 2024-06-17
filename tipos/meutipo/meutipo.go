package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	/*
		if n.entre(9.0, 10.0) {
			return "A"
		} else if n.entre(7.5, 8.99) {
			return "B"
		} else if n.entre(6.0, 7.99) {
			return "C"
		} else {
			return "D"
		}*/

	switch { // alternativa mais simples e "limpa", ao invés de usar if/elseif
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.5, 8.99):
		return "B"
	case n.entre(6.0, 7.5):
		return "C"
	case n.entre(3.5, 6.0):
		return "D"
	}

	return "E"
}

func main() {
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(7.7))
	fmt.Println(notaParaConceito(6.5))
	fmt.Println(notaParaConceito(5.1))
	fmt.Println(notaParaConceito(2.0))
}
