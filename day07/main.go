package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	parent *node
	path   string
	dirs   []*node
	files  map[string]int
}

type valid struct {
	node *node
	size int
}

var cdIn = regexp.MustCompile(`^\$ cd [a-zA-Z]+$`).MatchString
var cdOut = regexp.MustCompile(`^\$ cd ..$`).MatchString
var cdRoot = regexp.MustCompile(`^\$ cd /$`).MatchString
var ls = regexp.MustCompile(`^\$ ls$`).MatchString
var file = regexp.MustCompile(`^\d+ [a-zA-Z](\.[a-zA-Z]+)?`).MatchString
var dir = regexp.MustCompile(`^dir [a-zA-Z]$`).MatchString

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	tree := parseInput(lines)

	validNodes := tree.validNodes([]*valid{})

	var total int
	for _, x := range validNodes {
		total += x.size
	}

	fmt.Println("part 1:", total)

}

func parseInput(lines []string) *node {
	cwd := ""
	root := newNode(nil, "/")
	current := root

	for i := 0; i < len(lines); i++ {
		cmd := lines[i]
		switch {
		case cdIn(cmd):
			cwd = in(cmd, cwd)
			new := newNode(current, cwd)
			current.addDir(new)
			current = new
			fmt.Println(cwd)
		case cdOut(cmd):
			cwd = out(cwd)
			current = current.parent
			fmt.Println(cwd)
		case cdRoot(cmd):
			cwd = "/"
			current = root
			fmt.Println(cwd)
		case ls(cmd):
			i += list(lines[i+1:], current)
		}
	}
	return root
}

func in(cmd, cwd string) string {
	parts := strings.Split(cmd, " ")
	return cwd + parts[2] + "/"
}

func out(cwd string) string {
	current := strings.Split(cwd, "/")
	return strings.Join(current[:len(current)-2], "/") + "/"
}

func list(lines []string, current *node) int {
	numFiles := 0
	for _, x := range lines {
		if file(x) {
			numFiles++
			fileDetail := strings.Split(x, " ")
			size, _ := strconv.Atoi(fileDetail[0])
			name := fileDetail[1]
			current.addFile(name, size)
		} else if dir(x) {
			numFiles++
		} else {
			return numFiles
		}

	}
	return numFiles
}

func (n *node) addDir(dir *node) {
	n.dirs = append(n.dirs, dir)
}

func (n *node) addFile(name string, size int) {
	n.files[name] = size
}

func (n *node) size() int {
	total := 0
	for _, size := range n.files {
		total += size
	}

	return total
}

func (n *node) sizeOfTree() int {
	size := n.size()
	for _, x := range n.dirs {
		size += x.sizeOfTree()
	}
	return size
}

func (n *node) validNodes(v []*valid) []*valid {
	size := n.sizeOfTree()
	if size <= 100000 {
		v = append(v, &valid{node: n, size: size})
	}
	for _, x := range n.dirs {
		v = x.validNodes(v)
	}

	return v
}

func newNode(parent *node, path string) *node {
	return &node{
		parent: parent,
		path:   path,
		dirs:   []*node{},
		files:  map[string]int{},
	}
}
