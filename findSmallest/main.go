package main

import "fmt"

func LowestSearch(list []int) int {
	menor := list[0]
	menorIndice := 0

	for i := range list {
		if list[i] < menor {
			menor = list[i]
			menorIndice = i
		}
	}

	return menorIndice
}

func main() {
	lista := []int{5, 46, 346, 46, 2, 46, 75, 53, 745, 345}
	fmt.Println(LowestSearch(lista))
}
