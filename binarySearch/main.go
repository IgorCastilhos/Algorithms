package main

import "fmt"

func buscaBinaria(lista []int, item int) int {
	baixo := 0
	alto := len(lista) - 1

	for baixo <= alto {
		meio := (baixo + alto) / 2
		chute := lista[meio]

		if chute == item {
			return meio
		}
		if chute > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}

	return item
}

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 40, 50, 60, 75, 100}
	fmt.Println(buscaBinaria(lista, 100))
}
