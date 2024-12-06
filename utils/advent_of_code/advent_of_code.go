package aoc 

import (
	"errors"
  "strconv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// ManageAdventOfCodePuzzles sets up the puzzles for the given year.
func ManageAdventOfCodePuzzles(year int) error {
	if err := validateEnvironment(); err != nil {
		return err
	}

	today := time.Now().Truncate(2 * time.Hour)
	if !isDuringAdvent(today, year) {
		return nil
	}

	for day := 1; day <= today.Day(); day++ {
		if err := preparePuzzleFiles(day); err != nil {
			return err
		}
	}

	return nil
}

// validateEnvironment ensures the required environment variables are set.
func validateEnvironment() error {
	if os.Getenv("SESSION") == "" {
		return errors.New("Session not found in environment.")
	}
	return nil
}

// isDuringAdvent checks if today is during the Advent of Code period for the given year.
func isDuringAdvent(today time.Time, year int) bool {
	return today.Month() == time.December && today.Year() == year && today.Day() <= 25
}

// preparePuzzleFiles sets up the folder and files for a given day.
func preparePuzzleFiles(day int) error {
	dir := formatDayFolder(day)
	if err := ensureDirectoryExists(dir); err != nil {
		return err
	}

	if err := ensureFileExists(filepath.Join(dir, dir+".go"), func() error {
		return createGoPuzzleFileWithBoilerplate(day, filepath.Join(dir, dir+".go"))
	}); err != nil {
		return err
	}

	if err := ensureFileExists(filepath.Join(dir, "input.txt"), func() error {
		raw, err := GetInputHTTP(day)
		if err != nil {
			return err
		}
		return writeToFile(filepath.Join(dir, "input.txt"), raw)
	}); err != nil {
		return err
	}

	return nil
}

// ensureDirectoryExists creates a directory if it doesn't exist.
func ensureDirectoryExists(path string) error {
	return os.MkdirAll(path, 0755)
}

// ensureFileExists creates a file using the provided generator function if it doesn't exist.
func ensureFileExists(path string, generator func() error) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return generator()
		}
		return err
	}
	return nil
}

// formatDayFolder formats the folder name for a given day (e.g., "day01").
func formatDayFolder(day int) string {
	return fmt.Sprintf("day%02d", day)
}

// createGoPuzzleFileWithBoilerplate generates a Go file with boilerplate code.
func createGoPuzzleFileWithBoilerplate(day int, path string) error {
	content := fmt.Sprintf(`package main

import (
	"fmt"
	"advent_of_code_2024/utils/std"
	"advent_of_code_2024/utils/advent_of_code"
)

func main() {
	input, err := aoc.GetInputFile(%d)
	std.Error(err, 1)

	fmt.Println("Day %d - Part 1:")
	fmt.Println("Day %d - Part 2:")
}
`, day, day, day)
	return writeToFile(path, content)
}

// writeToFile writes a string to a file.
func writeToFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// GetInputHTTP retrieves the input for a specific day from the Advent of Code website.
func GetInputHTTP(day int) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: os.Getenv("SESSION")})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

// Reads file from respective day
func GetInputFile(day string) (string, error) {
	intDay, err := strconv.Atoi(day)
	if err != nil {
		return "", err
	}
	inputFileName := fmt.Sprintf("day%02d/input.txt", intDay)
	file, err := os.ReadFile(inputFileName)
	return string(file), err
}
