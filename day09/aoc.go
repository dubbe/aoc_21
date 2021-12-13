package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/fatih/color"
)

type Point struct {
	x, y int
}

var maxX = 0
var maxY = 0

func parsePane(input []string) map[Point]int {
	pane := map[Point]int{}
	for y, row := range input {
		for x, i := range row {
			height, err := strconv.Atoi(string(i))
			if err != nil {
				fmt.Printf("ERROR")
			}
			point := Point{x: x, y: y}
			pane[point] = height
			if x > maxX {
				maxX = x
			}
		}
		if y > maxY {
			maxY = y
		}
	}
	return pane
}

func checkIfLowest(point Point, height int, pane map[Point]int) bool {
	pointsToCheck := []Point{
		{x: point.x, y: point.y - 1},
		{x: point.x, y: point.y + 1},
		{x: point.x + 1, y: point.y},
		{x: point.x - 1, y: point.y},
	}

	heighest := false
	for _, p := range pointsToCheck {
		if p.x >= 0 && p.y >= 0 && p.x <= maxX && p.y <= maxY {
			if height >= pane[p] {
				heighest = true
			}
		}
	}

	return !heighest
}

func printPane(pane map[Point]int) {
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			point := Point{x: x, y: y}
			height := pane[point]
			if checkIfLowest(point, height, pane) {
				red := color.New(color.FgRed).SprintFunc()
				fmt.Printf("%s", red(height))
			} else {
				fmt.Printf("%d", height)
			}
		}
		fmt.Printf("\n")
	}
}

func contains(s []Point, e Point) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getPointsToCheck(point Point) []Point {
	return []Point{
		{x: point.x, y: point.y - 1},
		{x: point.x, y: point.y + 1},
		{x: point.x + 1, y: point.y},
		{x: point.x - 1, y: point.y},
	}
}

func getBasin(point Point, pane map[Point]int, foundPoints *[]Point) int {
	sum := 1
	*foundPoints = append(*foundPoints, point)
	for _, p := range getPointsToCheck(point) {
		if p.x >= 0 && p.y >= 0 && p.x <= maxX && p.y <= maxY && !contains(*foundPoints, p) {
			if pane[p] != 9 {
				sum += getBasin(p, pane, foundPoints)
			}
		}
	}
	return sum
}

func getSolutionPart1(input []string) int {
	sum := 0
	pane := parsePane(input)

	printPane(pane)

	for p, height := range pane {
		if checkIfLowest(p, height, pane) {
			sum += 1 + height
		}
	}

	return sum
}

func getSolutionPart2(input []string) int {
	pane := parsePane(input)

	basins := []int{}

	for p, height := range pane {
		if checkIfLowest(p, height, pane) {
			foundPoints := []Point{}
			b := getBasin(p, pane, &foundPoints)
			basins = append(basins, b)
		}
	}
	sort.Ints(basins)

	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]

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
