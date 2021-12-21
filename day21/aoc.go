package main

import (
	"container/ring"
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed input.txt
var input string


func getSolutionPart1(input string) int {
	inputs := strings.Split(input, "\n")
	var playerOneStartingPosition, playerTwoStartingPosition int
	fmt.Sscanf(inputs[0], "Player 1 starting position: %d", &playerOneStartingPosition)
	fmt.Sscanf(inputs[1], "Player 2 starting position: %d", &playerTwoStartingPosition)

	ints := []int{1,2,3,4,5,6,7,8,9,10}
	board1 := ring.New(len(ints))
	board2 := ring.New(len(ints))
	for i := 0; i < board1.Len(); i++ {
		board1.Value = ints[i]
		board1 = board1.Next()

		board2.Value = ints[i]
		board2 = board2.Next()
 	}

	board1 = board1.Move(playerOneStartingPosition-1)
	board2 = board2.Move(playerTwoStartingPosition-1)

	playerOnePoints := 0
	playerTwoPoints := 0

	turn := 0;
	i := 1
	for {
		rolls := []int{i, i+1, i+2}
		move := rolls[0] + rolls[1] + rolls[2]
		turn++
		i+=3
		if turn%2 == 1 {
			board1 = board1.Move(move)
			playerOnePoints += board1.Value.(int)
			if playerOnePoints >= 1000 {
				return (i-1)*playerTwoPoints
			}
		} else {
			board2 = board2.Move(move)
			playerTwoPoints += board2.Value.(int)
			if playerTwoPoints >= 1000 {
				return (i-1)*playerOnePoints
			}
		}	
	}
}
 
func checkCache(pos1, pos2, score1, score2, player int) (bool, [2]int64) {
	if v, ok := cache[[5]int{pos1, pos2, score1, score2, player}]; ok {
		return true, v
	}
	return false, [2]int64{0,0}
}
	
func play(position [2]int, score [2]int, player int) (int64, int64) {
	if found, v := checkCache(position[0], position[1], score[0], score[1], player); found {
		return v[0], v[1]
	}

	if score[0] >= 21 {
		return 1, 0
	} else if score[1] >= 21 {
		return 0, 1
	}

	var wins [2]int64
	for _, dice1 := range []int{1, 2, 3} {
		for _, dice2 := range []int{1, 2, 3} {
			for _, dice3 := range []int{1, 2, 3} {
				newPos := position
				newPos[player] = (position[player] + dice1 + dice2 + dice3) % 10
				newScore := score
				newScore[player] = score[player] + newPos[player] + 1
				nextPlayer := 1
				if player == 1 {
					nextPlayer = 0
				}
				newWinOne, newWinTwo := play(newPos, newScore, nextPlayer)
				wins[0] += newWinOne
				wins[1] += newWinTwo
			}
		}
	}
	cache[[5]int{position[0], position[1], score[0], score[1], player}] = wins
	return wins[0], wins[1]
}

var cache map[[5]int][2]int64

func getSolutionPart2(input string) int64 {
	inputs := strings.Split(input, "\n")
	var playerOneStartingPosition, playerTwoStartingPosition int
	fmt.Sscanf(inputs[0], "Player 1 starting position: %d", &playerOneStartingPosition)
	fmt.Sscanf(inputs[1], "Player 2 starting position: %d", &playerTwoStartingPosition)
	positions := [2]int{playerOneStartingPosition-1, playerTwoStartingPosition-1}
	cache = make(map[[5]int][2]int64)
	return max(play(positions, [2]int{0, 0}, 0))
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func main() {
	if os.Getenv("part") == "part2" {
		fmt.Println(getSolutionPart2(input))
	} else {
		fmt.Println(getSolutionPart1(input))
	}

}
