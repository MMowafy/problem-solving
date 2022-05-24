package main

import "fmt"

//https://www.algoexpert.io/questions/Validate%20Subsequence

func main() {
	fmt.Println(IsValidSubsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}))
}

func IsValidSubsequence(array []int, sequence []int) bool {
	position := 0
	for _, value := range array {
		if position == len(sequence) {
			break
		}
		if value == sequence[position] {
			position++
		}
	}
	return position == len(sequence)
}
