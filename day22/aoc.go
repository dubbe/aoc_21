package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed input.txt
var input string

type Cube struct  {
	x,y,z int
}

type CubeTwo struct {
	x1,x2,y1,y2,z1,z2 int
	state bool
}

func reboot(state bool, x, y, z  []int, cubes map[Cube]bool) map[Cube]bool {
	for xx:=x[0];xx<=x[1];xx++ {
		for yy:=y[0];yy<=y[1];yy++ {
			for zz:=z[0];zz<=z[1];zz++ {
				if state {
					cubes[Cube{x: xx, y: yy, z: zz}] = true
				} else {
					delete(cubes, Cube{xx, yy, zz})
				}
				
			}
		}
	}
	return cubes
}

func getSolutionPart1(input string) int {
	inputs := strings.Split(input, "\n")
	cubes := map[Cube]bool{}
	for _, input := range inputs {
		var st string
		var x1, x2, y1, y2, z1, z2 int
		fmt.Sscanf(input, "%s x=%d..%d,y=%d..%d,z=%d..%d", &st, &x1, &x2, &y1, &y2, &z1, &z2)
		if x1 > 50 || x1 < -50 || x2 > 50 || x2 < -50 {
			continue;
		}
		cubes = reboot(st == "on", []int{x1,x2}, []int{y1,y2}, []int{z1,z2}, cubes)

	}
	return len(cubes)
}



func getSolutionPart2(input string) int64 {
	inputs := strings.Split(input, "\n")
	cubes := []CubeTwo{}
	for _, input := range inputs {
		var st string
		var x1, x2, y1, y2, z1, z2 int
		fmt.Sscanf(input, "%s x=%d..%d,y=%d..%d,z=%d..%d", &st, &x1, &x2, &y1, &y2, &z1, &z2)
		cubes = append(cubes, CubeTwo{x1:x1, x2:x2,y1:y1, y2:y2, z1: z1, z2: z2, state: st == "on"})

	}
	return int64(count(cubes))
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func intersection(cube, other CubeTwo) *CubeTwo {
	x1 := MaxInt(cube.x1, other.x1)
	x2 := MinInt(cube.x2, other.x2)

	if x2 < x1 {
		return nil
	}

	y1 := MaxInt(cube.y1, other.y1)
	y2 := MinInt(cube.y2, other.y2)

	if y2 < y1 {
		return nil
	}

	z1 := MaxInt(cube.z1, other.z1)
	z2 := MinInt(cube.z2, other.z2)

	if z2 < z1 {
		return nil
	}

	return &CubeTwo{
		x1: x1,
		x2: x2,
		y1: y1,
		y2: y2,
		z1: z1,
		z2: z2,
	}
}

func volume(cube CubeTwo) int {
	vol := ((cube.x2 + 1) - cube.x1) *
		((cube.y2 + 1) - cube.y1) *
		((cube.z2 + 1) - cube.z1)

	return vol
}

func count(cubes []CubeTwo) (sum int) {
	for i := len(cubes) - 1; i >= 0; i-- {
		cube := cubes[i]

		if !cube.state {
			continue
		}

		intersections := []CubeTwo{}

		for _, next := range cubes[i+1:] {
			intersection := intersection(cube, next)

			if intersection == nil {
				continue
			}

			shouldCountVolume := true
			intersection.state = shouldCountVolume

			intersections = append(intersections, *intersection)
		}

		sum += volume(cube)
		sum -= count(intersections)
	}

	return
}


func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}

}
