package main

import "fmt"

var votou map[string]bool

func checkarVotacao(nome string) {
	if votou[nome] {
		fmt.Println("Mande embora!")
	} else {
		votou[nome] = true
		fmt.Println("Deixe votar!")
	}
}

func main() {
	votou = make(map[string]bool)
	checkarVotacao("tom")
	checkarVotacao("mike")
	checkarVotacao("mike")
}
