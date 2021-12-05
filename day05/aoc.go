package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	points []Point
	x1 int
	x2 int
	y1 int
	y2 int
}

type Point struct {
	x int
	y int
}

func parseLines(input string) Line {
	startStop := strings.Split(input, " -> ")
	x1y1 := strings.Split(startStop[0], ",")	
	x2y2 := strings.Split(startStop[1], ",")

	line := Line{
		points: []Point{},
	}

	line.x1, _ = strconv.Atoi(x1y1[0])
	line.y1, _ = strconv.Atoi(x1y1[1])
	line.x2, _ = strconv.Atoi(x2y2[0])
	line.y2, _ = strconv.Atoi(x2y2[1])

	// straight line
	if line.x1 == line.x2 || line.y1 == line.y2 {
		x:=line.x1
		for {
			y:=line.y1
			for {
				point := Point{x: x, y: y}			
				line.points = append(line.points, point)
				if y == line.y2 {
					break
				}

				if line.y2 >= line.y1 {
					y++
				} else {
					y--
				}
				
			}

			if x == line.x2 {
				break
			}

			if line.x2 >= line.x1 {
				x++
			} else {
				x--
			}
			
		}
	} else {
		// diagonal line
		x:=line.x1
		for {
			y:=line.y1
			for {
					point := Point{x: x, y: y}			
					line.points = append(line.points, point)

				if y == line.y2 {
					break
				}

				if line.y2 >= line.y1 {
					y++
				} else {
					y--
				}

				if line.x2 >= line.x1 {
					x++
				} else {
					x--
				}	
			}

			if x == line.x2 {
				break
			}
			
		}
	}


	return line
}

func getSolutionPart1(input []string) int{
	lines := []Line{}
	for _, inp := range input{
		line := parseLines(inp)
		lines = append(lines, line)
	}

	pane := map[int]map[int]int{}
	for _, line := range lines {
		if line.x1 == line.x2 || line.y1 == line.y2 {
			for _, p := range line.points {
				if len(pane[p.x]) == 0 {
					pane[p.x] = map[int]int{}
				}
				pane[p.x][p.y] += 1
			}
		}
	}

	sum := 0
	for _, x := range pane {
		for _, y := range x {
			if y > 1 {
				sum++
			}
		}
 	}


	return sum
}

func getSolutionPart2(input []string) int {
	lines := []Line{}
	for _, inp := range input{
		line := parseLines(inp)
		lines = append(lines, line)
	}

	pane := map[int]map[int]int{}
	for _, line := range lines {
		for _, p := range line.points {
			if len(pane[p.x]) == 0 {
				pane[p.x] = map[int]int{}
			}
			pane[p.x][p.y] += 1
		}
	}

	sum := 0
	for _, x := range pane {
		for _, y := range x {
			if y > 1 {
				sum++
			}
		}
 	}


	return sum
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
