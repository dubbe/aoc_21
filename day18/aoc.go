package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic("cannot convert")
	}
	return i
}

func findFirstToExplode(snailfish string) (bool, []int, string) {
	re := regexp.MustCompile(`\[\d+,\d+\]`)
	submatchesIndex := re.FindAllStringIndex(snailfish, -1)

	for _, submatch := range submatchesIndex {
		before := snailfish[:submatch[0]]
		d1 := strings.Count(before, "[")
		d2 := strings.Count(before, "]")
		if d1-d2 == 4 {
			return true, submatch, snailfish[submatch[0]:submatch[1]]
		}
	}

	return false, []int{}, ""
}

func findFirstToSplit(snailfish string) (bool, []int, int) {
	re := regexp.MustCompile(`\d{2,}`)
	submatchesIndex := re.FindAllStringIndex(snailfish, 1)

	if len(submatchesIndex) > 0 {
		return true, submatchesIndex[0], atoi(snailfish[submatchesIndex[0][0]:submatchesIndex[0][1]])
	}

	return false, []int{}, 0
}

func replaceLast(str string, repl int) string {
	re := regexp.MustCompile(`\d+`)
	submatchesIndex := re.FindAllStringIndex(str, -1)
	submatches := re.FindAllString(str, -1)
	if len(submatches) <= 0 {
		return str
	}
	i := len(submatchesIndex)

	intRepl, _ := strconv.Atoi(submatches[i-1])
	str = fmt.Sprintf("%s%d%s", str[:submatchesIndex[i-1][0]], intRepl+repl, str[submatchesIndex[i-1][1]:])

	return str
}

func replaceFirst(str string, repl int) string {
	re := regexp.MustCompile(`\d+`)
	submatchesIndex := re.FindAllStringIndex(str, 1)
	submatches := re.FindAllString(str, 1)
	if len(submatches) <= 0 {
		return str
	}
	intRepl, _ := strconv.Atoi(submatches[0])
	str = fmt.Sprintf("%s%d%s", str[:submatchesIndex[0][0]], intRepl+repl, str[submatchesIndex[0][1]:])

	return str
}

func getXY(snail string) (x int, y int) {
	fmt.Sscanf(snail, "[%d,%d]", &x, &y)
	return x, y
}

func explode(snailfish string) (string, bool) {
	found, i, snail := findFirstToExplode(snailfish)

	if !found {
		return snailfish, false
	}
	x, y := getXY(snail)

	right := replaceFirst(snailfish[i[1]:], y)
	left := replaceLast(snailfish[:i[0]], x)
	str := fmt.Sprintf("%s0%s", left, right)

	return str, true
}

func split(snailfish string) (string, bool) {
	found, i, n := findFirstToSplit(snailfish)
	if !found {
		return snailfish, false
	}

	half := float64(n) / float64(2)
	x := int(math.Floor(half))
	y := int(math.Ceil(half))

	str := fmt.Sprintf("%s[%d,%d]%s", snailfish[:i[0]], x, y, snailfish[i[1]:])

	return str, true
}

func addition(snail1, snail2 string) string {
	return fmt.Sprintf("[%s,%s]", snail1, snail2)
}

func doStuff(snail string) string {
	var couldExplode, couldSplit bool

	for {
		snail, couldExplode = explode(snail)

		if !couldExplode {
			snail, couldSplit = split(snail)
			if !couldSplit {
				return snail
			}
		}
	}
}

func calculateMagnitude(snail string) int {

	re := regexp.MustCompile(`\[\d+,\d+\]`)

	for {
		submatches := re.FindAllString(snail, -1)
		if len(submatches) == 0 {
			break
		}
		for _, match := range submatches {
			x, y := getXY(match)
			sum := (x * 3) + (y * 2)
			snail = strings.Replace(snail, match, fmt.Sprint(sum), 1)
		}
	}
	sum, _ := strconv.Atoi(snail)
	return sum
}

func getSolutionPart1(input string) int {
	rows := strings.Split(input, "\n")
	currentNode := rows[0]
	for i := range rows {
		if i == 0 {
			continue
		}
		added := addition(currentNode, rows[i])
		currentNode = doStuff(added)
	}

	return calculateMagnitude(currentNode)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getSolutionPart2(input string) int {
	rows := strings.Split(input, "\n")
	max := 0
	for r, row := range rows {
		for i := 0; i < len(rows); i++ {
			if i != r {

				added := doStuff(addition(row, rows[i]))
				magnitude := calculateMagnitude(added)
				max = getMax(max, magnitude)
			}
		}
	}
	return max
}

func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}

}
