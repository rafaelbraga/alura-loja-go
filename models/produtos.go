package models

import (
	"my_app/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaBancoDeDados()
	defer db.Close()

	selectDeTodosOsProdutos, err := db.Query("select* from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}
	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	return produtos

}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancoDeDados()
	defer db.Close()
	insereDadosNoBanco, err := db.Prepare("insert INTO produtos (nome, descricao,preco,quantidade) VALUES($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

}

func DeletaProduto(id string) {
	db := db.ConectaBancoDeDados()
	defer db.Close()
	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarOProduto.Exec(id)
}

func EditaProduto(id string) Produto {
	db := db.ConectaBancoDeDados()
	defer db.Close()
	editaOProduto, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	for editaOProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = editaOProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
	}
	return p
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBancoDeDados()
	defer db.Close()
	atualizaProduto, err := db.Prepare("update produtos set nome= $1,descricao=$2, preco=$3, quantidade = $4 where id = $5")
	if err != nil {
		panic(err.Error())
	}
	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)

}
