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

	head, tail := &point{0, 0}, &point{0, 0}
	visited := map[point]bool{{0, 0}: true}

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		direction := command[0]
		num, _ := strconv.Atoi(command[1])

		switch direction {
		case "U":
			for i := 0; i < num; i++ {
				head.rowID++
				tail.follow(head)
				visited[*tail] = true
			}
		case "D":
			for i := 0; i < num; i++ {
				head.rowID--
				tail.follow(head)
				visited[*tail] = true
			}
		case "R":
			for i := 0; i < num; i++ {
				head.colID++
				tail.follow(head)
				visited[*tail] = true
			}
		case "L":
			for i := 0; i < num; i++ {
				head.colID--
				tail.follow(head)
				visited[*tail] = true
			}
		}
	}

	fmt.Println("part 1:", len(visited))
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
	if (tail.rowID == head.rowID-2 && tail.colID == head.colID+1) || (tail.rowID == head.rowID-1 && tail.colID == head.colID+2) {
		tail.rowID++
		tail.colID--
	}
	// up right
	if (tail.rowID == head.rowID-2 && tail.colID == head.colID-1) || (tail.rowID == head.rowID-1 && tail.colID == head.colID-2) {
		tail.rowID++
		tail.colID++
	}
	// down left
	if (tail.rowID == head.rowID+1 && tail.colID == head.colID-2) || (tail.rowID == head.rowID+2 && tail.colID == head.colID-1) {
		tail.rowID--
		tail.colID++
	}
	// down right
	if (tail.rowID == head.rowID+2 && tail.colID == head.colID+1) || (tail.rowID == head.rowID+1 && tail.colID == head.colID+2) {
		tail.rowID--
		tail.colID--
	}
}
