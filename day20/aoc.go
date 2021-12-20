package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"strings"
)

//go:embed input.txt
var input string

type Point struct {
	x, y int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mapGrid(input string) map[Point]int {
	rows := strings.Split(input, "\n")
	grid := map[Point]int{} 
	for x, row := range rows {
		for y, char := range row {

			if char == '#' {
				grid[Point{x,y}] = 1
			} else {
				grid[Point{x,y}] = 0
			}
		}
	}
	return grid
}

func getMinMaxXY(grid map[Point]int) (minX, minY, maxX, maxY int) {
	minX, minY = math.MaxInt, math.MaxInt

	for g := range grid {
		maxX = max(maxX, g.x)
		maxY = max(maxY, g.y)
		minX = min(minX, g.x)
		minY = min(minY, g.y)
	}
	return
}


func printGrid(grid map[Point]int) string {
	minX, minY, maxX, maxY := getMinMaxXY(grid)
	ret := ""
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {	
			point := Point{x: x, y: y}
			if val, ok := grid[point]; ok {
				if val == 1 {
					ret += "#"
				} else {
					ret += "."
				}
			} else {
				ret += "."
			}
		}
		ret += "\n"
	}
	ret += "\n"
	return ret
}


func getBinaryCode(startX, startY int, grid map[Point]int, iteration int) int {
	var binary int
	for x := startX-1; x <= startX+1; x++ {
		for y := startY-1; y <= startY+1; y++ {
			binary = binary << 1
			if v, ok := grid[Point{x, y}]; ok {
				if v == 1 {
					binary |= 1
				}
			} else {
				if iteration%2 == 1 {
					binary |= 1
				}
			}
		}
	}
	
	return binary
}

func enhanceImage(grid map[Point]int, key string, iteration int) map[Point]int {
	minX, minY, maxX, maxY := getMinMaxXY(grid)
	newGrid := map[Point]int{}
	for x := minX-3; x <= maxX+3; x++ {
		for y := minY-3; y <= maxY+3; y++ {
			index := getBinaryCode(x,y,grid, iteration)
			k := key[index]
			if k == '#' {
				newGrid[Point{x,y}] = 1
			} else {
				newGrid[Point{x,y}] = 0
			}

		}
	}
	return newGrid
}


func count(in map[Point]int) (res int) {
	for _, i := range in {
		if i == 1 {
			res++
		}
	}
	return
}

func getSolutionPart1(input string) int {
	inp := strings.Split(input, "\n\n")
	key := strings.ReplaceAll(inp[0], "\n", "")
	grid := mapGrid(inp[1])
	grid = enhanceImage(grid, key, 0)
	grid = enhanceImage(grid, key, 1)

	return count(grid)
}

func getSolutionPart2(input string) int {

	inp := strings.Split(input, "\n\n")
	key := strings.ReplaceAll(inp[0], "\n", "")
	grid := mapGrid(inp[1])
	for i:=0;i<50;i++ {
		grid = enhanceImage(grid, key, i)
	}

	return count(grid)

}

func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}

}
