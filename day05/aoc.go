package main

import (
	"fmt"
	"os"
)

type Line struct {
	points []Point
	x1, x2, y1, y2 int
}

type Point struct {
	x, y int
}

func parseLines(input string) Line {
	line := Line{
		points: []Point{},
	}

	fmt.Sscanf(input, "%d,%d -> %d,%d", &line.x1, &line.y1, &line.x2, &line.y2)

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
		out:
		for {
			y:=line.y1
			for {
					point := Point{x: x, y: y}			
					line.points = append(line.points, point)

				if y == line.y2 && x == line.x2 {
					break out
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

	pane := map[Point]int{}
	for _, line := range lines {
		if line.x1 == line.x2 || line.y1 == line.y2 {
			for _, p := range line.points {
				pane[p] += 1
			}
		}
	}

	return getSum(pane)
}

func getSolutionPart2(input []string) int {
	lines := []Line{}
	for _, inp := range input{
		line := parseLines(inp)
		lines = append(lines, line)
	}

	pane := map[Point]int{}
	for _, line := range lines {
		for _, p := range line.points {
			pane[p] += 1
		}
	}
	return getSum(pane)
}

func getSum(pane map[Point]int) (sum int) {
	for _, x := range pane {	
			if x > 1 {
				sum++
			}
 	}
	 return
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
