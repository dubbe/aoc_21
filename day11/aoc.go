package main

import (
	"fmt"
	"os"
	"strconv"

	tm "github.com/buger/goterm"
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

func sprintPane(pane map[Point]int) string {
	str := ""
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			point := Point{x: x, y: y}
			i := pane[point]
			if i == 0 {
				red := color.New(color.Bold).SprintFunc()
				str += fmt.Sprintf("%s", red(i))
			} else {
				str += fmt.Sprintf("%d", i)
			}
		}
		str += fmt.Sprintf("\n")
	}
	str += fmt.Sprintf("\n")
	return str
}

func doNextStep(pane map[Point]int) {
	for point := range pane {
		pane[point]++
	}
}

func explode(pane map[Point]int, point Point) {
	pointsToCheck := []Point{
		{x: point.x, y: point.y - 1},
		{x: point.x, y: point.y + 1},
		{x: point.x + 1, y: point.y},
		{x: point.x - 1, y: point.y},
		{x: point.x - 1, y: point.y - 1},
		{x: point.x - 1, y: point.y + 1},
		{x: point.x + 1, y: point.y - 1},
		{x: point.x + 1, y: point.y + 1},
	}

	for _, p := range pointsToCheck {
		if pane[p] != 0 {
			pane[p]++
		}
	}
}

func doExplodingStuff(pane map[Point]int, startingSum int) int {
	sum := 0
	for point := range pane {
		if pane[point] > 9 {
			pane[point] = 0
			explode(pane, point)
			sum++
		}
	}
	if sum > 0 {
		sum = doExplodingStuff(pane, sum)
	}
	return startingSum + sum
}

func getSolutionPart1(input []string) int {
	tm.Clear() // Clear current screen
	sum := 0
	pane := parsePane(input)
	for i := 0; i < 100; i++ {
		doNextStep(pane)
		sum += doExplodingStuff(pane, 0)
		// time.Sleep(100 * time.Millisecond)
		// printPane(pane)
	}
	return sum
}

func printPane(pane map[Point]int) {
	tm.MoveCursor(1, 1)
	tm.Printf(sprintPane(pane))
	tm.Flush() // Call it every time at the end of rendering
}

func checkIfAllIsExploded(pane map[Point]int) bool {
	exploded := true
	for _, i := range pane {
		if i != 0 {
			exploded = false
			break
		}
	}
	return exploded
}

func getSolutionPart2(input []string) int {
	pane := parsePane(input)
	i := 0
	for {
		doNextStep(pane)
		_ = doExplodingStuff(pane, 0)
		if checkIfAllIsExploded(pane) {
			return i + 1
		}
		i++
	}
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
