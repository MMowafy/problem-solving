package main

//https://www.algoexpert.io/questions/First%20Duplicate%20Value

func main() {
	FirstDuplicateValue([]int{2, 1, 5, 2, 3, 3, 4})
}
func FirstDuplicateValue(array []int) int {
	// Write your code here.
	arrMap := make(map[int]int)
	for i := range array {
		arrMap[array[i]] += 1
		if arrMap[array[i]] > 1 {
			return array[i]
		}
	}
	return -1
}
