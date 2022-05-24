package main

import (
	"fmt"
)

//https://www.algoexpert.io/questions/Move%20Element%20To%20End

func main() {
	fmt.Println(MoveElementToEnd([]int{2, 1, 2, 2, 2, 3, 4, 2}, 2))
}
func MoveElementToEnd(array []int, toMove int) []int {
	// Write your code here.
	swapFirstIndex := 0
	for swapFirstIndex < len(array)-1 {
		if array[swapFirstIndex] == toMove {
			swapEndIndex := swapFirstIndex + 1
			for swapEndIndex < len(array)-1 {
				if array[swapEndIndex] == toMove {
					swapEndIndex++
				} else {
					break
				}
			}
			array[swapFirstIndex], array[swapEndIndex] = array[swapEndIndex], array[swapFirstIndex]
		}
		swapFirstIndex++
	}
	return array
}
