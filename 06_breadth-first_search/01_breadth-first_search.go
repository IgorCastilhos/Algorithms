package main

import "fmt"

// pessoaForVendedora retornará o nome da pessoa que tiver m no final da letra, esta pessoa será uma vendedora de manga
func pessoaForVendedora(nome string) bool {
	return nome[len(nome)-1] == 'e'
}

// grafo cria o hash contendo os nomes das pessoas como chaves e seus "amigos" como valores
var grafo = make(map[string][]string)

func main() {
	grafo["voce"] = []string{"alice", "bob", "claire"}
	grafo["bob"] = []string{"andre", "pedro"}
	grafo["alice"] = []string{"pedro"}
	grafo["claire"] = []string{"tom", "jon"}
	grafo["andre"] = []string{}
	grafo["pedro"] = []string{}
	grafo["tom"] = []string{}
	grafo["jon"] = []string{}
	pesquisa("voce")
	pesquisa("bob")
	pesquisa("alice")
}

func pesquisa(nome string) bool {
	var filaDePesquisa []string
	filaDePesquisa = append(filaDePesquisa, grafo[nome]...)
	var verificadas []string
	var pessoa string

	// Enquanto a fila não estiver vazia...
	for len(filaDePesquisa) != 0 {
		// ... pega a primeira pessoa da fila
		pessoa = filaDePesquisa[0]
		// Tira a pessoa que está verificando da lista
		filaDePesquisa = filaDePesquisa[1:]
		// Checa se essa pessoa já foi verificada
		if pessoaNaoEstaEmVerificadas(pessoa, verificadas) {
			// Verifica se essa pessoa é uma vendedora de mangas
			if pessoaForVendedora(pessoa) {
				// Sim, ela é uma vendedora de mangas
				fmt.Println(pessoa + " é vendedor(a) de manga!")
				return true
			}
			// Não, essa pessoa não é uma vendedora de mangas. Adiciona todos os amigos dessa pessoa à lista
			filaDePesquisa = append(filaDePesquisa, grafo[pessoa]...)
			verificadas = append(verificadas, pessoa)
		}
	}
	// Se chegou até aqui, é sinal que nenhuma pessoa da fila era uma vendedora de mangas
	return false
}

// pessoaNaoEstaEmVerificadas verifica se o nome da pessoa não está no slice de pessoas já verificadas
func pessoaNaoEstaEmVerificadas(nome string, verificadas []string) bool {
	for _, n := range verificadas {
		if n == nome {
			// Essa pessoa já foi verificada!
			return false
		}
	}
	// Essa pessoa ainda não foi verificada!
	return true
}
