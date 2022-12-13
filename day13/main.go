package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type node struct {
	location point
	height   rune
	gScore   float64
	fScore   float64
}

type point struct {
	rowID int
	colID int
}

type openSet map[*node]bool

var allNodes = [][]*node{}
var maxRowID, maxColID int

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	start, end := &node{}, &node{}

	startingPoint := []*node{}

	rowID := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []*node{}
		for colID, height := range strings.Split(line, "") {
			nn := newNode(rowID, colID, rune(height[0]))
			if height == "S" {
				nn.height = 'a'
				start = nn
			} else if height == "E" {
				nn.height = 'z'
				end = nn
			}
			if nn.height == 'a' {
				startingPoint = append(startingPoint, nn)
			}
			row = append(row, nn)
		}
		allNodes = append(allNodes, row)
		rowID++
	}

	maxRowID = len(allNodes)
	maxColID = len(allNodes[0])

	fmt.Println("part 1:", shortestDistance(start, end))

	allStarts := []float64{}
	for _, node := range startingPoint {
		allStarts = append(allStarts, shortestDistance(node, end))
	}

	sort.Float64s(allStarts)

	fmt.Println("part 2:", allStarts[0])

}

func shortestDistance(start, end *node) float64 {
	sld := start.distance(end)

	openSet := openSet{start: true}
	start.gScore = 0
	start.fScore = float64(sld)
	openSet[start] = true

	for len(openSet) > 0 {
		current := openSet.lowestFScore()
		if current == end {
			return current.gScore
		}
		delete(openSet, current)
		// check up
		if current.location.colID-1 >= 0 {
			up := allNodes[current.location.rowID][current.location.colID-1]
			current.checkNeighbour(up, openSet, end)
		}
		// check left
		if current.location.rowID-1 >= 0 {
			left := allNodes[current.location.rowID-1][current.location.colID]
			current.checkNeighbour(left, openSet, end)
		}
		// check down
		if current.location.colID+1 < maxColID {
			down := allNodes[current.location.rowID][current.location.colID+1]
			current.checkNeighbour(down, openSet, end)
		}
		// check right
		if current.location.rowID+1 < maxRowID {
			right := allNodes[current.location.rowID+1][current.location.colID]
			current.checkNeighbour(right, openSet, end)
		}
	}

	return math.Inf(1)
}

func (os openSet) lowestFScore() *node {
	var lowestNode *node
	lowest := math.Inf(1)
	for node := range os {
		if node.fScore <= lowest {
			lowestNode = node
			lowest = node.fScore
		}
	}
	return lowestNode
}

func (n *node) distance(o *node) float64 {
	return math.Abs(float64(n.location.rowID-o.location.rowID)) + math.Abs(float64(n.location.colID-o.location.colID))
}

func (n *node) checkNeighbour(o *node, openSet map[*node]bool, end *node) {
	if o.height <= n.height || o.height == n.height+1 {
		tentativeGScore := n.gScore + 1
		if tentativeGScore < o.gScore {
			o.gScore = tentativeGScore
			o.fScore = tentativeGScore + float64(o.distance(end))
			if !openSet[o] {
				openSet[o] = true
			}
		}
	}
}

func newNode(rowID, colID int, height rune) *node {
	n := node{
		location: point{
			rowID: rowID,
			colID: colID,
		},
		height: height,
		gScore: math.Inf(1),
		fScore: math.Inf(1),
	}
	return &n
}
