package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amar-preet/advent-of-code-2022/reader"
)

type ItemType interface{}

type Stack struct {
	items []ItemType
}
type Stacks []Stack

func (stack *Stack) Push(t ItemType) {
	if stack.items == nil {
		stack.items = []ItemType{}
	}
	stack.items = append(stack.items, t)
}

func (stack *Stack) Pop() *ItemType {
	if len(stack.items) == 0 {
		return nil
	}
	item := stack.items[len(stack.items)-1]
	stack.items = stack.items[0 : len(stack.items)-1]
	return &item
}

func main() {
	lines := strings.Split(string(reader.ReadFile("input.txt")), "\n")
	partOne(lines)
	partTwo(lines)
}

//    [D]
//[N] [C]
//[Z] [M] [P]
// 1   2   3

//[H]         [D]     [P]
//[W] [B]         [C] [Z] [D]
//[T] [J]     [T] [J] [D] [J]
//[H] [Z]     [H] [H] [W] [S]     [M]
//[P] [F] [R] [P] [Z] [F] [W]     [F]
//[J] [V] [T] [N] [F] [G] [Z] [S] [S]
//[C] [R] [P] [S] [V] [M] [V] [D] [Z]
//[F] [G] [H] [Z] [N] [P] [M] [N] [D]
// 1   2   3   4   5   6   7   8   9

func partOne(lines []string) {
	s1 := Stack{}
	s1.Push("Z")
	s1.Push("N")

	s2 := Stack{}
	s2.Push("M")
	s2.Push("C")
	s2.Push("D")

	s3 := Stack{}
	s3.Push("P")
	Stacks := Stacks{}
	Stacks = append(Stacks, s1, s2, s3)
	for _, n := range lines {
		var move, from, to string
		fmt.Sscanf(n, "move %s from %s to %s", &move, &from, &to)

		intMove, _ := strconv.Atoi(move)
		intFrom, _ := strconv.Atoi(from)
		intTo, _ := strconv.Atoi(to)
		for i := 0; i < int(intMove); i++ {

			k := Stacks[intFrom-1].Pop()
			Stacks[intTo-1].Push(*k)
		}
	}

	for i := 0; i < len(Stacks); i++ {
		fmt.Println(Stacks[i].items[len(Stacks[i].items)-1])
	}

}

func partTwo(lines []string) {
	s1 := Stack{}
	s1.Push("F")
	s1.Push("C")
	s1.Push("J")
	s1.Push("P")
	s1.Push("H")
	s1.Push("T")
	s1.Push("W")

	s2 := Stack{}
	s2.Push("G")
	s2.Push("R")
	s2.Push("V")
	s2.Push("F")
	s2.Push("Z")
	s2.Push("J")
	s2.Push("B")
	s2.Push("H")

	s3 := Stack{}
	s3.Push("H")
	s3.Push("P")
	s3.Push("T")
	s3.Push("R")

	s4 := Stack{}
	s4.Push("Z")
	s4.Push("S")
	s4.Push("N")
	s4.Push("P")
	s4.Push("H")
	s4.Push("T")

	s5 := Stack{}
	s5.Push("N")
	s5.Push("V")
	s5.Push("F")
	s5.Push("Z")
	s5.Push("H")
	s5.Push("J")
	s5.Push("C")
	s5.Push("D")

	s6 := Stack{}
	s6.Push("P")
	s6.Push("M")
	s6.Push("G")
	s6.Push("F")
	s6.Push("W")
	s6.Push("D")
	s6.Push("Z")

	s7 := Stack{}
	s7.Push("M")
	s7.Push("V")
	s7.Push("Z")
	s7.Push("W")
	s7.Push("S")
	s7.Push("J")
	s7.Push("D")
	s7.Push("P")

	s8 := Stack{}
	s8.Push("N")
	s8.Push("D")
	s8.Push("S")

	s9 := Stack{}
	s9.Push("D")
	s9.Push("Z")
	s9.Push("S")
	s9.Push("F")
	s9.Push("M")

	Stacks := Stacks{}
	Stacks = append(Stacks, s1, s2, s3, s4, s5, s6, s7, s8, s9)
	for _, n := range lines {
		var move, from, to string
		fmt.Sscanf(n, "move %s from %s to %s", &move, &from, &to)

		intMove, _ := strconv.Atoi(move)
		intFrom, _ := strconv.Atoi(from)
		intTo, _ := strconv.Atoi(to)
		for i := 0; i < int(intMove); i++ {
			if intMove == 1 {
				k := Stacks[intFrom-1].Pop()
				Stacks[intTo-1].Push(*k)
			} else {
				k := Stacks[intFrom-1].items[0]
				copy(Stacks[intFrom-1].items[0:], Stacks[intFrom-1].items[0+1:])
				Stacks[intFrom-1].items[len(Stacks[intFrom-1].items)-1] = ""
				Stacks[intFrom-1].items = Stacks[intFrom-1].items[:len(Stacks[intFrom-1].items)-1]
				Stacks[intTo-1].Push(k)

			}

		}
	}

	for i := 0; i < len(Stacks); i++ {
		fmt.Println(Stacks[i].items[len(Stacks[i].items)-1])
	}

}
