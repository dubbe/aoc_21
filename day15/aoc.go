package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

type Cell struct {
	x, y int
	val  uint16
}
type Heap []Cell

func (h Heap) Len() int { return len(h) }
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h Heap) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}
func (h *Heap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}



var maxY, maxX int

func parseGrid(input []string) map[Point]uint16 {
	maxY = len(input)
	maxX = len(input[0])
	grid := map[Point]uint16{}

	for y, row := range input {
		for x, risk := range row {
			r, _ := strconv.Atoi(string(risk))
			grid[Point{y, x}] = uint16(r)
		}
	}
	return grid

}

func printGrid(grid map[Point]uint16) {

	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {

			point := Point{x: x, y: y}
			if val, ok := grid[point]; ok {
				fmt.Printf("%d", val)
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func extendGrid(grid map[Point]uint16) map[Point]uint16 {
	newGrid := grid
	thisMaxY := maxY

	for i:=0;i<4;i++ {
		for x := 0; x < maxX; x++ {
			for y := 0; y < thisMaxY; y++ {
				v := newGrid[Point{x, (y+(thisMaxY*i))}]
			if v == 9 {
				v = 1
			} else {
				v++
			}

			newY := y+(thisMaxY*(i+1))
			newPoint := Point{x:x, y:newY}
			newGrid[newPoint] = v
			if newY > maxY {
				maxY = newY+1
			}
		}
		}
	}

	thisMaxX := maxX
	for i:=0;i<4;i++ {
		for x := 0; x < thisMaxX; x++ {
			for y := 0; y < maxY; y++ {
				v := newGrid[Point{(x+(thisMaxX*i)), y}]
			if v == 9 {
				v = 1
			} else {
				v++
			}

			newX := x+(thisMaxX*(i+1))
			newPoint := Point{x:newX, y:y}
			newGrid[newPoint] = v
			if newX > maxX {
				maxX = newX+1
			}
		}
		}
	}
	return newGrid
}

func getShortest(grid map[Point]uint16) int {
	seen := map[Point]bool{}
	dist := map[Point]uint16{}

	for k, _ := range grid {
		dist[k] = math.MaxUint16
	}

	// Check out of bounds
	ok := func(y, x int) bool {
		return y >= 0 && y < maxY && x >= 0 && x < maxX
	}

	// make heap
	h := make(Heap, 1, 100)
	h[0] = Cell{0, 0, 0}

	for {
		hp := heap.Pop(&h).(Cell)
		seen[Point{x: hp.x, y: hp.y}] = true
		if hp.x == maxY-1 && hp.y == maxX-1 {
			return int(hp.val)
		}
		for _, nei := range [][2]int{
			{hp.x + 1, hp.y}, {hp.x - 1, hp.y}, {hp.x, hp.y - 1}, {hp.x, hp.y + 1},
		} {
			x, y := nei[0], nei[1]
			if !ok(x, y) || seen[Point{x, y}] {
				continue
			}
			risk := hp.val + grid[Point{x, y}]
			if risk >= dist[Point{x, y}] {
				continue
			}
			dist[Point{x, y}] = risk
			heap.Push(&h, Cell{x, y, risk})
		}
	}
}

func getSolutionPart1(input []string) int {
	grid := parseGrid(input)
	return getShortest(grid)
}

func getSolutionPart2(input []string) int {
	grid := parseGrid(input)
	grid = extendGrid(grid)
	return getShortest(grid)
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
