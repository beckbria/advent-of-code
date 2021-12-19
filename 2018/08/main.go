package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Set to true to get debug spew and a printout of the tree
const debug = false

// Helper functions
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func sum(i []int64) int64 {
	s := int64(0)
	for _, v := range i {
		s += v
	}
	return s
}

// StringToInts Parses a string consisting of space-delimited numbers
func StringToInts(s string) []int64 {
	var i []int64
	for _, v := range strings.Split(s, " ") {
		intVal, err := strconv.ParseInt(v, 10, 64)
		check(err)
		i = append(i, intVal)
	}
	return i
}

// Node represents a node in a tree structure with attached data
type Node struct {
	children []Node
	metadata []int64
}

func makeNode() Node {
	nodeSlice := make([]Node, 0)
	intSlice := make([]int64, 0)
	return Node{children: nodeSlice, metadata: intSlice}
}

// ReadNodes reads the data into a tree
func ReadNodes(input []int64) Node {
	n, _ := readNode(input, 0)
	return n
}

// Reads a node starting at the specified index.  Returns the index after the node was done reading
func readNode(input []int64, index int) (Node, int) {
	n := makeNode()
	numChildren := input[index]
	index++
	numMetadata := input[index]
	index++
	for i := int64(0); i < numChildren; i++ {
		c, idx := readNode(input, index)
		index = idx
		n.children = append(n.children, c)
	}
	for i := int64(0); i < numMetadata; i++ {
		n.metadata = append(n.metadata, input[index])
		index++
	}
	return n, index
}

// MetadataSum sums up all of the metadata information in the tree
func MetadataSum(n Node) int64 {
	s := sum(n.metadata)
	for _, c := range n.children {
		s += MetadataSum(c)
	}
	return s
}

// Value computes the value of a node.  If a node has no children, its value
// is the sum of its metadata.  Otherwise, it's the sum of its metadata-indexed
// child nodes (metadata is 1-indexed)
func Value(n Node) int64 {
	if len(n.children) == 0 {
		return sum(n.metadata)
	} else {
		v := int64(0)
		for _, i := range n.metadata {
			index := i - 1
			if index < int64(len(n.children)) {
				v += Value(n.children[index])
			}
		}
		return v
	}
}

func printTree(n Node, depth int) {
	padding := ""
	for i := 0; i < depth; i++ {
		padding = padding + "    "
	}
	fmt.Print(padding + " (")
	for _, m := range n.metadata {
		fmt.Printf("%d, ", m)
	}
	fmt.Printf(") = %d [", Value(n))
	for _, c := range n.children {
		fmt.Printf("%d, ", Value(c))
	}
	fmt.Print("]\n")

	for _, c := range n.children {
		printTree(c, depth+1)
	}
}

func main() {
	file, err := os.Open("2018/08/input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	check(scanner.Err())
	start := time.Now()
	data := StringToInts(input[0])
	tree := ReadNodes(data)
	fmt.Println(MetadataSum(tree))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(Value(tree))
	fmt.Println(time.Since(start))
	if debug {
		printTree(tree, 0)
	}
}
