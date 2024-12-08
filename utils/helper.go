package utils

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
  "strings"
)

// Parses a string into pairs of integers and associated updates.
func GetPairs(input string) ([][]int, []string, error) {
	rawPairs := regexp.MustCompile(`\d+|\d+`).FindAllString(input, -1)
	rawUpdates := regexp.MustCompile(`(\d+,)+|\d+`).FindAllString(input, -1)

	pairs := make([][]int, len(rawPairs))
	for i, pair := range rawPairs {
		extracted, err := ExtractNumbers(pair)
		if err != nil {
			return nil, nil, err
		}
		pairs[i] = extracted
	}
	return pairs, rawUpdates, nil
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
        brokenRules, result := checkPertinentRules(pertinentRules, rawUpdates[index])
        if result {
          middleIndex := len(update) / 2
          sumOfMiddleNumber += update[middleIndex]
          continue
        }
        
        updatedUpdate := correctUpdate(update, brokenRules)
        if result {
          middleIndex := len(updatedUpdate) / 2
          sumOfMiddleNumberCorrected += updatedUpdate[middleIndex]
        }
        
    }
    return sumOfMiddleNumber, sumOfMiddleNumber + sumOfMiddleNumberCorrected, nil
}

func checkPertinentRules(pertinentRules [][]int, rawUpdate string) ([][]int, bool) {
  brokenRules := [][]int{}

  for _, rule := range pertinentRules {
    ruleCheck := regexp.MustCompile(strconv.Itoa(rule[0]) + ".*" + strconv.Itoa(rule[1])).FindString(rawUpdate)
    if ruleCheck == "" {
      brokenRules = append(brokenRules, rule)
    }
  }
  if len(brokenRules) == 0 {
    return brokenRules, true
  }
  return brokenRules, false 
}


func correctUpdate(update []int, brokenRules [][]int) ([]int) {
  corrected := []int{}
  copy(corrected, update)
  for _, brokenRule := range brokenRules {
    aux1 := 0
    for w, item := range update {
      if brokenRule[1] == item {
        aux1 = item
        corrected = RemoveAtIndex(corrected, w)
      }
      if brokenRule[0] == item {
        corrected = insertAtIndex(corrected, w+1, aux1)
        break
      }
    }
  }

  return corrected 
}

func ProcessDependencySequence(pairs [][]int, atualizations []string) int {
	pagesBefore := make(map[int]map[int]bool)
	pagesAfter := make(map[int]map[int]bool)

	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		if pagesBefore[b] == nil {
			pagesBefore[b] = make(map[int]bool)
		}
		pagesBefore[b][a] = true
		
		if pagesAfter[a] == nil {
			pagesAfter[a] = make(map[int]bool)
		}
		pagesAfter[a][b] = true
	}

	accumulated := 0

	for _, atu := range atualizations {
		seqStr := strings.Split(atu, ",")
		sequence := make([]int, len(seqStr))
		for i, s := range seqStr {
			sequence[i], _ = strconv.Atoi(s)
		}

		if !isSequenceValid(sequence, pagesBefore) {
			orderedSequence := makeSequenceValid(sequence, pagesBefore, pagesAfter)
			accumulated += orderedSequence[len(orderedSequence)/2]
		}
	}

	return accumulated
}

func isSequenceValid(sequence []int, pagesBefore map[int]map[int]bool) bool {
	for i, x := range sequence {
		for j, y := range sequence {
			if pagesBefore[y] != nil && pagesBefore[y][x] && i > j {
				return false
			}
		}
	}
	return true
}

func makeSequenceValid(sequence []int, pagesBefore map[int]map[int]bool, pagesAfter map[int]map[int]bool) []int {
	orderedSequence := []int{}
	queue := []int{}
	
	remainingDependencies := make(map[int]int)
	for _, v := range sequence {
		remainingDependencies[v] = 0
		for x := range pagesBefore[v] {
			if contains(sequence, x) {
				remainingDependencies[v]++
			}
		}
	}

	for v, deps := range remainingDependencies {
		if deps == 0 {
			queue = append(queue, v)
		}
	}

	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		orderedSequence = append(orderedSequence, x)

		if pagesAfter[x] != nil {
			for y := range pagesAfter[x] {
				if _, exists := remainingDependencies[y]; exists {
					remainingDependencies[y]--
					if remainingDependencies[y] == 0 {
						queue = append(queue, y)
					}
				}
			}
		}
	}

	return orderedSequence
}

