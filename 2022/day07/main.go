package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// In retrospect it totally may not have been necessary to construct
// a whole file system tree to solve this problem but ok

type node struct {
	name     string
	dsize    int
	ntype    string
	parent   *node
	children map[string]*node
}

func (n *node) walk(c func(*node) bool) []node {
	// walk tree and keep nodes matching some condition c
	// e.g. walk(head, func(n node) { node.value > 1000000} )
	out := []node{}
	if c(n) {
		out = append(out, *n)
	}
	for _, child := range n.children {
		out = append(out, child.walk(c)...)
	}
	return out
}

func (n *node) sizeOf() int {
	if n.ntype == "file" || n.dsize > 0 {
		return n.dsize
	}

	var dirSize int
	for _, child := range n.children {
		dirSize += child.sizeOf()
	}
	n.dsize = dirSize
	return dirSize
}

func parse(nav []string) *node {
	// it starts with `cd /` so we know we are starting at the root node

	var top = node{name: "/", children: make(map[string]*node)}
	var currentPointer = &top

	for _, line := range nav {
		// if it's an ls command, actually we can just continue reading lines
		// if it's a cd command, we have to move
		if line == "$ cd /" {
			currentPointer = &top
		} else if line == "$ cd .." {
			currentPointer = currentPointer.parent
		} else if strings.HasPrefix(line, "$ cd") {
			var dirName string
			fmt.Sscanf(line, "$ cd %s", &dirName)
			currentPointer = currentPointer.children[dirName]
		} else if line == "$ ls" {
			continue
		} else if strings.HasPrefix(line, "dir") {
			var dirName string
			fmt.Sscanf(line, "dir %s", &dirName)
			child := node{name: dirName, parent: currentPointer, children: make(map[string]*node), ntype: "dir"}
			currentPointer.children[dirName] = &child
		} else {
			// I think the only other option is that it's a file...
			var fname string
			var fsize int
			fmt.Sscanf(line, "%d %s", &fsize, &fname)
			currentPointer.children[fname] = &node{name: fname, dsize: fsize, children: make(map[string]*node), ntype: "file"}
		}

	}

	// return the pointer to the top of the file directory
	return &top
}

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fsHead := parse(strings.Split(string(contents), "\n"))
	dirs := fsHead.walk(func(n *node) bool { return n.ntype == "dir" && n.sizeOf() < 100000 })
	var total int
	for _, dir := range dirs {
		total += dir.sizeOf()
	}
	fmt.Printf("Part One: Sum of all directories of size < 100000: %d\n", total)

	// Part Two
	// current free space
	var totalSpace = 70000000
	var currentSpace = totalSpace - fsHead.sizeOf()
	var neededDelta = 30000000 - currentSpace

	dirs = fsHead.walk(func(n *node) bool { return n.ntype == "dir" && n.sizeOf() >= neededDelta })
	dirSizes := []int{}
	for _, dir := range dirs {
		dirSizes = append(dirSizes, dir.sizeOf())
	}
	sort.Ints(dirSizes)
	fmt.Printf("Part Two: Size of the smallest directory we can delete to free the required amount of space is %d\n", dirSizes[0])

}
