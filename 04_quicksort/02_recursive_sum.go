package main

import "fmt"

func soma(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return nums[0] + soma(nums[1:])
}

func main() {
	fmt.Println(soma([]int{1, 2, 3, 4, 5}))
}
