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

func selectionSort(list []int) []int {
	novoArr := make([]int, 0, len(list))

	for len(list) > 0 {
		menor := LowestSearch(list)
		novoArr = append(novoArr, list[menor])
		list = append(list[:menor], list[menor+1:]...)
	}
	return novoArr
}

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(lista)
}
