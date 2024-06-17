package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":  1500.78,
			"Gustavo Pereira": 846.78,
		},
		"J": {
			"José João": 11625.45,
		},
		"P": {
			"Pedro Junior": 1800.00,
		},
	}

	delete(funcsPorLetra, "J")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
