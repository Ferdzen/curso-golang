package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha")
	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.141516
	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f", x)

}
