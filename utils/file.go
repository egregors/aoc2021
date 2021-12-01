package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
)

// ReadInts reads text line by line and convert each of them into Int
func ReadInts(path string) (ints []int, err error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	for _, v := range lines {
		n, err := strconv.Atoi(v) //nolint:govet // it's ok
		if err != nil {
			return nil, err
		}
		ints = append(ints, n)
	}
	return ints, err
}

// ReadLines reads text input from the file line by line
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) { //nolint:gosec // meh
		_ = file.Close()
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
