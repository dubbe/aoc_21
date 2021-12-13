package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Instruction struct {
	direction string
	value     int
}

var maxX = 0
var maxY = 0

func parsePane(input []string) (pane map[Point]string, instructions []Instruction) {
	pane = map[Point]string{}
	instructions = []Instruction{}
	parseInstructions := false
	for _, row := range input {
		if row == "" {
			parseInstructions = true
			continue
		}

		if !parseInstructions {
			p := strings.Split(row, ",")
			x, _ := strconv.Atoi(p[0])
			y, _ := strconv.Atoi(p[1])
			point := Point{x: x, y: y}
			pane[point] = "#"

			maxX = findMax(maxX, x)
			maxY = findMax(maxY, y)

		} else {

			var parse string
			fmt.Sscanf(row, "fold along %s", &parse)
			splittedParse := strings.Split(parse, "=")
			direction := splittedParse[0]
			value, _ := strconv.Atoi(splittedParse[1])
			instructions = append(instructions, Instruction{direction: direction, value: value})
		}
	}

	return pane, instructions
}

func findMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func foldUp(pane map[Point]string, fold int) map[Point]string {
	newPane := map[Point]string{}
	for y := 0; y < fold; y++ {
		for p := range pane {
			if p.y == y || p.y == fold+fold-y {
				newPane[Point{p.x, y}] = "#"
			}
		}
	}

	maxY = fold - 1

	return newPane
}

func foldLeft(pane map[Point]string, fold int) map[Point]string {
	newPane := map[Point]string{}
	for x := 0; x < fold; x++ {
		for p := range pane {
			if p.x == x || p.x == fold+fold-x {
				newPane[Point{x, p.y}] = "#"
			}
		}
	}

	maxX = fold - 1

	return newPane
}

func fold(pane map[Point]string, where int, direction string) map[Point]string {
	switch direction {
	case "y":
		return foldUp(pane, where)
	case "x":
		return foldLeft(pane, where)
	}
	return pane
}

func printPane(pane map[Point]string) string {
	ret := ""
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {

			point := Point{x: x, y: y}
			if val, ok := pane[point]; ok {
				ret += fmt.Sprintf("%s", val)
			} else {
				ret += fmt.Sprintf(".")
			}
		}
		ret += fmt.Sprintf("\n")
	}
	ret += fmt.Sprintf("\n")
	return ret
}

func calculatePoints(pane map[Point]string) int {
	sum := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			point := Point{x: x, y: y}
			if _, ok := pane[point]; ok {
				sum++
			}
		}
	}
	return sum
}

func getSolutionPart1(input []string) int {
	pane, instructions := parsePane(input)
	pane = fold(pane, instructions[0].value, instructions[0].direction)
	sum := calculatePoints(pane)
	return sum
}

func getSolutionPart2(input []string) string {
	pane, instructions := parsePane(input)
	for i := range instructions {
		pane = fold(pane, instructions[i].value, instructions[i].direction)
	}
	return printPane(pane)
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
