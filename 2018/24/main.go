package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	debug                = false
	debugRead            = debug && false
	debugTargetSelection = debug && false
	debugAttack          = debug && false
	debugMinimumSearch   = debug && false
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
	inf faction = 2 // A group of infection cells
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

// A unit is dead if it runs out of units and should not take part in combat
func (g *group) alive() bool {
	return g.count > 0
}

func (g *group) projectedDamage(target *group) int {
	// If the target is immune, it takes no damage
	if _, present := target.immuneTo[g.attack]; present {
		return 0
	}

	damage := g.power()
	// If the target is weak, it takes double damage
	if _, present := target.weakTo[g.attack]; present {
		damage *= 2
	}

	return damage
}

// A body has immune cells and infection cells battling in it
type body struct {
	groups []*group
}

func (b *body) liveImmuneGroups() int {
	count := 0
	for _, g := range b.groups {
		if g.alive() && g.team == imm {
			count++
		}
	}
	return count
}

func (b *body) liveInfectedGroups() int {
	count := 0
	for _, g := range b.groups {
		if g.alive() && g.team == inf {
			count++
		}
	}
	return count
}

func (b *body) sortByPower() {
	sort.Slice(b.groups, func(i, j int) bool {
		// Sort by power descending
		if b.groups[i].power() > b.groups[j].power() {
			return true
		} else if b.groups[i].power() < b.groups[j].power() {
			return false
		}
		// Fall back to initiative descending
		return b.groups[i].initiative > b.groups[j].initiative
	})
}

func (b *body) sortByInitiative() {
	sort.Slice(b.groups, func(i, j int) bool {
		// Sort by initiative descending
		return b.groups[i].initiative > b.groups[j].initiative
	})
}

// Two teams enter.  One team leaves.  Unless it's a stalemate.  Then return true.
func (b *body) battle() bool {
	for (b.liveImmuneGroups() > 0) && (b.liveInfectedGroups() > 0) {
		// Target Selection
		targets := make(map[int]*group)
		targetedBy := make(map[int]int)
		b.sortByPower()
		for _, attacker := range b.groups {
			if !attacker.alive() {
				if debugTargetSelection {
					fmt.Printf("TARGET: attacker %d is dead, skipping\n", attacker.id)
				}
				continue
			}
			bestTarget := (*group)(nil)
			for _, target := range b.groups {
				if !target.alive() || (target.team == attacker.team) {
					// We shouldn't attack deceased groups or groups on our side
					if debugTargetSelection {
						if !target.alive() {
							fmt.Printf("TARGET: target %d is dead\n", target.id)
						} else {
							fmt.Printf("TARGET: target %d is on same team as attacker %d\n", target.id, attacker.id)
						}
					}
					continue
				}
				if _, present := targetedBy[target.id]; present {
					// Each group may only be targeted by one unit
					if debugTargetSelection {
						fmt.Printf("TARGET: target %d is already targeted\n", target.id)
					}
					continue
				}
				if isBetterTarget(attacker, target, bestTarget) {
					bestTarget = target
					if debugTargetSelection {
						fmt.Printf("TARGET: attacker %d best target = %d\n",
							attacker.id,
							bestTarget.id)
					}
				}
			}
			if bestTarget != nil {
				targets[attacker.id] = bestTarget
				targetedBy[bestTarget.id] = attacker.id
			} else if debugTargetSelection {
				fmt.Printf("TARGET: attacker %d selected no target\n", attacker.id)
			}
		}

		// Attacking
		b.sortByInitiative()
		totalUnitsKilled := 0
		for _, attacker := range b.groups {
			if !attacker.alive() {
				continue
			}
			// If a unit has no target, it doesn't attack
			if _, present := targets[attacker.id]; !present {
				continue
			}
			target := targets[attacker.id]
			damage := attacker.projectedDamage(target)
			deadUnits := int(damage / target.hp)
			target.count -= deadUnits
			totalUnitsKilled += deadUnits
			if debugAttack {
				fmt.Printf("ATTACK: attacker %d deals %d damage to group %d (%d units die, %d remain)\n",
					attacker.id,
					damage,
					target.id,
					deadUnits,
					target.count)
			}
		}
		if totalUnitsKilled == 0 {
			// Stalemate
			if debugAttack {
				fmt.Println("Stalemate detected")
			}
			return true
		}
	}
	return false
}

func isBetterTarget(attacker, target, bestTarget *group) bool {
	damage := attacker.projectedDamage(target)
	if debugTargetSelection {
		fmt.Printf("TARGET: Attacker %d would deal %d damage to target %d\n", attacker.id, damage, target.id)
	}
	if damage < 1 {
		// If a unit wouldn't be damaged, it shoudln't be targeted
		return false
	}

	if bestTarget == nil {
		// Any target is better than no target
		return true
	}

	bestDamage := attacker.projectedDamage(bestTarget)
	if damage == bestDamage {
		// If the damage to be done is equal, choose the target with the highest
		// effective power.  If tie, fallback to highest initiative
		if target.power() == bestTarget.power() {
			return target.initiative > bestTarget.initiative
		}
		return target.power() > bestTarget.power()
	}
	return damage > bestDamage
}

func (b *body) winningArmyCount() (int, faction) {
	var f faction
	stalemate := b.battle()
	if stalemate {
		// Stalemate means the immune system doesn't win.
		return 0, inf
	}
	total := 0
	for _, g := range b.groups {
		if g.alive() {
			total += g.count
			f = g.team
		}
	}
	return total, f
}

// Finds the minimum attack boost to the cause the immune team to win.
// Returns the minimum boost and the number of units left after that boost
func findMinimumBoost(input []string) (int, int) {
	lower := 0
	upper := 5000

	// Ensure the upper bound issane
	for {
		_, winningTeam := countWithBoost(input, upper)
		if winningTeam == inf {
			if debugMinimumSearch {
				fmt.Printf("%d is too low, doubling upper bound\n", upper)
			}
			lower = upper
			upper *= 2
		} else {
			if debugMinimumSearch {
				fmt.Printf("Upper bound: %d\n", upper)
			}
			break
		}
	}

	// Binary Search
	for upper-lower > 1 {
		boost := lower + (upper-lower)/2
		if debugMinimumSearch {
			fmt.Printf("Range [%d-%d], Guess: %d\n", lower, upper, boost)
		}
		_, winningTeam := countWithBoost(input, boost)

		if winningTeam == imm {
			upper = boost
		} else {
			lower = boost
		}
	}

	// Measure the final amounts
	remaining, _ := countWithBoost(input, upper)
	return upper, remaining
}

func countWithBoost(input []string, boost int) (int, faction) {
	b := readBody(input)
	for _, g := range b.groups {
		if g.team == imm {
			g.damage += boost
		}
	}
	return b.winningArmyCount()
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
	groups := make([]*group, 0)
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
		groupStr := fmt.Sprint(g)
		if readingImmune {
			g.team = imm
			if debugRead {
				fmt.Printf("Read Immune Group: %s\n", groupStr)
			}
		} else {
			g.team = inf
			if debugRead {
				fmt.Printf("Read Infection Group: %s\n", groupStr)
			}
		}
		groups = append(groups, &g)
	}
	return body{groups: groups}
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
	start = time.Now()
	minBoost, remaining := findMinimumBoost(input)
	fmt.Printf("Minimum boost of %d leaves %d units remaining\n", minBoost, remaining)
	fmt.Println(time.Since(start))
}
