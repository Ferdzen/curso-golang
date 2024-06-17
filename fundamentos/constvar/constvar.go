package main

import (
	"fmt"    //import voltado a entradas e saídas formatadas
	m "math" //renomeia import caso precise
)

func main() {

	const PI float64 = 3.1415
	var raio = 3.2 //forma inferida pelo compilador (float64)

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)

	fmt.Println("A área da circuferência é:", area)

	const ( // definindo variaveis junto
		a = 1
		b = 2
	)

	var ( // definindo variaveis junto
		c = 3
		d = 4
	)

	println(a, b, c, d)

	g, h, i := 2, false, "uepa!"

	println(g, h, i)
}
