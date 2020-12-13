package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Pass struct {
	Row    int
	Column int
}

func (p *Pass) SeatId() int {
	return p.Row*8 + p.Column
}

func ParsePass(line string) Pass {
	rowMin := 0
	rowWindow := 64
	for _, r := range line[:7] {
		if r == 'B' {
			rowMin += rowWindow
		}
		rowWindow /= 2
	}
	colMin := 0
	colWindow := 4
	for _, c := range line[7:] {
		if c == 'R' {
			colMin += colWindow
		}
		colWindow /= 2
	}
	return Pass{rowMin, colMin}
}

func task1(passes []Pass) {
	maxSeatId := 0
	for _, p := range passes {
		if p.SeatId() > maxSeatId {
			maxSeatId = p.SeatId()
		}
	}
	fmt.Printf("Task 1 result: %s\n", maxSeatId)
}

func task2(passes []Pass) {
	ids := make([]int, 0)
	for _, p := range passes {
		ids = append(ids, p.SeatId())
	}
	sort.Ints(ids)

	prev := ids[0]
	for _, id := range ids[1 : len(ids)-1] {
		// fmt.Printf("DEBUG: %s %s\n", prev, id)
		if id-prev > 1 {
			fmt.Printf("Task 2 result: %s\n", id-1)
			return
		}
		prev = id
	}
}

func main() {
	examples := [4]string{"FBFBBFFRLR", "BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	for _, example := range examples {
		pass := ParsePass(example)
		fmt.Printf("Example '%s': %s\n", example, pass)
	}
	passes := make([]Pass, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		passes = append(passes, ParsePass(scanner.Text()))
	}
	task1(passes)
	task2(passes)
}
