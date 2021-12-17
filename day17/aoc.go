package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed input.txt
var input string

func maxMin(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

func testTrajectory(x, y, x1, x2, y1, y2 int) (bool, int) {

	xPos, yPos := 0, 0
	// i := 0
	maxY := 0
	for {
		xPos += x
		yPos += y
		if x > 0 {
			x--
		}
		y--

		maxY, _ = maxMin(maxY, yPos)

		if xPos >= x1 && xPos <= x2 && yPos <= y1 && yPos >= y2 {

			return true, maxY
		}

		if (x == 0 && xPos < x1) || xPos > x2 || yPos < y2 {

			return false, 0
		}

		// // Failguard
		// i++
		// if i > 10000 {
		// 	return false, 0
		// }
	}
	return false, 0
}

func getSolutionPart1(input string) int {
	var sx1, sx2, sy1, sy2, x1, x2, y1, y2 int
	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &sx1, &sx2, &sy1, &sy2)
	x2, x1 = maxMin(sx1, sx2)
	y1, y2 = maxMin(sy1, sy2)

	maxY := 0
	for startX := 0; startX < 100; startX++ {
		for startY := 0; startY < 100; startY++ {
			_, max := testTrajectory(startX, startY, x1, x2, y1, y2)
			maxY, _ = maxMin(maxY, max)
		}
	}
	return maxY
}

func getSolutionPart2(input string) int {
	var sx1, sx2, sy1, sy2, x1, x2, y1, y2 int
	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d\n", &sx1, &sx2, &sy1, &sy2)
	
	x2, x1 = maxMin(sx1, sx2)
	y1, y2 = maxMin(sy1, sy2)
	
	hits := 0
	maxStarty := -1000
	for startX := 0; startX <= x2; startX++ {
		for startY := y2; startY < y2*-1; startY++ {
			found, _ := testTrajectory(startX, startY, x1, x2, y1, y2)
			if found {
				maxStarty, _ = maxMin(maxStarty, startY)
				hits++
			}
		}
	}

	return hits
}

func main() {

	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}

}
