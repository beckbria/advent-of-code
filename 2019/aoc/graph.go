package aoc

import(
	"fmt"
)

// Node represents a node in a tree structure with attached data
type TreeNode struct {
	Children []TreeNode
	Data interface{}
}

func (n *TreeNode) Print() {
	n.printTree(0)
}

func (n *TreeNode) printTree(depth int) {
	padding := ""
	for i := 0; i < depth; i++ {
		padding = padding + "    "
	}
	fmt.Print(padding + " (")
	fmt.Print(n.Data)
	fmt.Print(")\n")
	for _, c := range n.Children {
		c.printTree(depth+1)
	}
}