package main
  
  import (
    "fmt"
    "os"
    "advent_of_code_2024/utils"
  )

  func main() {
    raw, err := utils.GetInputFile("5")
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    numberOfMiddleValues, numberAfterCorrected, err := utils.CountCorrectUpdatesMiddleNumber(raw)
    if err != nil {
      fmt.Println(err)
      os.Exit(2)
    }
    fmt.Println("Day 5 - Part 1:", numberOfMiddleValues)
    fmt.Println("Day 5 - Part 2:", numberAfterCorrected)
  }
  
