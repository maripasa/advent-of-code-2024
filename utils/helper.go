package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
  "strings"
)

func ManageAdventOfCodePuzzles(year int) (error) {
  if os.Getenv("SESSION") == "" {
    return errors.New("Session not found in environment.")
  }


  today := time.Now().Truncate(2 * time.Hour)
  if today.Day() > 25 || today.Year() != year || today.Month() != time.December {
    return nil
  }
  
  for day := 1 ; day <= today.Day() ; day++ {
    filename := "day"
    if day < 10 { filename += "0" }
    filename += strconv.Itoa(day)
    
    _, err := os.Stat(filename)

    if err != nil {
      if !os.IsNotExist(err) {
        return err
      }
      err := os.MkdirAll(filename, 0755)
      if err != nil {
        return err
      }
    }

    _, err = os.Stat(filename + "/" + filename + ".go")

    if err != nil {
      if !os.IsNotExist(err) {
        return err
      }
      err = createGoPuzzleFileWithBoilerplate(strconv.Itoa(day), filename + "/" + filename + ".go")
      if err != nil {
        return err
      }
    }
    
    _, err = os.Stat(filename + "/input.txt")

    if err != nil {
      if !os.IsNotExist(err) {
        return err
      }
      raw, err := GetInputHTTP(strconv.Itoa(day))
      if err != nil {
        return err
      }
      err = createFileAndWriteString(filename + "/input.txt", raw)
      if err != nil {
        return err
      }
      }

    _, err = os.Stat(filename + "go.mod")

    if err != nil {
      if !os.IsNotExist(err) {
        return err
      }
      raw, err := GetInputHTTP(strconv.Itoa(day))
      if err != nil {
        return err
      }
      err = createFileAndWriteString(filename + "/input.txt", raw)
      if err != nil {
        return err
      }
    }
  }
  return nil 
}


func createGoPuzzleFileWithBoilerplate(day string, filename string) (error) {
  content := `package main
  
  import (
    "fmt"
    "os"
    "advent_of_code_2024/utils"
  )

  func main() {
    raw, err := utils.GetInputFile(` + day + `)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    fmt.Println("Day ` + day + ` - Part 1:")
    fmt.Println("Day ` + day + ` - Part 2:")
  }
  `
  return createFileAndWriteString(filename, content)
}

func createFileAndWriteString(filename string, content string) (error) {
  file, err := os.Create(filename)
  if err != nil {
    return err
  }
  defer file.Close()
	_, err = file.WriteString(content)
  return err 
}

func GetInputFile(day string) (string, error){
  inputFileName := "day"
  intDay, err := strconv.Atoi(day) 
  if err != nil { return "", err }
  if intDay < 10 { inputFileName += "0" }
  inputFileName += day + "/input.txt"
  file, err := os.ReadFile(inputFileName)
  return string(file), err
}

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

