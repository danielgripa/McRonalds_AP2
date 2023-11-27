package produtos

import (
	m "McRonalds/metricas"
	"fmt"
	"strings"
)

type ProdutoNode struct {
	Produto
	Next *ProdutoNode
}

var primeiroProduto *ProdutoNode
var totalProdutos = 0

func tentarCriar(nome, descricao string, preco float64, id int) *ProdutoNode {
	if id != -1 {
		_, idProcurado := BuscarId(id)
		if idProcurado != nil {
			return nil // Retorna nil se o ID já existir
		}
	}

	maxExistingID := 0
	current := primeiroProduto
	for current != nil {
		if current.Id > maxExistingID {
			maxExistingID = current.Id
		}
		current = current.Next
	}
	nextAvailableID := maxExistingID + 1

	return &ProdutoNode{Produto: criar(nome, descricao, preco, nextAvailableID)}
}

func AdicionarUnico(nome, descricao string, preco float64, id int) int {
	if totalProdutos == 0 {
		produtoCriado := tentarCriar(nome, descricao, preco, id)
		if produtoCriado == nil {
			return -2
		}
		primeiroProduto = produtoCriado
		totalProdutos++
		m.M.SomaProdutosCadastrados(1)
		return totalProdutos
	}

	// Verificar se o produto já existe
	current := primeiroProduto
	for current != nil {
		if current.Nome == nome {
			return -2
		}
		current = current.Next
	}

	produtoCriado := tentarCriar(nome, descricao, preco, id)
	if produtoCriado == nil {
		return -3
	}

	// Adicionar no início da lista
	produtoCriado.Next = primeiroProduto
	primeiroProduto = produtoCriado
	totalProdutos++
	m.M.SomaProdutosCadastrados(1)
	return totalProdutos
}

func BuscarId(id int) (*ProdutoNode, *ProdutoNode) {
	var anterior *ProdutoNode
	current := primeiroProduto
	for current != nil {
		if current.Id == id {
			return current, anterior
		}
		anterior = current
		current = current.Next
	}
	return nil, nil
}

func BuscarNome(comecaCom string) ([]Produto, int) {
	var produtosEncontrados []Produto

	current := primeiroProduto
	for current != nil {
		if strings.HasPrefix(current.Nome, comecaCom) {
			produtosEncontrados = append(produtosEncontrados, current.Produto)
		}
		current = current.Next
	}
	return produtosEncontrados, len(produtosEncontrados)
}

func Exibir() {
	current := primeiroProduto
	for current != nil {
		fmt.Println("\nProduto", current.Id)
		fmt.Println("Nome:", current.Nome)
		fmt.Println("Descrição:", current.Descricao)
		fmt.Printf("Preço: R$ %.2f\n", current.Preco)

		current = current.Next
	}
}

func Excluir(id int) int {
	if totalProdutos == 0 {
		return -2
	}

	produto, anterior := BuscarId(id)
	if produto == nil {
		return -1
	}

	if anterior == nil {
		// Se o produto for o primeiro na lista
		primeiroProduto = produto.Next
	} else {
		anterior.Next = produto.Next
	}

	totalProdutos--
	m.M.SomaProdutosCadastrados(-1)
	return 0
}

/*
Cria uma cópia da lista de produtos e organiza em ordem alfabética
*/
func ExibirOrdenadoPorNome() {

	copia := CopiarLista(primeiroProduto) // Cria uma cópia da lista encadeada

	// Bubble Sort para classificação dos produtos em ordem alfabética
	for i := 0; i < totalProdutos; i++ {
		atual := copia
		for j := 0; j < totalProdutos-i-1; j++ {
			if atual.Produto.Nome > atual.Next.Produto.Nome {
				atual.Produto, atual.Next.Produto = atual.Next.Produto, atual.Produto
			}

			atual = atual.Next
		}
	}

	for copia != nil {
		copia.Produto.Exibir()
		copia = copia.Next
	}
}

/*
Chama a função BuscarID para encontrar o produto que precisa ter o preço modificado e faz a alteração.
*/
func AtualizarPreco(id int, novoPreco float64) int {
	produto, _ := BuscarId(id)
	if produto == nil {
		return -1 // Produto não encontrado
	}

	produto.Preco = novoPreco
	return 0 // Produto encontrado
}

/*
Cria uma cópia da lista encadeada para que a função ExibirOrdenadoPorNome não desorganize a função Exibir que está por ordem de cadastro de ID.
*/
func CopiarLista(original *ProdutoNode) *ProdutoNode {

	if original == nil {
		return nil
	}

	copia := &ProdutoNode{
		Produto: original.Produto,
		Next:    nil,
	}

	ultimoCopia := copia

	for atual := original.Next; atual != nil; atual = atual.Next {
		novoNo := &ProdutoNode{
			Produto: atual.Produto,
			Next:    nil,
		}
		ultimoCopia.Next = novoNo
		ultimoCopia = novoNo
	}

	return copia
}
