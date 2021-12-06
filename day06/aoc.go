package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcLanternFishes(initialNums []string, iterations int) int {
	
	lanternfishes := map[int]int{}
	for _, num := range initialNums {
		n, err := strconv.Atoi(num)
		if err != nil {
			return 0
		}
		lanternfishes[n]++
	}

	for i:=0;i<iterations;i++ {
		newLanternfishes := map[int]int{}
		for k, f := range lanternfishes {
			k--
			if k == -1 {
				k = 6
				newLanternfishes[8] += f
			}
			newLanternfishes[k] += f
		}
		lanternfishes = newLanternfishes
	}

	sum := 0
	for _, f := range lanternfishes {
		sum += f
	}

	return sum
}

func getSolutionPart1(input string) int{
	return calcLanternFishes(strings.Split(input, ","), 80)
}

func getSolutionPart2(input string) int {
	return calcLanternFishes(strings.Split(input, ","), 256)
}

func main() {
	if input, err := readInput("input.txt"); err != nil {
		panic("could not read file")
	} else {
		if os.Getenv("part") == "part2" {
			fmt.Println(getSolutionPart2(input))
		} else {
			fmt.Println(getSolutionPart1(input))
		}
	}
}
