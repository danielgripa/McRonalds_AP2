package itens

import (
	p "McRonalds/produtos"
)

type Item struct {
	Prod  *p.ProdutoNode // Alterado para *p.ProdutoNode
	Quant int
}

// CalcularPrecoParcial calcula o preço parcial do item.
func (i *Item) CalcularPrecoParcial() float64 {
	if i == nil || i.Prod == nil {
		return 0.0
	}
	return i.Prod.Produto.Preco * float64(i.Quant) // Alterado para i.Prod.Produto.Preco
}

/*
Criar retorna um Item com as informações solicitadas.
Se o id não existir para um produto, retorna um Item vazio.
*/
func Criar(id int, quant int) Item {
	produtoNode, _ := p.BuscarId(id) // Alterado para produtoNode
	if produtoNode == nil {
		return Item{}
	}

	return Item{Prod: produtoNode, Quant: quant}
}
