package main

import (
	"fmt"
	"sort"
)

//https://www.algoexpert.io/questions/Sorted%20Squared%20Array

func main() {
	fmt.Println(SortedSquaredArray([]int{1, 2, 3, 5, 6, 8, 9}))
}
func SortedSquaredArray(array []int) []int {
	result := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		result[i] = array[i] * array[i]
	}
	sort.Ints(result)
	return result
}
