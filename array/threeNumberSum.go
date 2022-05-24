package main

import (
	"fmt"
	"sort"
)

//https://www.algoexpert.io/questions/Three%20Number%20Sum

func main() {
	fmt.Println(ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0))
}

func ThreeNumberSum(array []int, target int) [][]int {
	// Write your code here.
	sort.Ints(array)
	result := [][]int{}
	for i := 0; i < len(array); i++ {
		left, right := i+1, len(array)-1
		for left < right {
			currentSum := array[i] + array[left] + array[right]
			if currentSum == target {
				result = append(result, []int{array[i], array[left], array[right]})
				left++
				right--
			} else if currentSum > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
