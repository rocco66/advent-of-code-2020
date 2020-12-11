package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PasswordLine struct {
	Min      int
	Max      int
	Char     byte
	Password string
}

func task1(passwords []PasswordLine) {
	var validCounter int

	for _, p := range passwords {
		var charCount, i int
		for i < len(p.Password) {
			if p.Password[i] == p.Char {
				charCount++
			}
			i++
		}
		if charCount >= p.Min && charCount <= p.Max {
			validCounter++
		}
	}

	fmt.Printf("Task 1 result: %s\n", validCounter)
}

func task2(passwords []PasswordLine) {
	var validCounter int

	for _, p := range passwords {
		if (p.Password[p.Min-1] == p.Char && p.Password[p.Max-1] != p.Char) || (p.Password[p.Min-1] != p.Char && p.Password[p.Max-1] == p.Char) {
			validCounter++
		}
	}

	fmt.Printf("Task 2 result: %s\n", validCounter)
}

func main() {
	passwords := make([]PasswordLine, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		minMax := strings.Split(words[0], "-")
		parsedMin, _ := strconv.Atoi(minMax[0])
		parsedMax, _ := strconv.Atoi(minMax[1])
		passwords = append(passwords, PasswordLine{parsedMin, parsedMax, words[1][0], words[2]})
	}

	task1(passwords)
	task2(passwords)
}
