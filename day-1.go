package main

import (
	"fmt"
	"os"
	"sort"
)

//To use must export session
//
//Done with:
//export SESSION=<your_advent_of_code_session>

func Day1() {
	input, err := GetInputHTTP("1")
	if err != nil {
		os.Exit(1)
	}
	nums, err := ExtractNumbers(input)

	if err != nil {
		os.Exit(2)
	}

	left, right := make([]int, 0, len(nums)/2), make([]int, 0, len(nums)/2)

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			left = append(left, nums[i])
			continue
		}
		right = append(right, nums[i])
	}

	fmt.Println("Day 1 - Part 1:", calculateErrorSum(left, right))
	fmt.Println("Day 1 - Part 2:", calculateSimilarityScore(left, right))
}

func calculateErrorSum(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	
	errorSum := 0
	for i := range left {
		errorSum += Abs(left[i] - right[i])
	}
	return errorSum
}

func calculateSimilarityScore(left, right []int) int {
	score := 0
	for index := range left {
		score += left[index] * Count(right, left[index])
	}
	return score
}


