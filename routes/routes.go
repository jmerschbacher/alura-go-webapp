package routes

import (
	"alura-go-webapp/controllers"
	"net/http"
)

func CarregarRotas() {
	// Mapeia uma rota a uma determinada funcao do controller
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
