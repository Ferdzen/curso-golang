package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}

	var slice1 []int

	//array1 = append(array1, 4, 5, 6) -- Comando Append funciona só com slice
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2) // slice de tamanho 2
	copy(slice2, slice1)     // copia os 2 primeiros elementos do slice1
	fmt.Println(slice2)      // o copy não aumenta o tamanho do slice como o append
}
