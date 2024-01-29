package main

import "fmt"

func quicksort(list []int) []int {
	if len(list) < 2 {
		return list
	} else {
		pivot := list[0]

		var menor = []int{}
		var maior = []int{}
		for _, num := range list[1:] {
			if pivot > num {
				menor = append(menor, num)
			} else {
				maior = append(maior, num)
			}
		}
		menor = append(quicksort(menor), pivot)
		maior = quicksort(maior)
		return append(menor, maior...)
	}
}

func main() {
	fmt.Println(quicksort([]int{10, 5, 2, 3}))
}
