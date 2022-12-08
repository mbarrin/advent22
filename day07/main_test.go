package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}
var root = &node{
	path: "/", parent: nil,
	files: map[string]int{
		"b.txt": 14848514,
		"c.dat": 8504156,
	},
	dirs: []*node{},
}

var a = &node{
	path: "/a/", parent: root,
	files: map[string]int{
		"f":     29116,
		"g":     2557,
		"h.lst": 62596,
	},
	dirs: []*node{},
}

var d = &node{
	path: "/d/", parent: root,
	files: map[string]int{
		"j":     4060174,
		"d.log": 8033020,
		"d.ext": 5626152,
		"k":     7214296,
	},
	dirs: []*node{},
}

var e = &node{
	path: "/a/e/", parent: a,
	files: map[string]int{"i": 584},
	dirs:  []*node{},
}

func TestParseInput(t *testing.T) {
	root.dirs = []*node{a, d}
	a.dirs = []*node{e}

	tests := map[string]struct {
		input  []string
		output *node
	}{
		"sample": {input: input, output: root},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := parseInput(tc.input)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestIn(t *testing.T) {
	tests := map[string]struct {
		cmd    string
		cwd    string
		output string
	}{
		"one":   {cwd: "/", cmd: "$ cd a", output: "/a/"},
		"two":   {cwd: "/a/", cmd: "$ cd b", output: "/a/b/"},
		"three": {cwd: "/a/", cmd: "$ cd bsss", output: "/a/bsss/"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := in(tc.cmd, tc.cwd)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestOut(t *testing.T) {
	tests := map[string]struct {
		cwd    string
		output string
	}{
		"one":   {cwd: "/a/", output: "/"},
		"two":   {cwd: "/a/b/", output: "/a/"},
		"three": {cwd: "/a/b/c/", output: "/a/b/"},
		"four":  {cwd: "/a/b/csss/", output: "/a/b/"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := out(tc.cwd)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestList(t *testing.T) {
	lsInputOne := []string{"4060174 j ", "8033020 d.log", "5626152 d.ext", "7214296 k "}
	lsInputTwo := []string{"584 i", "$ cd .."}
	root := &node{path: "/", files: map[string]int{}}
	child := &node{parent: root, path: "/a", files: map[string]int{}}

	tests := map[string]struct {
		node   *node
		input  []string
		output int
	}{
		"one": {input: lsInputOne, output: 4, node: root},
		"two": {input: lsInputTwo, output: 1, node: child},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := list(tc.input, tc.node)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestSize(t *testing.T) {
	root.dirs = []*node{a, d}
	a.dirs = []*node{e}

	tests := map[string]struct {
		input  *node
		output int
	}{
		"/": {input: root, output: 23352670},
		"a": {input: a, output: 94269},
		"d": {input: d, output: 24933642},
		"e": {input: e, output: 584},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := tc.input.size()
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestSizeOfTree(t *testing.T) {
	root.dirs = []*node{a, d}
	a.dirs = []*node{e}

	tests := map[string]struct {
		input  *node
		output int
	}{
		"/": {input: root, output: 48381165},
		"a": {input: a, output: 94853},
		"d": {input: d, output: 24933642},
		"e": {input: e, output: 584},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := tc.input.sizeOfTree()
			assert.Equal(t, tc.output, actual)
		})

	}
}

func TestValidNodes(t *testing.T) {
	tests := map[string]struct {
		input  *node
		output map[string]int
	}{
		"/": {input: root, output: map[string]int{a.path: 94853, e.path: 584}},
		"a": {input: a, output: map[string]int{a.path: 94853, e.path: 584}},
		"d": {input: d, output: map[string]int{}},
		"e": {input: e, output: map[string]int{e.path: 584}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := tc.input.validNodes(map[string]int{}, 100000)
			assert.Equal(t, tc.output, actual)
		})
	}

	tests = map[string]struct {
		input  *node
		output map[string]int
	}{
		"/ no limit": {input: root, output: map[string]int{root.path: 48381165, a.path: 94853, d.path: 24933642, e.path: 584}},
		"a no limit": {input: a, output: map[string]int{a.path: 94853, e.path: 584}},
		"d no limit": {input: d, output: map[string]int{d.path: 24933642}},
		"e no limit": {input: e, output: map[string]int{e.path: 584}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := tc.input.validNodes(map[string]int{}, 0)
			assert.Equal(t, tc.output, actual)
		})

	}

}

func TestNodeToDeleteSize(t *testing.T) {

	tests := map[string]struct {
		input  map[string]int
		output int
	}{
		"/": {input: root.validNodes(map[string]int{}, 0), output: 24933642},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := nodeToDeleteSize(tc.input)
			assert.Equal(t, tc.output, actual)
		})
	}
}
