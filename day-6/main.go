package main

import (
	"fmt"

	"github.com/amar-preet/advent-of-code-2022/reader"
)

func main() {
	lines := reader.ReadFile("input.txt")
	partOne(lines)
	partTwo(lines)
}

func partOne(lines []byte) {

	for i := 0; i < len(lines)-3; i++ {
		found := false
		fourChars := ""
		count := 0
		for count < 4 {
			fourChars = fourChars + string(lines[i+count])
			count++
		}

	Exit:
		for j := 0; j < len(fourChars); j++ {
			for k := j + 1; k < len(fourChars); k++ {
				if string(fourChars[j]) == string(fourChars[k]) {
					found = true
					break Exit
				}
			}
		}
		if found == false {
			fmt.Println(i + 4)
			break
		}
	}
}

func partTwo(lines []byte) {

	for i := 0; i < len(lines)-13; i++ {
		found := false
		count := 0
		fourteenChars := ""
		for count < 14 {
			fourteenChars = fourteenChars + string(lines[i+count])
			count++
		}

	Exit:
		for j := 0; j < len(fourteenChars); j++ {
			for k := j + 1; k < len(fourteenChars); k++ {
				if string(fourteenChars[j]) == string(fourteenChars[k]) {
					found = true
					break Exit
				}
			}
		}
		if found == false {
			fmt.Println(i + 14)
			break
		}
	}
}
