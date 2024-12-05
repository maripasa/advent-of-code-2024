package main

import (
  "os"
  "fmt"
  "advent_of_code_2024/utils"
)

func main() {
  rawInput, err := utils.GetInputFile("2")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  nums, err := utils.ExtractNumbersByLine(rawInput)
  if err != nil {
    fmt.Println(err)
    os.Exit(2)
  }

	fmt.Println("Day 2 - Part 1:", utils.NumberOfSafeLines(nums))
	fmt.Println("Day 2 - Part 2:", utils.NumberOfSafeLinesWithDampening(nums))
}

