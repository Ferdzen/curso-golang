package main

import "fmt"

func main() {

	/*
		i := 1
		for i <= 10 { // não tem while em Go
			fmt.Printf("%d", i)
		}
	*/

	fmt.Println("\nMisturando....")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Par")
		} else {
			fmt.Printf("Impar")
		}
	}

}
