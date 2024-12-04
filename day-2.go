package main

import (
  "os"
  "fmt"
)

func Day2() {
  rawInput, err := GetInputHTTP("2")
  if err != nil {
    os.Exit(1)
  }

  nums, err := ExtractNumbersByLine(rawInput)
  if err != nil {
    os.Exit(2)
  }

	fmt.Println("Day 2 - Part 1:", NumberOfSafeLines(nums))
	fmt.Println("Day 2 - Part 2:", NumberOfSafeLinesWithDampening(nums))
}

