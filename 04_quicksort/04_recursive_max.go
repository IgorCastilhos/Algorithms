package main

import "fmt"

func maximo(list []int) int {
	if len(list) == 2 {
		if list[0] > list[1] {
			return list[0]
		}
		return list[1]
	}

	sub_max := maximo(list[1:])
	if list[0] > sub_max {
		return list[0]
	}
	return sub_max
}

func main() {
	fmt.Println(maximo([]int{1, 5, 10, 25, 16, 1}))
}
