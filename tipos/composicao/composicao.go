package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	//adicione mais métodos se quiser
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo....")
}

func (b bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var b esportivoLuxuoso = bmw7{}

	b.fazerBaliza()
	b.ligarTurbo()

}
