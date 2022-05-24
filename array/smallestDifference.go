package main

import (
	"math"
	"sort"
)

//https://www.algoexpert.io/questions/Smallest%20Difference

func main() {
	SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17})
}

func SmallestDifference(array1, array2 []int) []int {
	// Write your code here.
	sort.Ints(array1)
	sort.Ints(array2)
	smallest := math.MaxInt
	current := math.MaxInt
	arr1Index := 0
	arr2Index := 0
	smallestPair := []int{}

	for arr1Index < len(array1) && arr2Index < len(array2) {
		first, second := array1[arr1Index], array2[arr2Index]
		if first == second {
			return []int{first, second}
		}
		if first > second {
			current = first - second
			arr2Index++
		} else {
			current = second - first
			arr1Index++
		}
		if current < smallest {
			smallest = current
			smallestPair = []int{first, second}
		}
	}
	return smallestPair
}
