package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	rowID int
	colID int
}

var lowestRow = 0

func main() {
	file, _ := os.Open("sample.txt")
	scanner := bufio.NewScanner(file)

	cave := map[point]bool{}

	for scanner.Scan() {
		parseLine(scanner.Text(), cave)
	}

	fmt.Println("part 1:", dropAllSand(cave))
}

func dropAllSand(cave map[point]bool) int {
	total := 0

	for {
		sand := dropSand(point{rowID: 0, colID: 500}, cave)
		if sand == nil {
			return total
		}
		total++
		cave[*sand] = true
	}
}

func dropSand(sand point, cave map[point]bool) *point {
	for {
		if sand.rowID >= lowestRow {
			return nil
		}

		if (cave[point{sand.rowID + 1, sand.colID}] && cave[point{sand.rowID + 1, sand.colID - 1}] && cave[point{sand.rowID + 1, sand.colID + 1}]) {
			return &sand
		}

		if !cave[point{sand.rowID + 1, sand.colID}] {
			sand.rowID++
		} else if !cave[point{sand.rowID + 1, sand.colID - 1}] {
			sand.rowID++
			sand.colID--
		} else if !cave[point{sand.rowID + 1, sand.colID + 1}] {
			sand.rowID++
			sand.colID++
		}
	}
}

func parseLine(line string, cave map[point]bool) {
	coords := strings.Split(line, " -> ")
	for i := 0; i < len(coords)-1; i++ {
		coord := strings.Split(coords[i], ",")
		colID, _ := strconv.Atoi(coord[0])
		rowID, _ := strconv.Atoi(coord[1])
		first := point{rowID: rowID, colID: colID}

		if rowID > lowestRow {
			lowestRow = rowID
		}

		coord = strings.Split(coords[i+1], ",")
		colID, _ = strconv.Atoi(coord[0])
		rowID, _ = strconv.Atoi(coord[1])
		second := point{rowID: rowID, colID: colID}

		if rowID > lowestRow {
			lowestRow = rowID
		}

		drawLine(first, second, cave)
	}
}

func drawLine(start, end point, cave map[point]bool) {
	if start.rowID == end.rowID {
		if start.colID < end.colID {
			for i := start.colID; i <= end.colID; i++ {
				cave[point{rowID: start.rowID, colID: i}] = true
			}
		} else {
			for i := end.colID; i <= start.colID; i++ {
				cave[point{rowID: start.rowID, colID: i}] = true
			}
		}
	} else {
		if start.rowID < end.rowID {
			for i := start.rowID; i <= end.rowID; i++ {
				cave[point{rowID: i, colID: start.colID}] = true
			}
		} else {
			for i := end.rowID; i <= start.rowID; i++ {
				cave[point{rowID: i, colID: start.colID}] = true
			}
		}
	}
}
