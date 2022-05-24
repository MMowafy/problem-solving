package main

//https://www.algoexpert.io/questions/Array%20Of%20Products

func main() {
	ArrayOfProducts([]int{5, 1, 4, 2})
}

func ArrayOfProducts(array []int) []int {
	// Write your code here.
	products := make([]int, len(array))
	for i := range products {
		products[i] = 1
	}

	leftSideProducts := 1
	for i := range array {
		products[i] = leftSideProducts
		leftSideProducts *= array[i]
	}

	rightSideProducts := 1
	for i := len(array) - 1; i >= 0; i-- {
		products[i] *= rightSideProducts
		rightSideProducts *= array[i]
	}
	return products
}
