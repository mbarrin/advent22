package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	rowID, colID int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("part 1:", visited(lines, 2))
	fmt.Println("part 2:", visited(lines, 10))
}

func visited(lines []string, size int) int {
	visited := map[point]bool{{0, 0}: true}
	rope := make([]*point, size)

	for i := range rope {
		rope[i] = &point{0, 0}
	}

	for _, line := range lines {
		command := strings.Split(line, " ")
		direction := command[0]
		num, _ := strconv.Atoi(command[1])

		switch direction {
		case "U":
			for i := 0; i < num; i++ {
				rope[0].rowID++
				for i := range rope[1:] {
					rope[i+1].follow(rope[i])
				}
				visited[*rope[len(rope)-1]] = true
			}
		case "D":
			for i := 0; i < num; i++ {
				rope[0].rowID--
				for i := range rope[1:] {
					rope[i+1].follow(rope[i])
				}
				visited[*rope[len(rope)-1]] = true
			}
		case "R":
			for i := 0; i < num; i++ {
				rope[0].colID++
				for i := range rope[1:] {
					rope[i+1].follow(rope[i])
				}
				visited[*rope[len(rope)-1]] = true
			}
		case "L":
			for i := 0; i < num; i++ {
				rope[0].colID--
				for i := range rope[1:] {
					rope[i+1].follow(rope[i])
				}
				visited[*rope[len(rope)-1]] = true
			}
		}
	}

	return len(visited)
}

func (tail *point) follow(head *point) {
	// up && down
	if tail.rowID == head.rowID {
		if tail.colID == head.colID-2 {
			tail.colID++
		} else if tail.colID == head.colID+2 {
			tail.colID--
		}
	}
	// left && right
	if tail.colID == head.colID {
		if tail.rowID == head.rowID-2 {
			tail.rowID++
		} else if tail.rowID == head.rowID+2 {
			tail.rowID--
		}
	}
	// up left
	if (tail.rowID == head.rowID-2 && tail.colID >= head.colID+1) || (tail.rowID <= head.rowID-1 && tail.colID == head.colID+2) {
		tail.rowID++
		tail.colID--
	}
	// up right
	if (tail.rowID == head.rowID-2 && tail.colID >= head.colID-1) || (tail.rowID <= head.rowID-1 && tail.colID == head.colID-2) {
		tail.rowID++
		tail.colID++
	}
	// down left
	if (tail.rowID >= head.rowID+1 && tail.colID == head.colID-2) || (tail.rowID == head.rowID+2 && tail.colID <= head.colID-1) {
		tail.rowID--
		tail.colID++
	}
	// down right
	if (tail.rowID == head.rowID+2 && tail.colID >= head.colID+1) || (tail.rowID >= head.rowID+1 && tail.colID == head.colID+2) {
		tail.rowID--
		tail.colID--
	}
}
