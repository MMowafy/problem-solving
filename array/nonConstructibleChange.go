package main

import (
	"fmt"
	"sort"
)

//https://www.algoexpert.io/questions/Non-Constructible%20Change

func main() {
	fmt.Println(NonConstructibleChange([]int{1, 2, 3, 4, 5, 6, 7}))
}
func NonConstructibleChange(coins []int) int {
	if len(coins) == 0 {
		return 1
	}
	sort.Ints(coins)
	sum := 0
	for i := 0; i < len(coins); i++ {
		if coins[i] > sum+1 {
			return sum + 1
		}
		sum += coins[i]
	}
	return sum + 1
}
