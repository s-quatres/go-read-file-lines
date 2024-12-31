package goreadfilelines

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "example.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write test data to the file
	content := "Line one\nLine two\nLine three\n"
	if _, err := tempFile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// Test the ReadFile function
	lines, err := ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	// Verify the result
	expected := []string{"Line one", "Line two", "Line three"}
	if len(lines) != len(expected) {
		t.Fatalf("Expected %d lines, got %d", len(expected), len(lines))
	}
	for i, line := range expected {
		if lines[i] != line {
			t.Errorf("Expected line %d to be %q, got %q", i+1, line, lines[i])
		}
	}
}

func TestReadFile_EmptyFile(t *testing.T) {
	// Create an empty temporary file
	tempFile, err := os.CreateTemp("", "empty.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	// Test the ReadFile function on an empty file
	lines, err := ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Error reading file: %v", err)
	}

	// Verify the result
	if len(lines) != 0 {
		t.Fatalf("Expected 0 lines, got %d", len(lines))
	}
}

func TestReadFile_NonExistentFile(t *testing.T) {
	// Test the ReadFile function with a non-existent file
	_, err := ReadFile("nonexistent.txt")
	if err == nil {
		t.Fatalf("Expected an error when reading a non-existent file, got nil")
	}
}