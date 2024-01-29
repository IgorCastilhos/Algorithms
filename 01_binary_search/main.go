package main

import "fmt"

func buscaBinaria(lista []int, item int) int {
	baixo := 0
	alto := len(lista) - 1
	fmt.Println(alto)
	for baixo <= alto {
		meio := (baixo + alto) / 2
		if lista[meio] == item {
			return meio
		}
		if lista[meio] > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(buscaBinaria([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 40, 50, 60, 75, 100, 102}, 24))
	fmt.Println(buscaBinaria([]int{-1, 0, 3, 5, 9, 12}, 9))
}
