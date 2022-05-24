package main

import "fmt"

// https://www.algoexpert.io/questions/Spiral%20Traverse

func main() {
	fmt.Println(SpiralTraverse([][]int{
		{1},
		{3},
		{2},
		{5},
		{4},
		{7},
		{6},
	}))
}

func SpiralTraverse(array [][]int) []int {
	// Write your code here.
	capacity := len(array) * len(array[0])
	i := 0
	startingRow := 0
	endingRow := len(array) - 1
	startingColumn := 0
	endingColumn := len(array[0]) - 1
	var result []int
	for i < capacity {

		for column := startingColumn; column <= endingColumn; column++ {
			result = append(result, array[startingRow][column])
			i++
		}

		for row := startingRow + 1; row <= endingRow; row++ {
			result = append(result, array[row][endingColumn])
			i++
		}

		for column := endingColumn - 1; column >= startingColumn; column-- {
			if startingRow == endingRow {
				break
			}
			result = append(result, array[endingRow][column])
			i++
		}

		for row := endingRow - 1; row > startingRow; row-- {
			if startingColumn == endingColumn {
				break
			}
			result = append(result, array[row][startingColumn])
			i++
		}
		startingRow++
		endingRow--
		startingColumn++
		endingColumn--
	}
	return result
}
