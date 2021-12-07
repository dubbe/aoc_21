package main

import (
	"bufio"
	"os"
)

func readInput(path string) (lines string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = scanner.Text()
	}

	return lines, scanner.Err()
}
