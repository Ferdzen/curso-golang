package main

import (
	"fmt"
	"time"
)

// A variável do tipo interface é uma variável versátil, isto é, uma variável não tipada que pode receber valores de qualquer tipo.

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "String"
	case func():
		return "função"
	default:
		return "Não sei"

	}
}

func main() {
	fmt.Println(tipo(2.5))
	fmt.Println(tipo(1))
	fmt.Println(tipo("top"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
