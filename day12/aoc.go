package main

import (
	"fmt"
	"os"
	"strings"
)

func parseAdjacent(input []string) map[string][]string {
	adjacentMap := map[string][]string{}
	for _, i := range input {
		p := strings.Split(i, "-")
		
		if adjacentMap[p[0]] == nil {
			adjacentMap[p[0]] = []string{}
		}

		if adjacentMap[p[1]] == nil {
			adjacentMap[p[1]] = []string{}
		}

		adjacentMap[p[0]] = append(adjacentMap[p[0]], p[1])
		adjacentMap[p[1]] = append(adjacentMap[p[1]], p[0])
	}
	return adjacentMap
}

func contains(s []string, e string) bool {
	for _, a := range s {
			if a == e {
					return true
			}
	}
	return false
}

func findPath(adjacentMap map[string][]string, start string, visited []string, road []string, paths *int)  {

	if strings.ToLower(start) == start {
		visited = append(visited, start)
	}
	road = append(road, start)

	for _, a := range adjacentMap[start] {
		if a == "end" {
			*paths++
			road = append(road, a)
		} else if !contains(visited, a) {
			findPath(adjacentMap, a, visited, road, paths)
		}
	}
}

func findPath2(adjacentMap map[string][]string, start string, visited []string, visitedTwice bool, road []string, paths *int)  {

	if strings.ToLower(start) == start {
		visited = append(visited, start)
	}
	road = append(road, start)

	for _, a := range adjacentMap[start] {
		if a == "end" {
			*paths++
			road = append(road, a)
		} else if !contains(visited, a)  {
			findPath2(adjacentMap, a, visited, visitedTwice, road, paths)
		} else if contains(visited, a) && !visitedTwice && a != "start" {
			findPath2(adjacentMap, a, visited, true, road, paths)
		}
	}
}

func getSolutionPart1(input []string) int{
	adjacentMap := parseAdjacent(input)

	visitedOnce := []string{}
	
	road := []string{}
	sum := 0
	findPath(adjacentMap, "start", visitedOnce, road, &sum)
	return sum
}

func getSolutionPart2(input []string) int {
	adjacentMap := parseAdjacent(input)

	visitedOnce := []string{"start"}
	visitedTwice := false
	road := []string{}
	sum := 0
	findPath2(adjacentMap, "start", visitedOnce, visitedTwice, road, &sum)
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
