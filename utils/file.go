package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ReadIntsLines reads all lines in file and split each of them by sep. Expect lines with Ints
func ReadIntsLines(path, sep string) (ints [][]int, err error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		var intLine []int
		for _, num := range strings.Split(line, sep) {
			n, _ := strconv.Atoi(num)
			intLine = append(intLine, n)
		}
		ints = append(ints, intLine)
	}

	return ints, nil
}

// ReadIntsOneLine reads one line of Ints, separated by coma
func ReadIntsOneLine(path string) (ints []int, err error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	for _, num := range strings.Split(lines[0], ",") {
		n, _ := strconv.Atoi(num)
		ints = append(ints, n)
	}
	return ints, nil
}

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
