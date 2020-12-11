package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	ints := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parsedInt, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, parsedInt)
	}

	for _, f := range ints {
		for _, s := range ints {
			if f+s == 2020 {
				fmt.Printf("Result: %s\n", f*s)
			}
		}
	}

	for _, f := range ints {
		for _, s := range ints {
			for _, t := range ints {
				if f+s+t == 2020 {
					fmt.Printf("Result: %s\n", f*s*t)
				}
			}
		}
	}
}
