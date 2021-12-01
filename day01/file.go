package main

import (
	"bufio"
	"os"
	"strconv"
)

func readInts(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
			return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
			number, err := strconv.Atoi(scanner.Text())
			if err == nil {
				lines = append(lines, number)
			}
	}
	return lines, scanner.Err()
}