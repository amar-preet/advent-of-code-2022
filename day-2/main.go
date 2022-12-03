package main

import (
	"fmt"
	"strings"

	"github.com/amar-preet/advent-of-code-2022/reader"
)

func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	//A, X = Rock = 1
	//B, Y = paper = 2
	//C, Z = Scissors = 3

	//X = Rock
	//Y = Paper
	//Z = Scissors

	// win = 6, loss = 0, draw = 3
	sum := 0
	for _, n := range lines {
		if n != "" {
			arr := strings.Split(n, " ")
			if arr[1] == "X" {
				if arr[0] == "A" {
					sum = sum + 1 + 3
				}
				if arr[0] == "B" {
					sum = sum + 1 + 0
				}
				if arr[0] == "C" {
					sum = sum + 1 + 6
				}

			} else if arr[1] == "Y" {
				if arr[0] == "A" {
					sum = sum + 2 + 6
				}
				if arr[0] == "B" {
					sum = sum + 2 + 3
				}
				if arr[0] == "C" {
					sum = sum + 2 + 0
				}
			} else if arr[1] == "Z" {
				if arr[0] == "A" {
					sum = sum + 3 + 0
				}
				if arr[0] == "B" {
					sum = sum + 3 + 6
				}
				if arr[0] == "C" {
					sum = sum + 3 + 3
				}
			}
		}
	}
	fmt.Println(sum)
}

func partTwo(lines []string) {
	//X = lose = 0
	//Y = draw = 3
	//Z = win = 6
	//A, X = Rock = 1
	//B, Y = paper = 2
	//C, Z = Scissors = 3
	sum := 0
	for _, n := range lines {
		if n != "" {
			arr := strings.Split(n, " ")
			if arr[1] == "X" {
				// ROCK
				// need to loose
				if arr[0] == "A" {
					sum = sum + 3 + 0
				}
				if arr[0] == "B" {
					sum = sum + 1 + 0
				}
				if arr[0] == "C" {
					sum = sum + 2 + 0
				}

			} else if arr[1] == "Y" {
				// paper
				// need to draw
				if arr[0] == "A" {
					sum = sum + 1 + 3
				}
				if arr[0] == "B" {
					sum = sum + 2 + 3
				}
				if arr[0] == "C" {
					sum = sum + 3 + 3
				}
			} else if arr[1] == "Z" {
				// scissors
				// need to win
				if arr[0] == "A" {
					sum = sum + 2 + 6
				}
				if arr[0] == "B" {
					sum = sum + 3 + 6
				}
				if arr[0] == "C" {
					sum = sum + 1 + 6
				}
			}
		}
	}
	fmt.Println(sum)
}
