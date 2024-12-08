package std

import (
  "strconv"
  "strings"
  "fmt"
  "regexp"
)

// ===== String and Parsing Utilities ===== stringutils =====

// Converts an array of integers to a single concatenated string.
func ArrayToString(arr []int) string {
	var sb strings.Builder
	for _, num := range arr {
		sb.WriteString(fmt.Sprintf("%d", num))
	}
	return sb.String()
}

// Extracts numbers from a string into a slice of integers.
func ExtractNumbers(input string) ([]int, error) {
	return ExtractNumbersFromMatches(regexp.MustCompile(`\d+`).FindAllString(input, -1))
}

// Splits a string into lines, excluding the final empty line if present.
func ExtractByLine(input string) []string {
	lines := regexp.MustCompile(`(?m)^.*$`).FindAllString(input, -1)
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

// Extracts numbers from each line of a string into a slice of integer slices.
func ExtractNumbersByLine(input string) ([][]int, error) {
	lines := ExtractByLine(input)
	var result [][]int
	for _, line := range lines {
		numbers, err := ExtractNumbers(line)
		if err != nil {
			return result, err
		}
		result = append(result, numbers)
	}
	return result, nil
}

func ExtractNumbersFromMatches(matches []string) ([]int, error) {
	numbers := make([]int, len(matches))
	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, err
		}
		numbers[i] = num
	}
	return numbers, nil
}

// Function that creates a new slice with the indexed value x set to another value. Useful for immutable slices.
func ReplaceByIndex(str string, index int, value string) string {
  return str[:index] + value + str[index+1:]
}
