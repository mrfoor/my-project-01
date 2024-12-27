package main

import (
	"fmt"
	"os"
)

func ReadFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("Error reading file %s: %v", fileName, err)
	}
	return string(data), nil
}

func WriteFile(fileName, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("Error creating file %s: %v", fileName, err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return fmt.Errorf("Error writing to file %s: %v", fileName, err)
	}
	return nil
}
