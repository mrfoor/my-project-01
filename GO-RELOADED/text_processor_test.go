
package main

import (
	"os"
	"testing"
)

// Tests for text processing functions

func TestReplaceHex(t *testing.T) {
	input := "Add 1F (hex) to the count."
	expected := "Add 31 to the count."
	result := ReplaceHex(input)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestReplaceBin(t *testing.T) {
	input := "It is 1010 (bin)."
	expected := "It is 10."
	result := ReplaceBin(input)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestFixArticles(t *testing.T) {
	input := "A apple is tasty."
	expected := "An apple is tasty."
	result := FixArticles(input)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestFixPunctuation(t *testing.T) {
	input := "Hello , world !"
	expected := "Hello, world!"
	result := FixPunctuation(input)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

// Tests for file utility functions

func TestFileUtils(t *testing.T) {
	testFile := "test.txt"
	testContent := "Hello, Go!"

	// Test WriteFile
	err := WriteFile(testFile, testContent)
	if err != nil {
		t.Errorf("WriteFile failed: %v", err)
	}

	// Test ReadFile
	content, err := ReadFile(testFile)
	if err != nil {
		t.Errorf("ReadFile failed: %v", err)
	}

	// Check content
	if content != testContent {
		t.Errorf("Expected '%s', got '%s'", testContent, content)
	}

	// Cleanup
	os.Remove(testFile)
}
