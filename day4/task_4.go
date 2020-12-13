package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	Byr int
	Iyr int
	Eyr int
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid int
}

func (p *Passport) IsWeakValid() bool {
	return p.Byr > 0 && p.Iyr > 0 && p.Eyr > 0 && p.Hgt != "" && p.Hcl != "" && p.Ecl != "" && p.Pid != ""
}

var hgtCmRe *regexp.Regexp
var hgtInRe *regexp.Regexp
var hclRe *regexp.Regexp
var eclRe *regexp.Regexp
var pidRe *regexp.Regexp

func initRegex() {
	hgtCmRe = regexp.MustCompile(`^(\d{3})cm$`)
	hgtInRe = regexp.MustCompile(`^(\d{2})in$`)
	hclRe = regexp.MustCompile(`^#[\da-f]{6}$`)
	eclRe = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	pidRe = regexp.MustCompile(`^\d{9}$`)
}

func (p *Passport) IsStrongValid() bool {
	var hgtMatched bool
	hgtCmMatched := hgtCmRe.FindStringSubmatch(p.Hgt)
	if len(hgtCmMatched) > 0 {
		parsedHgtCm, _ := strconv.Atoi(hgtCmMatched[1])
		hgtMatched = parsedHgtCm >= 150 && parsedHgtCm <= 193
	} else {
		hgtInMatched := hgtInRe.FindStringSubmatch(p.Hgt)
		if len(hgtInMatched) > 0 {
			parsedHgtIn, _ := strconv.Atoi(hgtInMatched[1])
			hgtMatched = parsedHgtIn >= 59 && parsedHgtIn <= 76
		}
	}

	return p.Byr >= 1920 && p.Byr <= 2002 &&
		p.Iyr >= 2010 && p.Iyr <= 2020 &&
		p.Eyr >= 2020 && p.Eyr <= 2030 &&
		hgtMatched && hclRe.MatchString(p.Hcl) && eclRe.MatchString(p.Ecl) && pidRe.MatchString(p.Pid)
}

func ParsePassport(in string) Passport {
	p := Passport{0, 0, 0, "", "", "", "", 0}
	for _, line := range strings.Split(in, "\n") {
		for _, pair := range strings.Fields(line) {
			tokens := strings.Split(pair, ":")
			switch tokens[0] {
			case "byr":
				p.Byr, _ = strconv.Atoi(tokens[1])
			case "iyr":
				p.Iyr, _ = strconv.Atoi(tokens[1])
			case "eyr":
				p.Eyr, _ = strconv.Atoi(tokens[1])
			case "hgt":
				p.Hgt = tokens[1]
			case "hcl":
				p.Hcl = tokens[1]
			case "ecl":
				p.Ecl = tokens[1]
			case "pid":
				p.Pid = tokens[1]
			case "cid":
				p.Cid, _ = strconv.Atoi(tokens[1])
			}
		}
	}
	return p
}

func task1(passports []Passport) {
	resultCounter := 0
	for _, p := range passports {
		if p.IsWeakValid() {
			resultCounter++
		}
	}
	fmt.Printf("Task 1 result: %s\n", resultCounter)
}

func task2(passports []Passport) {
	resultCounter := 0
	for _, p := range passports {
		if p.IsStrongValid() {
			resultCounter++
		}
	}
	fmt.Printf("Task 2 result: %s\n", resultCounter)
}

func main() {
	initRegex()
	passports := make([]Passport, 0)
	bytes, _ := ioutil.ReadAll(os.Stdin)
	for _, passwordStrings := range strings.Split(string(bytes), "\n\n") {
		passports = append(passports, ParsePassport(passwordStrings))
	}
	task1(passports)
	task2(passports)
}