func ExtractByLine(input string) ([]string) {
  extracted := regexp.MustCompile(`(?m)^.*$`).FindAllString(input, -1)
  return extracted[:len(extracted)-1]
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

func FindX_mas(matrix []string) (int) {
  output := 0
  for i, line := range matrix {
    for j, item := range line {
      if item != 'A'{
        continue
      }
      if i < 1 || j < 1 || i > len(matrix) - 2 || j > len(matrix[i]) - 2 {
        continue
      }
      if matrix[i-1][j-1] == 'M' && matrix[i+1][j-1] == 'M'{
        if matrix[i+1][j+1] == 'S' && matrix[i-1][j+1] == 'S'{
          output += 1
        }
      }
      if matrix[i-1][j-1] == 'M' && matrix[i-1][j+1] == 'M'{
        if matrix[i+1][j+1] == 'S' && matrix[i+1][j-1] == 'S'{
          output += 1
        }
      }
      if matrix[i-1][j+1] == 'M' && matrix[i+1][j+1] == 'M'{
        if matrix[i+1][j-1] == 'S' && matrix[i-1][j-1] == 'S'{
          output += 1
        }
      }
      if matrix[i+1][j-1] == 'M' && matrix[i+1][j+1] == 'M'{
        if matrix[i-1][j+1] == 'S' && matrix[i-1][j-1] == 'S'{
          output += 1
        }
      }

    }
  }
  return output
}

func FindXmas(matrix []string) (int) {
  output := 0
  for i, line := range matrix {
    for j, item := range line {
      if item != 'X'{
        continue
      }
      // UP
      if i >= 3 {
        if matrix[i-1][j] == 'M' && matrix[i-2][j] == 'A' && matrix[i-3][j] == 'S' {
          output += 1
        }
      }
      // DOWN
      if i <= len(matrix)-4 {
        if matrix[i+1][j] == 'M' && matrix[i+2][j] == 'A' && matrix[i+3][j] == 'S' {
          output += 1
        }
      }
      // LEFT
      if j >= 3 {
        if matrix[i][j-1] == 'M' && matrix[i][j-2] == 'A' && matrix[i][j-3] == 'S' {
          output += 1
        }
      }
      // RIGHT
      if j <= len(line) - 4 {
        if matrix[i][j+1] == 'M' && matrix[i][j+2] == 'A' && matrix[i][j+3] == 'S' {
          output += 1
        }
      }
      // left upper diagonal
      if j >= 3 && i >= 3 {
        if matrix[i-1][j-1] == 'M' && matrix[i-2][j-2] == 'A' && matrix[i-3][j-3] == 'S' {
          output += 1
        }
      }
      // right upper diagonal
      if j <= len(line) - 4 && i >= 3 {
        if matrix[i-1][j+1] == 'M' && matrix[i-2][j+2] == 'A' && matrix[i-3][j+3] == 'S' {
          output += 1
        }
      }
      // left downwards diagonal
      if j >= 3 && i <= len(matrix)-4 {
        if matrix[i+1][j-1] == 'M' && matrix[i+2][j-2] == 'A' && matrix[i+3][j-3] == 'S' {
          output += 1
        }
      }
      // right downwards diagonal
      if j <= len(line)-4 && i <= len(matrix)-4 {
        if matrix[i+1][j+1] == 'M' && matrix[i+2][j+2] == 'A' && matrix[i+3][j+3] == 'S' {
          output += 1
        }
      }
    }
  }
  return output
}

/* func Filter()
func Reduce()
func Map() */

func CalculateErrorSum(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	
	errorSum := 0
	for i := range left {
		errorSum += Abs(left[i] - right[i])
	}
	return errorSum
}

func CalculateSimilarityScore(left, right []int) int {
	score := 0
	for index := range left {
		score += left[index] * Count(right, left[index])
	}
	return score
}

func CountCorrectUpdatesMiddleNumber(input string) (int, int, error) {
    sumOfMiddleNumber := 0
    sumOfMiddleNumberCorrected := 0

    rawUpdates := regexp.MustCompile(`(\d+,)+\d+`).FindAllString(input, -1)
    updates := make([][]int, len(rawUpdates))

    for i, update := range rawUpdates {
        extracted, err := ExtractNumbers(update)
        if err != nil {
            return 0, 0, err
        }
        updates[i] = extracted
    }

    updateNumber := len(updates)

    for index, update := range updates {
        fmt.Println("Checking:", index+1, "/", updateNumber)

        possibleRules := [][]int{}
        for i := 0; i < len(update); i++ {
            for j := 0; j < len(update); j++ {
                if i != j {
                    pair := []int{update[i], update[j]}
                    possibleRules = append(possibleRules, pair)
                }
            }
        }

        pertinentRules := [][]int{}
        for _, rule := range possibleRules {
            FilterRule := regexp.MustCompile(strconv.Itoa(rule[0]) + `\|` + strconv.Itoa(rule[1])).FindString(input)
            if FilterRule == "" {
                continue
            }
            pertinentRules = append(pertinentRules, rule)
        }

        // Check if the update satisfies the pertinent rules
        if checkPertinentRules(pertinentRules, rawUpdates[index]) {
          middleIndex := len(update) / 2
          sumOfMiddleNumber += update[middleIndex]
          continue
        }
        
        numOfPerm := Factorial(len(update))
        currentPerm := 0
        for perm := range generatePermutations(update) {
          currentPerm += 1
          fmt.Println("Permutation:", currentPerm, "/", numOfPerm)
          permString := arrayToString(perm)
          if checkPertinentRules(pertinentRules, permString) {
            middleIndex := len(perm) / 2
            sumOfMiddleNumberCorrected += perm[middleIndex]
            continue
          }
        }
    }
    return sumOfMiddleNumber, sumOfMiddleNumber + sumOfMiddleNumberCorrected, nil
}

func checkPertinentRules(pertinentRules [][]int, rawUpdate string) bool {
  for _, rule := range pertinentRules {
    ruleCheck := regexp.MustCompile(strconv.Itoa(rule[0]) + ".*" + strconv.Itoa(rule[1])).FindString(rawUpdate)
    if ruleCheck == "" {
      return false
    }
  }
  return true
}

func arrayToString(arr []int) string {
  var sb strings.Builder
  for _, num := range arr {
    sb.WriteString(fmt.Sprintf("%d", num))
  }
  return sb.String()
}

func generatePermutations(original []int) <-chan []int {
  ch := make(chan []int)
  
  go func() {
    defer close(ch)
    
    if len(original) <= 1 {
      ch <- original
      return
    }
    
    var permute func([]int, int)
    permute = func(arr []int, k int) {
      if k == len(arr)-1 {
        perm := make([]int, len(arr))
        copy(perm, arr)
        ch <- perm
        return
      }
      
      for i := k; i < len(arr); i++ {
        arr[k], arr[i] = arr[i], arr[k]
        permute(arr, k+1)
        arr[k], arr[i] = arr[i], arr[k]
      }
    }
    
    permute(original, 0)
  }()
  
  return ch
}

func Factorial(n int) int {
  result := 1
  for i := 1; i <= n; i++ {
    result *= i
  }
  return result
}
