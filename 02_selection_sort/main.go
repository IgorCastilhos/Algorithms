package main

import "fmt"

// Encontra o menor valor em um array
func findSmallest(arr []int) int {
	// Armazena o menor valor
	smallest := arr[0]

	// Armazena o índice do menor valor
	smallest_index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}
func selectionSort(arr []int) []int {
	size := len(arr)
	newArr := make([]int, size)
	for i := 0; i < size; i++ {
		// Encontra o menor elemento no array e o adiciona ao novo array
		smallest := findSmallest(arr)
		newArr[i] = arr[smallest]
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}
func main() {
	s := []int{5, 3, 6, 2, 20}
	fmt.Println(selectionSort(s))
}
