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
	count := 0
	for _, n := range lines {
		var p1, q1, p2, q2 int
		fmt.Sscanf(n, "%d-%d,%d-%d", &p1, &q1, &p2, &q2)

		if p1 <= p2 && q1 >= q2 || p1 >= p2 && q1 <= q2 {
			count++
		}
	}
	fmt.Println(count)
}

func partTwo(lines []string) {
	count := 0

	for _, n := range lines {
		var p1, q1, p2, q2 int
		fmt.Sscanf(n, "%d-%d,%d-%d", &p1, &q1, &p2, &q2)
		if p1 <= q2 && q1 >= p2 {
			count++
		}
	}
	fmt.Println(count)
}
