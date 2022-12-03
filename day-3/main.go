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
	found := ""
Exit:
	for _, n := range lines {
		if n != "" {
			var first []string
			var second []string
			for i := 0; i < len(n)/2; i++ {
				first = append(first, string(n[i]))
			}
			for j := len(n) / 2; j < len(n); j++ {
				second = append(second, string(n[j]))
			}

			for i := 0; i < len(first); i++ {
				for j := 0; j < len(second); j++ {
					if first[i] == second[j] {
						found = found + first[i]
						continue Exit
					}
				}
			}

		}
	}

	fmt.Println("SUM", getSum(found))
}

func partTwo(lines []string) {
	found := ""
	line := 0
Exit:
	for l := line; l < len(lines); l = line + 2 {
		var subarray []string
		count := 0
		for count < 3 {
			subarray = append(subarray, lines[line])
			line++
			count++
		}

		firstRow := subarray[0]
		secondRow := subarray[1]
		thirdRow := subarray[2]
		for i := 0; i < len(firstRow); i++ {
			for j := 0; j < len(secondRow); j++ {
				for k := 0; k < len(thirdRow); k++ {
					if firstRow[i] == secondRow[j] && firstRow[i] == thirdRow[k] {
						found = found + string(firstRow[i])
						continue Exit
					}
				}
			}
		}
	}
	fmt.Println(getSum(found))
}

func getSum(found string) int {
	lowerMap := lowerCase()
	upperMap := upperCase()
	sum := 0
	for _, n := range strings.Split(found, "") {

		value, exists := lowerMap[n]
		if exists {
			sum = sum + value
		} else {
			value2, exists := upperMap[n]
			if exists {
				sum = sum + value2
			}
		}
	}
	return sum
}

func lowerCase() map[string]int {
	input := make(map[string]int)
	count := 0
	for ch := 'a'; ch <= 'z'; ch++ {
		count = count + 1
		s := fmt.Sprintf("%c", ch)
		input[s] = count
	}
	return input
}

func upperCase() map[string]int {
	input := make(map[string]int)
	count := 26
	for ch := 'A'; ch <= 'Z'; ch++ {
		count = count + 1
		s := fmt.Sprintf("%c", ch)
		input[s] = count
	}
	return input
}
