package main
  
  import (
    "fmt"
    "os"
    "advent_of_code_2024/utils/std"
    "advent_of_code_2024/utils/advent_of_code"
  )

  func main() {
    raw, err := utils.GetInputFile("5")
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    numberOfMiddleValues, _, err := utils.CountCorrectUpdatesMiddleNumber(raw)
    if err != nil {
      fmt.Println(err)
      os.Exit(2)
    }
    
    pairs, updates, err := utils.GetPairs(raw)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    
    
    
    fmt.Println("Day 5 - Part 1:", numberOfMiddleValues)
    fmt.Println("Day 5 - Part 2:", utils.ProcessDependencySequence(pairs, updates))
  }
  
