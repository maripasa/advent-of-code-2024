package main

import (
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func GetInputHTTP(day string) (string, error) {
	req, err := http.NewRequest("GET", "https://adventofcode.com/2024/day/" + day + "/input", nil)
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: os.Getenv("SESSION"),
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Count(slice []int, val int) int {
	count := 0
	for index := range slice {
		if slice[index] == val {
			count++
		}
	}
	return count
}

func ExtractNumbersByLine(input string) ([][]int, error) {
  lines := regexp.MustCompile(`(?m)^.*$`).FindAllString(input, -1)
	extracted := [][]int{}
  for _, line := range lines {
    result, err := ExtractNumbers(line)
    if err != nil {
		  return extracted, err
    }
    extracted = append(extracted, result)
  }
  extracted = extracted[:len(extracted)-1]
  return extracted, nil
}

func ExtractNumbers(input string) ([]int, error) {
	nums := regexp.MustCompile(`\d+`).FindAllString(input, -1)
	extracted := make([]int, len(nums))
	
	for i, num := range nums {
		result, err := strconv.Atoi(num)
    if err != nil {
		  return extracted, err
    }

    extracted[i] = result
	}
	return extracted, nil 
}

func NumberOfSafeLines(lines [][]int) (int) {
  output := 0
  for _, line := range lines {
    var crescent bool
    if len(line) == 1 {
      output += 1
      continue
    }

    if line[0] > line[1] { crescent = false }
    if line[0] < line[1] { crescent = true }
    if line[0] == line[1] { continue }

    if isSequencePart2(line, crescent) { output += 1 }
  }
  return output
}

func isSequencePart2(line []int, crescent bool) (bool) {
  for i := range line {
    if i == 0 {
      continue
    }
    if Abs(line[i-1]-line[i]) > 3 { return false }
    if line[i-1] == line[i] { return false }
    if crescent {
      if line[i-1] > line[i] { return false }
      continue
    }
    if line[i-1] < line[i] { return false }
    continue
  }
  return true
}

func NumberOfSafeLinesWithDampening(lines [][]int) (int) {
  output := 0
  for _, line := range lines {
    var crescent bool
    if len(line) == 1 {
      output += 1
      continue
    }

    if line[0] > line[1] { crescent = false }
    if line[0] < line[1] { crescent = true }

    if isSequencePart2(line, crescent) {
      output += 1
      continue
    }
    
    dampeningVersions := createCombinationsWithOneMissing(line)
    for _, newLine := range dampeningVersions {
      if len(newLine) == 1 {
        output += 1
        break 
      }

      if newLine[0] > newLine[1] { crescent = false }
      if newLine[0] < newLine[1] { crescent = true }

      if isSequencePart2(newLine, crescent) {
        output += 1
        break 
    }
    }

  }
  return output
}

func createCombinationsWithOneMissing(line []int) ([][]int){
	extracted := make([][]int, len(line))
  for i := range line {
    extracted[i] = append([]int(nil), line[:i]...)
    extracted[i] = append(extracted[i], line[i+1:]...)
  }
  return extracted
}

func ExtractMuls(input string) ([][]int, error) {
  muls := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllString(input, -1)

	extracted := [][]int{}

  for _, line := range muls {
    result, err := ExtractNumbers(line)
    if err != nil {
		  return extracted, err
    }
    extracted = append(extracted, result)
  }

  return extracted, nil
}

func RemoveNewLines(input string) (string) {
  lines := regexp.MustCompile(`(?m)^.*$`).FindAllString(input, -1)
  output := ""
  for _, line := range lines {output += line}
  
  return output
}

func RemoveDontDo(input string) (string) {
  return regexp.MustCompile(`don't\(\).*?(do\(\)|$)`).ReplaceAllString(input, "")
}

