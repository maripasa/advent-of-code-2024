package main

import (
	"fmt"
  "strconv"
	"advent_of_code_2024/utils/std"
	"advent_of_code_2024/utils/aoc"
)

func main() {
	input, err := aoc.GetInputFile("9")
	std.Error(err, 1)
  
  fmt.Println(uncompressData(input))

	fmt.Println("Day 9 - Part 1:")
	fmt.Println("Day 9 - Part 2:")
}

func uncompressData(input string) []struct{
  size int
  id int
} {
  uncompressedData := []struct{
    size int
    id int
  }{}

  for i, value := range input {
    value, _ := strconv.Atoi(string(value))
    if i % 2 == 0 {
      uncompressedData = append(uncompressedData, struct{size int; id int}{value, i/2})
      continue
    }
    uncompressedData = append(uncompressedData, struct{size int; id int}{value, -1})
  }
  return uncompressedData
}

func compressData(input []struct{
  size int
  id int
}) []struct{
  size int
  id int
}{
  for i := 0 ; i < len(input) ; i++{
    if input[i].id != -1 {
      continue
    }
    firstBlock := input[i]
    lastBlock := len(input) - 1
    for firstBlock.size != 0 {
      for ; input[lastBlock].id == -1 ; lastBlock-- {}
      input = std.InsertAtIndex(input, i, struct{
        size int
        id int
      }{ 0, input[lastBlock].id })
      for ; 
    }
  }
}
