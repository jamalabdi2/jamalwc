package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func GetCwd() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Error getting current working directory: %w", err)
	}
	return cwd, nil
}

func ProcessStatistics(data []byte) (int, int, int) {
	ByteCounts := len(data)

	stringData := string(data)
	words := strings.Fields(stringData)
	WordCounts := len(words)
	NewLineCount := bytes.Count(data, []byte("\n"))

	return NewLineCount, WordCounts, ByteCounts
}

func ProcessReader(r io.Reader) ([]byte, error) {
	buffer := bufio.NewReader(r)

	data, err := io.ReadAll(buffer)

	if err != nil {
		return nil, fmt.Errorf("Error reading from the reader: %w ", err)
	}
	return data, nil
}

func ProcessFile(filename string, cwd string) (*WcResult, error) {
	filePath := path.Join(cwd, filename)

	results := new(WcResult)

	file, err := os.Open(filePath)
	if err != nil {
		return results, fmt.Errorf("Error opening file %q: %w", filename, err)
	}

	defer file.Close()

	byteData, err := ProcessReader(file)

	if err != nil {
		return results, fmt.Errorf("Error reading file %q: %w ", filename, err)
	}
	results.PackStats(byteData)

	results.Filename = filename
	results.Filepath = filePath
	return results, nil
}
