package data

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const filePath = "data.txt"

const (
	AssetPath = iota
	HighScore
	Score
)

func GetAsUint32(line uint32) uint32 {
	ui := Get(line)
	pui, err := strconv.ParseUint(ui, 10, 32)
	if err != nil {
		log.Printf("Error parsing high score: %v", err)
	}
	return uint32(pui)
}

// Get retrieves the value from a specific line in the file
func Get(line uint32) string {
	// Read the file
	contentBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	// Split the content into lines
	lines := strings.Split(string(contentBytes), "\n")

	// Ensure the line number is valid
	if int(line) >= len(lines) {
		fmt.Printf("Line %d does not exist in the file.\n", line)
	}

	// Return the value from the specified line
	return lines[line]
}

// Set updates a specific line in the file with the given value
func Set(line uint32, value interface{}) error {
	// Read the current content of the file
	contentBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// Split the content into lines
	lines := strings.Split(string(contentBytes), "\n")

	// Ensure the line number is valid
	if int(line) >= len(lines) {
		return fmt.Errorf("line %d does not exist in the file", line)
	}

	// Replace the specified line with the new value
	lines[line] = fmt.Sprint(value)

	// Join the lines back into a single string
	newContent := strings.Join(lines, "\n")

	// Write the updated content back to the file
	err = ioutil.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
