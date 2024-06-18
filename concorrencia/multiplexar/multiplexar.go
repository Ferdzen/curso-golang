package main

import (
	"fmt"

	"github.com/cursoGo/libsGo/html"
)

// multiplexar - misturar (mensagens) num canal

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.youtube.com/", "https://www.kabum.com.br/"),
		html.Titulo("https://pkg.go.dev/golang.org/x/tools/gopls/internal/analysis/deprecated", "https://pt.wikipedia.org/wiki/Wikip%C3%A9dia:P%C3%A1gina_principal"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
