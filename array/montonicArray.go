package main

import "fmt"

//algoexpert.io/questions/Monotonic%20Array
func main() {
	fmt.Println(IsMonotonic([]int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}))
}

func IsMonotonic(array []int) bool {
	if len(array) < 2 {
		return true
	}
	directionFound := false
	direction := 0
	for i := 0; i < len(array)-1; i++ {
		if array[i] != array[i+1] && !directionFound {
			direction = getDirection(array[i], array[i+1])
			directionFound = true
		}
		if direction == 0 && array[i] < array[i+1] {
			return false
		}
		if direction == 1 && array[i] > array[i+1] {
			return false
		}
	}
	return true
}

func getDirection(first int, second int) int {
	if first > second {
		return 0
	}
	return 1
}
