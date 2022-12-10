package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	parent *node
	path   string
	dirs   []*node
	files  map[string]int
}

var cdIn = regexp.MustCompile(`^\$ cd [a-zA-Z]+$`).MatchString
var cdOut = regexp.MustCompile(`^\$ cd ..$`).MatchString
var cdRoot = regexp.MustCompile(`^\$ cd /$`).MatchString
var file = regexp.MustCompile(`^\d+ [a-zA-Z](\.[a-zA-Z]+)?`).MatchString

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	tree := parseInput(lines)

	validNodes, allNodes := map[string]int{}, map[string]int{}

	tree.validNodes(validNodes, 100000)
	tree.validNodes(allNodes, 0)

	var total int
	for _, s := range validNodes {
		total += s
	}
	fmt.Println("part 1:", total)

	fmt.Println("part 2:", nodeToDeleteSize(allNodes))

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
		case cdOut(cmd):
			cwd = out(cwd)
			current = current.parent
		case cdRoot(cmd):
			cwd = "/"
			current = root
		case file(cmd):
			fileDetail := strings.Split(cmd, " ")
			size, _ := strconv.Atoi(fileDetail[0])
			name := fileDetail[1]
			current.addFile(name, size)
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

func nodeToDeleteSize(nodes map[string]int) int {
	totalSpace := 70000000
	neededSpace := 30000000
	unUsed := totalSpace - nodes["/"]

	canDelete := []int{}
	for _, s := range nodes {
		if unUsed+s >= neededSpace {
			canDelete = append(canDelete, s)
		}
	}
	sort.Ints(canDelete)
	return canDelete[0]
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

func (n *node) validNodes(m map[string]int, limit int) {
	size := n.sizeOfTree()
	if size <= limit || limit == 0 {
		m[n.path] = size
	}
	for _, x := range n.dirs {
		x.validNodes(m, limit)
	}
}

func newNode(parent *node, path string) *node {
	return &node{
		parent: parent,
		path:   path,
		dirs:   []*node{},
		files:  map[string]int{},
	}
}
