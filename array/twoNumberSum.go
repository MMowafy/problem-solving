package main

import "fmt"

//https://www.algoexpert.io/questions/Two%20Number%20Sum

func main() {
	fmt.Println(TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10))
}

func TwoNumberSum(array []int, target int) []int {
	arrMap := make(map[int]bool, len(array))
	for _, item := range array {
		arrMap[item] = true
	}
	for _, item := range array {
		remaining := target - item
		_, ok := arrMap[remaining]
		if ok && remaining != item {
			return []int{item, remaining}
		}
	}
	return []int{}
}
