package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amar-preet/advent-of-code-2022/reader"
)

func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	mostCalories := 0
	count := 0
	for i, n := range lines {

		if n != "" {
			a, _ := strconv.Atoi(lines[i])
			count = count + a
		} else {
			if count > mostCalories {
				mostCalories = count
			}
			count = 0
		}

	}
	fmt.Println(mostCalories)
}

func partTwo(lines []string) {
	count := 0
	var s []int
	for _, n := range lines {

		if n != "" {
			a, _ := strconv.Atoi(n)
			count = count + a
		} else {
			s = append(s, count)
			count = 0
		}

	}
	s = append(s, count)
	findTopThreeCalories(s)
}

func findTopThreeCalories(arr []int) {
	var first int = 0
	var second int = 0
	var third int = 0

	for _, n := range arr {
		if n > first {
			third = second
			second = first
			first = n
		} else if n > second && n != first {
			third = second
			second = n
		} else if n > third && n != second {
			third = n
		}
	}
	fmt.Println("sum", first+second+third)
}
