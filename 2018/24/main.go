package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// Input format: "956 units each with 7120 hit points (weak to bludgeoning, slashing) with an attack that does 71 radiation damage at initiative 7"
	groupRegEx = regexp.MustCompile(
		"^(\\d+) units each with (\\d+) hit points (\\(?.*\\)? ?)with an attack that does " +
			"(\\d+) ([a-z]+) damage at initiative (\\d+)$")

	lastGroupID = 0
)

type faction int

const (
	imm faction = 1 // A group of immune system cells
	inf faction = 1 // A group of infection cells
)

type group struct {
	id         int
	team       faction
	count      int             // How many units
	hp         int             // HP per unit
	damage     int             // Attack damage
	attack     string          // Attack elemental type
	weakTo     map[string]bool // Elemental weaknesses
	immuneTo   map[string]bool // Elemental immunity
	initiative int             // Turn order
}

// The group's attack power
func (g *group) power() int {
	return g.count * g.damage
}

// A unit is dead if it runs out of HP and should not take part in combat
func (g *group) alive() bool {
	return g.hp > 0
}

// A body has immune cells and infection cells battling in it
type body struct {
	immune    []group
	infection []group
}

func (b *body) liveImmuneGroups() int {
	count := 0
	for _, g := range b.immune {
		if g.alive() {
			count++
		}
	}
	return count
}

func (b *body) liveInfectedGroups() int {
	count := 0
	for _, g := range b.infection {
		if g.alive() {
			count++
		}
	}
	return count
}

// Two teams enter.  One team leaves.
func (b *body) battle() {
	for (b.liveImmuneGroups() > 0) && (b.liveInfectedGroups() > 0) {

	}
}

func (b *body) winningArmyCount() int {
	b.battle()
	total := 0
	for _, g := range b.immune {
		if g.alive() {
			total += g.count
		}
	}
	for _, g := range b.infection {
		if g.alive() {
			total += g.count
		}
	}
	return total
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func readGroup(input string) group {
	tokens := groupRegEx.FindStringSubmatch(input)
	if len(tokens) < 1 {
		log.Fatalf("Error reading group: \"%s\"", input)
	}
	count, err := strconv.Atoi(tokens[1])
	check(err)
	hp, err := strconv.Atoi(tokens[2])
	check(err)
	weak, immune := parseVulnerabilities(tokens[3])
	check(err)
	damage, err := strconv.Atoi(tokens[4])
	check(err)
	attack := tokens[5]
	check(err)
	init, err := strconv.Atoi(tokens[6])
	check(err)

	lastGroupID++
	return group{
		count: count, hp: hp, weakTo: weak, immuneTo: immune,
		damage: damage, attack: attack, initiative: init, id: lastGroupID}
}

// Reads the weaknesses/immunities
// Format examples:
// weak to bludgeoning, slashing
// weak to bludgeoning; immune to cold
// weak to cold; immune to bludgeoning, slashing
// immune to slashing; weak to bludgeoning
// Returns weakTo, immuneTo
func parseVulnerabilities(input string) (map[string]bool, map[string]bool) {
	weakTo := make(map[string]bool)
	immuneTo := make(map[string]bool)
	if len(input) > 0 {
		// Trim off the leading ( and the trailing ") "
		input = input[1 : len(input)-2]
		chunks := strings.Split(input, "; ")
		for _, c := range chunks {
			if c[0:7] == "weak to" {
				for _, w := range strings.Split(c[8:], ", ") {
					weakTo[w] = true
				}
			} else if c[0:9] == "immune to" {
				for _, w := range strings.Split(c[10:], ", ") {
					immuneTo[w] = true
				}
			} else {
				log.Fatalf("Unexpected vulnerabilities: %s\n", input)
			}
		}
	}
	return weakTo, immuneTo
}

// Returns immune, infection
func readBody(input []string) body {
	immune := make([]group, 0)
	infection := make([]group, 0)
	readingImmune := true
	for _, s := range input {
		if (len(s) == 0) || (s[:2] == "Im") {
			continue
		}
		if s[:2] == "In" {
			readingImmune = false
			continue
		}

		g := readGroup(s)
		if readingImmune {
			g.team = imm
			immune = append(immune, g)
		} else {
			g.team = inf
			infection = append(infection, g)
		}
	}
	return body{immune: immune, infection: infection}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())
	start := time.Now()
	b := readBody(input)
	fmt.Println(b.winningArmyCount())
	fmt.Println(time.Since(start))
}
