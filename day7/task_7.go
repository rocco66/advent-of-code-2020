package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Bag string

type Rule struct {
	Type  Bag
	Count int
}

type RulesIndex struct {
	Origin map[Bag][]Rule
	Index  map[Bag]bool
}

func ParseBagRule(line string) (Bag, []Rule) {
	var rules []Rule
	splitted := strings.Split(line, "s contain ")
	for _, ruleString := range strings.Split(splitted[1][:len(splitted[1])-1], ", ") {
		ruleTokens := strings.SplitN(ruleString, " ", 2)
		count, _ := strconv.Atoi(ruleTokens[0])
		var bagName string
		if count == 1 {
			bagName = ruleTokens[1]
		} else {
			bagName = ruleTokens[1][:len(ruleTokens[1])-1]
		}
		rules = append(rules, Rule{Bag(bagName), count})
	}
	return Bag(splitted[0]), rules
}

func MakeIndex(rules map[Bag][]Rule) RulesIndex {
	var ruleIndex RulesIndex
	ruleIndex.Origin = rules
	index := make(map[Bag]bool)
	mutexes := make(map[Bag]sync.Mutex)
	flags := make(map[Bag]*sync.Cond)
	for bag, _ := range rules {
		mutexes[bag] = sync.Mutex{}
		flags[bag] = sync.NewCond(&mutexes[bag])
	}
	for bag, rules := range rules {
		_, ok := index[bag]
		if !ok {
			for _, rule := range rules {

			}
		}
	}
	return ruleIndex

}

func task1(ri RulesIndex) {
	res := 0
	for _, mightContainsTarget := range ri.Index {
		if mightContainsTarget {
			res++
		}
	}
	fmt.Printf("Task 1 result: %s\n", res)
}

func task2(ri RulesIndex) {
	// fmt.Printf("Task 2 result: %s\n", res)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rules := make(map[Bag][]Rule)
	for scanner.Scan() {
		bag, bagRules := ParseBagRule(scanner.Text())
		rules[bag] = bagRules
	}
	ruleIndex := MakeIndex(rules)
	task1(ruleIndex)
	task2(ruleIndex)
}
