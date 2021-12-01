package main

import (
	"bufio"
	"os"
	"strconv"
)

func readInts(path string) (lines []int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err == nil {
			lines = append(lines, number)
		}
	}

	return lines, scanner.Err()
}
