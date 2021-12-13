package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func strToIntArr(s string) []int {
	strs := strings.Split(s, ",")
	res := []int{}
	for _, str := range strs {
		x, _ := strconv.Atoi(str)
		res = append(res, x)
	}
	return res
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func calcMean(numbers []int) (int, int) {
	mean := float64(sum(numbers)) / float64(len(numbers))
	return int(math.Floor(mean)), int(math.Ceil(mean))
}

func calcMedian(numbers []int) int {
	mNumber := len(numbers) / 2
	return numbers[mNumber]
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func getSolutionPart1(input string) int {

	numbers := strToIntArr(input)
	sort.Ints(numbers)
	median := calcMedian(numbers)

	sum := 0
	for _, num := range numbers {
		sum += diff(num, median)
	}

	return sum
}

func calculateFuelPart2(num []int, median int) int {
	fuel := 0
	for _, x := range num {
		n := diff(median, x)
		cost := (n * (n + 1)) / 2
		fuel += cost
	}
	return fuel
}

func getSolutionPart2(input string) int {
	numbers := strToIntArr(input)
	sort.Ints(numbers)
	meanFloor, meanCeil := calcMean(numbers)

	startFuel := calculateFuelPart2(numbers, meanFloor)
	starFuelCeil := calculateFuelPart2(numbers, meanCeil)
	if starFuelCeil < startFuel {
		startFuel = starFuelCeil
	}
	return startFuel

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
