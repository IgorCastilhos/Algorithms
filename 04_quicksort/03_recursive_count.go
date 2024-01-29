package main

import "fmt"

func contar(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return 1 + contar(list[1:])
}

func main() {
	fmt.Println(contar([]int{0, 1, 2, 3, 4, 5}))
}
