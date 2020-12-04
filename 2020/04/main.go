package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"../../aoc"
)

// https://adventofcode.com/2020/day/4
// Scan a list of passport data and check for valid passports

func main() {
	lines := aoc.ReadFileLines("input.txt")
	sw := aoc.NewStopwatch()
	fmt.Println("Step 1:")
	pps := readPassports(lines)
	fmt.Println(countPresent(pps))
	fmt.Println(sw.Elapsed())

	sw.Reset()
	fmt.Println("Step 2:")
	fmt.Println(countValid(pps))
	fmt.Println(sw.Elapsed())
}

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func (p *passport) AllPresent() bool {
	return len(p.byr) > 0 && len(p.iyr) > 0 && len(p.eyr) > 0 && len(p.hgt) > 0 && len(p.hcl) > 0 && len(p.ecl) > 0 && len(p.pid) > 0
}

var (
	heightRegex = regexp.MustCompile("^([0-9]+)(cm|in)$")
	hairRegex   = regexp.MustCompile("^#[0-9a-f]{6}$")
	pidRegex    = regexp.MustCompile("^[0-9]{9}$")
	eyeRegex    = regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
)

func (p *passport) Valid() bool {
	return p.AllPresent() &&
		p.validBirthYear() &&
		p.validIssueYear() &&
		p.validExpYear() &&
		p.validHeight() &&
		p.validHair() &&
		p.validEye() &&
		p.validPassportId()
}

func (p *passport) validBirthYear() bool {
	y, err := strconv.Atoi(p.byr)
	return err == nil && y >= 1920 && y <= 2002
}

func (p *passport) validIssueYear() bool {
	y, err := strconv.Atoi(p.iyr)
	return err == nil && y >= 2010 && y <= 2020
}

func (p *passport) validExpYear() bool {
	y, err := strconv.Atoi(p.eyr)
	return err == nil && y >= 2020 && y <= 2030
}

func (p *passport) validHeight() bool {
	h := heightRegex.FindStringSubmatch(p.hgt)
	if h == nil {
		return false
	}
	ht, err := strconv.Atoi(h[1])
	if err != nil || (h[2] == "cm" && (ht < 150 || ht > 193)) || (h[2] == "in" && (ht < 59 || ht > 76)) {
		return false
	}
	return true
}

func (p *passport) validHair() bool {
	return hairRegex.FindStringSubmatch(p.hcl) != nil
}

func (p *passport) validEye() bool {
	return eyeRegex.FindStringSubmatch(p.ecl) != nil
}

func (p *passport) validPassportId() bool {
	return pidRegex.FindStringSubmatch(p.pid) != nil
}

func countPresent(pps []passport) int {
	v := 0
	for _, p := range pps {
		if p.AllPresent() {
			v++
		}
	}
	return v
}

func countValid(pps []passport) int {
	v := 0
	for _, p := range pps {
		if p.Valid() {
			v++
		}
	}
	return v
}

func readPassports(lines []string) []passport {
	var pps []passport
	var p passport
	readData := true
	for _, l := range lines {
		if len(l) == 0 {
			pps = append(pps, p)
			p = passport{}
			readData = false
			continue
		}
		for _, attrib := range strings.Split(l, " ") {
			a := strings.Split(attrib, ":")
			switch a[0] {
			case "byr":
				p.byr = a[1]
			case "iyr":
				p.iyr = a[1]
			case "eyr":
				p.eyr = a[1]
			case "hgt":
				p.hgt = a[1]
			case "hcl":
				p.hcl = a[1]
			case "ecl":
				p.ecl = a[1]
			case "pid":
				p.pid = a[1]
			case "cid":
				p.cid = a[1]
			default:
				fmt.Println("Unexpected: " + a[0])
			}
		}
		readData = true
	}
	if readData {
		pps = append(pps, p)
	}
	return pps
}
