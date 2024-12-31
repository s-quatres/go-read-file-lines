package goreadfilelines

import (
	"bufio"
	"os"
)

// ReadFile reads a file and returns its lines as a slice of strings.
func ReadFile(filePath string) ([]string, error) {
    var lines []string

    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return lines, nil
}

