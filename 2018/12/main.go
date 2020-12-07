package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	// Input format: "initial state: #.#.#...#..##..###.##.#...#.##.#....#..#.#....##.#.##...###.#...#######.....##.###.####.#....#.#..##"
	stateRegEx = regexp.MustCompile(`^initial state: ([.#]+)$`)
	// Input format: "#...# => #"
	ruleRegEx = regexp.MustCompile(`^([.#][.#][.#][.#][.#]) => ([.#])$`)
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func hasPlant(r rune) bool {
	return r == '#'
}

// PotRow represents which pots in a row have plants
type PotRow map[int64]bool

// State represents the state of the row of pots
type State struct {
	plants   PotRow
	minPlant int64
	maxPlant int64
}

func min(i, j int64) int64 {
	if i < j {
		return i
	}
	return j
}

func max(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}

// ReadInitialState reads the initial state of the pots
func ReadInitialState(s string) State {
	state := make(PotRow)
	tokens := stateRegEx.FindStringSubmatch(s)
	initialState := []rune(tokens[1])
	minPlant := int64(len(initialState))
	maxPlant := int64(0)
	for i := 0; i < len(initialState); i++ {
		if hasPlant(initialState[i]) {
			state[int64(i)] = true
			minPlant = min(minPlant, int64(i))
			maxPlant = max(maxPlant, int64(i))
		}
	}
	return State{plants: state, minPlant: minPlant, maxPlant: maxPlant}
}

// Rule describes a rule on how plants change across generations
type Rule struct {
	leftTwo  bool
	leftOne  bool
	current  bool
	rightOne bool
	rightTwo bool
	result   bool
}

// RuleSet allows for quick hash lookup of rules
type RuleSet map[bool]map[bool]map[bool]map[bool]map[bool]bool
type ruleSet1 map[bool]map[bool]map[bool]map[bool]bool
type ruleSet2 map[bool]map[bool]map[bool]bool
type ruleSet3 map[bool]map[bool]bool
type ruleSet4 map[bool]bool

// MakeRuleSet creates a RuleSet from a list of rules
func MakeRuleSet(rules []Rule) RuleSet {
	rs := make(RuleSet)
	rs[true] = makeRuleSet1()
	rs[false] = makeRuleSet1()

	for _, r := range rules {
		rs[r.leftTwo][r.leftOne][r.current][r.rightOne][r.rightTwo] = r.result
	}

	return rs
}

func makeRuleSet1() ruleSet1 {
	rs := make(ruleSet1)
	rs[true] = makeRuleSet2()
	rs[false] = makeRuleSet2()
	return rs
}

func makeRuleSet2() ruleSet2 {
	rs := make(ruleSet2)
	rs[true] = makeRuleSet3()
	rs[false] = makeRuleSet3()
	return rs
}

func makeRuleSet3() ruleSet3 {
	rs := make(ruleSet3)
	rs[true] = make(ruleSet4)
	rs[false] = make(ruleSet4)
	return rs
}

// ReadRule parses a line from the input file
func ReadRule(input string) Rule {
	tokens := ruleRegEx.FindStringSubmatch(input)
	condition := []rune(tokens[1])
	result := []rune(tokens[2])

	return Rule{
		leftTwo:  hasPlant(condition[0]),
		leftOne:  hasPlant(condition[1]),
		current:  hasPlant(condition[2]),
		rightOne: hasPlant(condition[3]),
		rightTwo: hasPlant(condition[4]),
		result:   hasPlant(result[0])}
}

// ReadRules parses the rules lines from the input file
func ReadRules(input []string) []Rule {
	var rules []Rule
	for _, s := range input {
		rules = append(rules, ReadRule(s))
	}
	return rules
}

// Advance advances the pots by the requested number of generations
func Advance(initialState *State, rules RuleSet, generations int64) State {
	state := *initialState
	for i := int64(0); i < generations; i++ {
		state = advanceGeneration(&state, rules)
		if i%1000000 == 0 {
			fmt.Printf("%d\n", i)
		}
	}
	return state
}

func advanceGeneration(initialState *State, rules RuleSet) State {
	newState := State{plants: make(PotRow), minPlant: initialState.maxPlant, maxPlant: initialState.minPlant}
	foundPlant := false
	for i := initialState.minPlant - 2; i <= initialState.maxPlant+2; i++ {
		if willHavePlant(initialState.plants, rules, i) {
			newState.plants[i] = true
			newState.minPlant = min(newState.minPlant, i)
			newState.maxPlant = max(newState.maxPlant, i)
			foundPlant = true
		}
	}

	if !foundPlant {
		newState.minPlant, newState.maxPlant = int64(0), int64(0)
	}
	return newState
}

// Match to a rule
func willHavePlant(plants PotRow, rules RuleSet, index int64) bool {
	return rules[plants[index-2]][plants[index-1]][plants[index]][plants[index+1]][plants[index+2]]
}

// Score sums all of the pot numbers that have plants
func Score(state *State) int64 {
	score := int64(0)
	for k, v := range state.plants {
		if v {
			score += k
		}
	}
	return score
}

// FindCycle attempts to find a pattern in the generations
func FindCycle(state State, rules RuleSet) (int64, int64) {
	generations := make(map[string]int64)
	for t := int64(0); true; t++ {
		st := stateToString(&state)
		if previousGeneration, seen := generations[st]; seen {
			// Return the start of the cycle and its length
			return previousGeneration, t - previousGeneration
		}
		generations[st] = t
		state = advanceGeneration(&state, rules)
	}
	return -1, -1
}

func stateToString(state *State) string {
	var sb strings.Builder
	for i := state.minPlant; i <= state.maxPlant; i++ {
		if state.plants[i] {
			sb.WriteString("#")
		} else {
			sb.WriteString(".")
		}
	}
	return sb.String()
}

func printScoreSummary(state State, rules RuleSet, generations int) {
	for i := 0; i <= generations; i++ {
		fmt.Printf("%d: %d %s [%d...%d (%d)]\n", i, Score(&state), stateToString(&state), state.minPlant, state.maxPlant, state.maxPlant-state.minPlant)
		state = advanceGeneration(&state, rules)
	}
}

func finalScore(generation int64) int64 {
	// After generation 100, we end up in a repeating pattern shifting to the right:
	// 998: 60480 #..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..##....#..##..##..##..#..#..##....#..## [953...1097 (144)]
	// 999: 60539 #..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..##....#..##..##..##..#..#..##....#..## [954...1098 (144)]
	minValue := generation - 45
	pattern := []rune("#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..#..#..##..##....#..##..##..##..#..#..##....#..##")
	score := int64(0)
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '#' {
			score += (minValue + int64(i))
		}
	}
	return score
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	state := ReadInitialState(scanner.Text())
	scanner.Scan() // Read the blank line
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())
	start := time.Now()
	rules := MakeRuleSet(ReadRules(input))
	newState := Advance(&state, rules, 20)
	fmt.Println(Score(&newState))
	fmt.Println(time.Since(start))
	start = time.Now()
	cycleStart, cycleLength := FindCycle(state, rules)
	fmt.Printf("Found cycle starting at generation %d of length %d\n", cycleStart, cycleLength)
	fmt.Printf("Score after generation 50000000000: %d\n", finalScore(50000000000))
	fmt.Println(time.Since(start))
}
