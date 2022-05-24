package main

import (
	"fmt"
	"sort"
)

//https://www.algoexpert.io/questions/Merge%20Overlapping%20Intervals

func main() {
	fmt.Println(MergeOverlappingIntervals([][]int{
		{1, 22},
		{-20, 30}}))
}

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	// Write your code here.
	sortedIntervals := make([][]int, len(intervals))
	copy(sortedIntervals, intervals)
	sort.Slice(sortedIntervals, func(i, j int) bool {
		return sortedIntervals[i][0] < sortedIntervals[j][0]
	})

	mergedIntervals := make([][]int, 0)
	currentInterval := sortedIntervals[0]
	mergedIntervals = append(mergedIntervals, currentInterval)
	for _, nextInterval := range sortedIntervals {
		currentIntervalEnd := currentInterval[1]
		nextIntervalStart := nextInterval[0]
		nextIntervalEnd := nextInterval[1]
		if currentIntervalEnd >= nextIntervalStart {
			currentInterval[1] = max(currentIntervalEnd, nextIntervalEnd)
		} else {
			currentInterval = nextInterval
			mergedIntervals = append(mergedIntervals, currentInterval)
		}
	}
	return mergedIntervals
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
