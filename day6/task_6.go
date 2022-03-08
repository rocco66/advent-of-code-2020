package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type GroupAnswers struct {
	PersonsAnswers []string
}

func (ga *GroupAnswers) SumOfAnswers() map[rune]int {
	res := make(map[rune]int)
	for _, personAnswers := range ga.PersonsAnswers {
		for _, answer := range personAnswers {
			res[answer]++
		}
	}
	return res
}

func (ga *GroupAnswers) UniqAnswers() int {
	return len(ga.SumOfAnswers())
}

func (ga *GroupAnswers) EveryoneYes() int {
	var res int
	for _, yesCount := range ga.SumOfAnswers() {
		if yesCount == len(ga.PersonsAnswers) {
			res++
		}
	}
	return res
}

func ParseGroupAnswers(in string) GroupAnswers {
	return GroupAnswers{strings.Split(in, "\n")}
}

func task1(groupAnswers []GroupAnswers) {
	var res int
	for _, group := range groupAnswers {
		res += group.UniqAnswers()
	}
	fmt.Printf("Task 1 result: %s\n", res)
}

func task2(groupAnswers []GroupAnswers) {
	var res int
	for _, group := range groupAnswers {
		res += group.EveryoneYes()
	}
	fmt.Printf("Task 2 result: %s\n", res)
}

func readSections() <-chan string {
	ch := make(chan string)
	go func() {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		for _, sectionString := range strings.Split(string(bytes), "\n\n") {
			ch <- sectionString
		}
		close(ch)
	}()
	return ch
}

func main() {
	answers := make([]GroupAnswers, 0)
	for section := range readSections() {
		answers = append(answers, ParseGroupAnswers(section))
	}
	task1(answers)
	task2(answers)
}
