package controllers

import (
	"alura-go-webapp/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Indica que os arquivos html da pasta template serão responsáveis
// pelas telas da aplicação
var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscarTodosProdutos()

	// Cria um template (pagina HTML), a partir de um nome
	// (no caso, Index) mapeado diretamente num arquivo HTML da
	// pasta "templates" neste projeto
	err := temp.ExecuteTemplate(w, "Index", produtos)

	if err != nil {
		panic(err.Error())
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "New", nil)

	if err != nil {
		panic(err.Error())
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := models.Produto{}
		produto.Nome = r.FormValue("nome")
		produto.Descricao = r.FormValue("descricao")

		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}
		produto.Preco = preco

		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}
		produto.Quantidade = quantidade

		models.CriarProduto(&produto)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.ExcluirProdutoPorId(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	produto := models.BuscarProdutoPorId(idProduto)

	err := temp.ExecuteTemplate(w, "Edit", produto)

	if err != nil {
		panic(err.Error())
	}

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := models.Produto{}
		produto.Nome = r.FormValue("nome")
		produto.Descricao = r.FormValue("descricao")

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("Erro na conversão do id", err)
		}
		produto.Id = id

		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}
		produto.Preco = preco

		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}
		produto.Quantidade = quantidade

		models.AtualizarProduto(&produto)
	}
	http.Redirect(w, r, "/", 301)
}
