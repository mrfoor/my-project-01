package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read the input file
	content, err := ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Process the text
	processedContent := ProcessText(content)

	// Write the processed content to the output file
	err = WriteFile(outputFile, processedContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Le texte a été traité avec succès et sauvegardé dans", outputFile)
}
