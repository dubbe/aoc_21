package main

import (
	_ "embed"
	"fmt"
	"reflect"
	"strings"
)

//go:embed input.txt
var input string
var maxX, maxY int
type Point struct{
	x,y int
}

func parseGrid(input string) (map[Point]bool, map[Point]bool) {
	maxX, maxY = 0, 0
	rows := strings.Split(input, "\n")
	gridEast := map[Point]bool{}
	gridSouth := map[Point]bool{}
	for x, row := range rows {
		for y, point := range row {
			maxX = max(maxX, x)
			maxY = max(maxY, y)
			if point == '>' {
				gridEast[Point{x,y}] = true
			} else if point == 'v' {
				gridSouth[Point{x,y}] = true
			}
		}
	}
	return gridEast, gridSouth
}

func moveEast(gridEast, gridSouth map[Point]bool) map[Point]bool {
	newGridEast := map[Point]bool{}
	
	for point, _ := range gridEast {
		newPointY := point.y + 1
		if newPointY > maxY {
			newPointY = 0
		}
		if gridEast[Point{point.x, newPointY}] || gridSouth[Point{point.x, newPointY}] {
			newGridEast[point] = true
			continue
		} else {
			newGridEast[Point{point.x, newPointY}] = true
		}
	}
	return newGridEast
}

func moveSouth(gridEast, gridSouth map[Point]bool) map[Point]bool {
	newGridSouth := map[Point]bool{}
	
	for point, _ := range gridSouth {
		newPointX := point.x + 1
		if newPointX > maxX {
			newPointX = 0
		}
		if gridEast[Point{newPointX, point.y}] || gridSouth[Point{newPointX, point.y}] {
			newGridSouth[point] = true
			continue
		} else {
			newGridSouth[Point{newPointX, point.y}] = true
		}
	}
	return newGridSouth
}

func move(gridEast, gridSouth map[Point]bool) (map[Point]bool, map[Point]bool) {
	gridEast = moveEast(gridEast, gridSouth)
	gridSouth = moveSouth(gridEast, gridSouth)
	return gridEast, gridSouth
}


func getSolutionPart1(input string) int {
	gridEast, gridSouth := parseGrid(input)
	i:=0
	prevGridEast, prevGridSouth := gridEast, gridSouth
	for {
		gridEast, gridSouth = move(gridEast, gridSouth)
		i++

		if reflect.DeepEqual(gridEast, prevGridEast) && reflect.DeepEqual(gridSouth, prevGridSouth) {
			break
		} else {
			prevGridEast, prevGridSouth = gridEast, gridSouth
		}
	}
	return i
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(getSolutionPart1(input))
}
