package main

import (
	"fmt"
	"os"
  "advent_of_code_2024/utils"
)

func main() {
	input, err := utils.GetInputFile("1")
	if err != nil {
    fmt.Println(err)
		os.Exit(1)
	}
	nums, err := utils.ExtractNumbers(input)

	if err != nil {
    fmt.Println(err)
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

	fmt.Println("Day 1 - Part 1:", utils.CalculateErrorSum(left, right))
	fmt.Println("Day 1 - Part 2:", utils.CalculateSimilarityScore(left, right))
}



