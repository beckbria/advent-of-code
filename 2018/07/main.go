package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"time"
)

const (
	debug = false // If debug spew should be printed to the console
)

var (
	// Input format: "Step Y must be finished before step A can begin."
	requirementRegEx = regexp.MustCompile(`^Step (.) must be finished before step (.) can begin.`)
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// CharacterMap represents a mapping from characters to characters.  It is a map[byte, map[byte, bool]] because
// go has neither the concept of a set nor the concept of easily erasing from a slice by value.
type CharacterMap map[byte]map[byte]bool

// Apparently sort.Sort doesn't understand byte arrays, and you can't implement
// Len/Less/Swap on []byte directly.  Go REALLY needs generics.
type byteSlice []byte

func (a byteSlice) Len() int           { return len(a) }
func (a byteSlice) Less(i, j int) bool { return a[i] < a[j] }
func (a byteSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// ReadDependencies reads the input (a series of step dependencies) and returns two maps:
// First: map[node]dependsOn  Second: map[node]isDependedOnBy.
func ReadDependencies(input []string) (CharacterMap, CharacterMap) {
	dependsOn := make(CharacterMap)
	isDependedOnBy := make(CharacterMap)

	for _, s := range input {
		tokens := requirementRegEx.FindStringSubmatch(s)
		dependent := []byte(tokens[2])[0]
		dependency := []byte(tokens[1])[0]

		// Ensure all maps exist
		if _, present := dependsOn[dependency]; !present {
			dependsOn[dependency] = make(map[byte]bool)
		}
		if _, present := dependsOn[dependent]; !present {
			dependsOn[dependent] = make(map[byte]bool)
		}
		if _, present := isDependedOnBy[dependency]; !present {
			isDependedOnBy[dependency] = make(map[byte]bool)
		}
		if _, present := isDependedOnBy[dependent]; !present {
			isDependedOnBy[dependent] = make(map[byte]bool)
		}

		dependsOn[dependent][dependency] = true
		isDependedOnBy[dependency][dependent] = true
	}

	return dependsOn, isDependedOnBy
}

// copyCharacterMap does a deep copy of a map[map].
func copyCharacterMap(m CharacterMap) CharacterMap {
	cp := make(CharacterMap)
	for k, v := range m {
		cp[k] = make(map[byte]bool)
		for vk, vv := range v {
			cp[k][vk] = vv
		}
	}
	return cp
}

// ExecutionOrder sorts the dependencies and returns an order in which the nodes can be executed
func ExecutionOrder(dep, isDep CharacterMap) string {
	// Make a copy of each map so that we can modify them
	var order []byte
	dependsOn := copyCharacterMap(dep)
	isDependedOnBy := copyCharacterMap(isDep)

	for len(dependsOn) > 0 {
		next := firstAvailableNode(dependsOn, nil)
		order = append(order, next)
		delete(dependsOn, next)
		// Remove the dependency from every other option
		for v := range isDependedOnBy[next] {
			delete(dependsOn[v], next)
		}
	}

	return string(order)
}

// firstAvailableNode returns first node with no dependencies.  Will select nodes that
// became available first, with the tiebreaker being alphabetical order.
func firstAvailableNode(dependsOn CharacterMap, nodeAvailableTime map[byte]int) byte {
	var available byteSlice
	for k, v := range dependsOn {
		if len(v) == 0 {
			// Compile a list with equal start time
			if len(available) == 0 || nodeAvailableTime == nil || nodeAvailableTime[k] == nodeAvailableTime[available[0]] {
				available = append(available, k)
			} else if nodeAvailableTime[k] < nodeAvailableTime[available[0]] {
				// If an earlier start time is found, discard the current hits and start anew
				available = make(byteSlice, 1)
				available[0] = k
			}
		}
	}
	// Take the first alphabetically
	sort.Sort(available)
	return available[0]
}

// firstAvailableWorker Returns the index of the worker who will become available first
func firstAvailableWorker(workerCompleteTime []int) int {
	available := 0
	for idx, t := range workerCompleteTime {
		if t < workerCompleteTime[available] {
			available = idx
		}
	}
	return available
}

// NodeTime computes how long a node takes.  Each task takes the baseTime + ordinal(task) time,
// which means task A takes 60+1 = 61s, while task Z takes 60+26=86s.
func NodeTime(node byte, baseTime int) int {
	return baseTime + int(node-byte('A')) + 1
}

// ExecutionTime computes how long it would take to complete the given series of tasks
// given the number of workers.
func ExecutionTime(dep, isDep CharacterMap, workers int, baseTime int) int {
	dependsOn := copyCharacterMap(dep)
	isDependedOnBy := copyCharacterMap(isDep)
	nodeAvailableTime := make(map[byte]int)
	nodeCompleteTime := make(map[byte]int)
	workerCompleteTime := make([]int, workers)
	for i := 0; i < workers; i++ {
		// Workers start as available
		workerCompleteTime[i] = 0
	}

	// All 0-dependency nodes have an available time of 0
	for k, v := range dependsOn {
		if len(v) == 0 {
			nodeAvailableTime[k] = 0
		}
	}

	for len(dependsOn) > 0 {
		node := firstAvailableNode(dependsOn, nodeAvailableTime)
		worker := firstAvailableWorker(workerCompleteTime)
		startTime := max(workerCompleteTime[worker], nodeAvailableTime[node])
		duration := NodeTime(node, baseTime)
		endTime := startTime + duration
		if debug {
			fmt.Printf("Worker %d (avail@%d) taking task %c (avail@%d), complete at %d\n", worker, workerCompleteTime[worker], node, nodeAvailableTime[node], endTime)
		}
		nodeCompleteTime[node] = endTime
		workerCompleteTime[worker] = endTime

		// Clean up the dependencies
		delete(dependsOn, node)
		for v := range isDependedOnBy[node] {
			delete(dependsOn[v], node)
			if len(dependsOn[v]) == 0 {
				nodeAvailableTime[v] = endTime
			}
		}
	}

	endTime := 0
	for _, v := range nodeCompleteTime {
		endTime = max(endTime, v)
	}
	return endTime
}

func main() {
	var input []string
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())

	start := time.Now()
	dependsOn, isDependedOnBy := ReadDependencies(input)
	fmt.Println(ExecutionOrder(dependsOn, isDependedOnBy))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(ExecutionTime(dependsOn, isDependedOnBy, 5, 60))
	fmt.Println(time.Since(start))
}
