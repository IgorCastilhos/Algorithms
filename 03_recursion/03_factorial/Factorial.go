package main

import "fmt"

func fact(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * fact(num-1)
	}
}

func main() {
	fmt.Println(fact(5))
}
