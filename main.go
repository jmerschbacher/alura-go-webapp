package main

import (
	"alura-go-webapp/routes"
	"net/http"
)

func main() {
	routes.CarregarRotas()
	// Ouve requisicoes a partir da porta 8000 (sobe um server)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err.Error())
	}
}
