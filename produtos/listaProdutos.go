package produtos

import (
	m "McRonalds/metricas"
	"fmt"
	"strings"
)

const maxProdutos = 50

var Produtos [maxProdutos]Produto
var totalProdutos = 0

func tentarCriar(nome, descricao string, preco float64, id int) Produto {
	if id != -1 {
		_, idProcurado := BuscarId(id)
		if idProcurado != -1 {
			return Produto{} // Retorna um produto vazio se o ID já existir
		}
	}

	// Procurar o maior ID existente e atribuir o próximo ID disponível
	maxExistingID := 0
	for _, produto := range Produtos {
		if produto.Id > maxExistingID {
			maxExistingID = produto.Id
		}
	}
	nextAvailableID := maxExistingID + 1

	return criar(nome, descricao, preco, nextAvailableID)
}

func AdicionarUnico(nome, descricao string, preco float64, id int) int {
	if totalProdutos == maxProdutos {
		return -1
	} // Overflow

	for _, produto := range Produtos {
		if (produto == Produto{}) {
			break
		}
		if produto.Nome == nome {
			return -2
		}
	}

	produtoCriado := tentarCriar(nome, descricao, preco, id)
	if (produtoCriado == Produto{}) {
		return -3
	}

	Produtos[totalProdutos] = produtoCriado
	totalProdutos++
	m.M.SomaProdutosCadastrados(1)
	return totalProdutos
}

func BuscarId(id int) (Produto, int) {
	for ind, produto := range Produtos {
		if (produto == Produto{}) {
			break
		}
		if produto.Id == id {
			return produto, ind
		}
	}

	return Produto{}, -1
}

func BuscarNome(comecaCom string) ([]Produto, int) {
	var produtosEncontrados []Produto

	for _, produto := range Produtos {
		if (produto == Produto{}) {
			break
		}

		if strings.HasPrefix(produto.Nome, comecaCom) {
			produtosEncontrados = append(produtosEncontrados, produto)
		}
	}
	return produtosEncontrados, len(produtosEncontrados)
}

func Exibir() {
	for _, produto := range Produtos {
		if (produto != Produto{}) {
			fmt.Println("\nProduto", produto.Id)
			fmt.Println("Nome:", produto.Nome)
			fmt.Println("Descrição:", produto.Descricao)
			fmt.Printf("Preço: R$ %.2f\n", produto.Preco)
		}
	}
}

func Excluir(id int) int {
	if totalProdutos == 0 {
		return -2
	}

	_, ind := BuscarId(id)
	if ind == -1 {
		return -1
	}

	for i := ind; i < totalProdutos-1; i++ {
		Produtos[i] = Produtos[i+1]
	}
	totalProdutos--
	Produtos[totalProdutos] = Produto{}
	m.M.SomaProdutosCadastrados(-1)
	return 0
}

// ExibirOrdenadoPorNome exibe os produtos ordenados por nome.
func ExibirOrdenadoPorNome() {
	// Implementação simplificada do Bubble Sort
	for i := 0; i < totalProdutos; i++ {
		for j := 0; j < totalProdutos-i-1; j++ {
			if Produtos[j].Nome > Produtos[j+1].Nome {
				Produtos[j], Produtos[j+1] = Produtos[j+1], Produtos[j]
			}
		}
	}

	// Exibindo produtos ordenados
	for _, produto := range Produtos {
		if (produto == Produto{}) {
			break
		}
		produto.Exibir()
	}
}

// AtualizarPreco atualiza o preço de um produto existente.
func AtualizarPreco(id int, novoPreco float64) int {
	produto, ind := BuscarId(id)
	if ind == -1 {
		return -1 // Produto não encontrado
	}

	produto.Preco = novoPreco
	Produtos[ind] = produto
	return 0 // Produto encontrado
}
