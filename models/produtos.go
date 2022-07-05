package models

import (
	"alura-go-webapp/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscarTodosProdutos() []Produto {
	conn := db.ConexaoDB()
	defer conn.Close()

	resultSet, err := conn.Query("SELECT * FROM tb_produtos")
	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}

	for resultSet.Next() {
		produto := Produto{}

		err = resultSet.Scan(
			&produto.Id,
			&produto.Nome,
			&produto.Descricao,
			&produto.Preco,
			&produto.Quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtos = append(produtos, produto)
	}

	return produtos
}

func CriarProduto(p *Produto) {
	conn := db.ConexaoDB()
	defer conn.Close()

	stmt, err := conn.Prepare("INSERT INTO tb_produtos (nome, descricao, preco, quantidade) VALUES (?, ?, ?, ?)")

	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(
		p.Nome,
		p.Descricao,
		p.Preco,
		p.Quantidade)
	if err != nil {
		panic(err.Error())
	}
}

func ExcluirProdutoPorId(id string) {
	conn := db.ConexaoDB()
	defer conn.Close()

	stmt, err := conn.Prepare("DELETE FROM tb_produtos WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}
}

func BuscarProdutoPorId(id string) *Produto {
	conn := db.ConexaoDB()
	defer conn.Close()

	resultSet, err := conn.Query("SELECT * FROM tb_produtos WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for resultSet.Next() {

		err = resultSet.Scan(
			&produto.Id,
			&produto.Nome,
			&produto.Descricao,
			&produto.Preco,
			&produto.Quantidade)

		if err != nil {
			panic(err.Error())
		}
	}
	return &produto
}

func AtualizarProduto(p *Produto) {
	conn := db.ConexaoDB()
	defer conn.Close()

	stmt, err := conn.Prepare("UPDATE tb_produtos SET nome = ?, descricao = ?, preco = ?, quantidade = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(
		p.Nome,
		p.Descricao,
		p.Preco,
		p.Quantidade,
		p.Id)
	if err != nil {
		panic(err.Error())
	}
}
