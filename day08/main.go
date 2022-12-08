package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	grid := [][]int{}
	seen := [][]bool{}

	for scanner.Scan() {
		row := []int{}
		seenRow := []bool{}

		tmp := strings.Split(scanner.Text(), "")
		for _, x := range tmp {
			i, _ := strconv.Atoi(x)
			row = append(row, i)
			seenRow = append(seenRow, false)
		}
		grid = append(grid, row)
		seen = append(seen, seenRow)
	}

	fmt.Println("part 1:", visibleTrees(grid, &seen))
	fmt.Println("part 2", suitableTree(grid))
}

func visibleTrees(grid [][]int, seen *[][]bool) int {
	outside := (len(grid[0]) - 1) * 4
	var inside int

	for r := 1; r < len(grid)-1; r++ {
		inside += lookRight(grid[r], r, seen)
		inside += lookLeft(grid[r], r, seen)
	}

	for c := 1; c < len(grid)-1; c++ {
		col := []int{}
		for _, r := range grid {
			col = append(col, r[c])
		}
		inside += lookDown(col, c, seen)
		inside += lookUp(col, c, seen)
	}

	return outside + inside
}

func suitableTree(grid [][]int) int {
	var max int
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid)-1; c++ {
			up, down, left, right := 0, 0, 0, 0
			origin := grid[r][c]

			// up
			for i := r - 1; i >= 0; i-- {
				up++
				if grid[i][c] >= origin {
					break
				}
			}

			// left
			for i := c - 1; i >= 0; i-- {
				left++
				if grid[r][i] >= origin {
					break
				}
			}

			// right
			for i := c + 1; i < len(grid); i++ {
				right++
				if grid[r][i] >= origin {
					break
				}
			}

			// down
			for i := r + 1; i < len(grid); i++ {
				down++
				if grid[i][c] >= origin {
					break
				}
			}

			total := right * left * down * up
			if total > max {
				max = total
			}
		}
	}

	return max
}

func lookRight(row []int, rowID int, seen *[][]bool) int {
	tallest := row[0]
	var visible int

	for c := 1; c < len(row)-1; c++ {
		if row[c] > tallest {
			tallest = row[c]

			if !(*seen)[rowID][c] {
				visible++
				(*seen)[rowID][c] = true
			}
		}
	}
	return visible
}

func lookLeft(row []int, rowID int, seen *[][]bool) int {
	tallest := row[len(row)-1]
	var visible int

	for c := len(row) - 2; c > 0; c-- {
		if row[c] > tallest {
			tallest = row[c]

			if !(*seen)[rowID][c] {
				visible++
				(*seen)[rowID][c] = true
			}
		}
	}
	return visible
}

func lookDown(column []int, columnID int, seen *[][]bool) int {
	tallest := column[0]
	var visible int

	for r := 1; r < len(column)-1; r++ {
		if column[r] > tallest {
			tallest = column[r]

			if !(*seen)[r][columnID] {
				visible++
				(*seen)[r][columnID] = true
			}
		}
	}
	return visible
}

func lookUp(column []int, columnID int, seen *[][]bool) int {
	tallest := column[len(column)-1]
	var visible int

	for r := len(column) - 2; r > 0; r-- {
		if column[r] > tallest {
			tallest = column[r]

			if !(*seen)[r][columnID] {
				visible++
				(*seen)[r][columnID] = true
			}
		}
	}
	return visible
}
