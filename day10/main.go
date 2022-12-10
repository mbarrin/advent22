package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	tick, register, total := 1, 1, 0

	screen := make([][]string, 6)

	for row := range screen {
		foo := make([]string, 40)
		for i := range foo {
			foo[i] = "."
		}
		screen[row] = foo
	}

	for scanner.Scan() {
		if (tick+20)%40 == 0 {
			total += (register * tick)
		}
		rowID := (tick - 1) / 40
		colID := (tick - 1) % 40
		if colID == register-1 || colID == register || colID == register+1 {
			screen[rowID][colID] = "#"
		}

		line := scanner.Text()
		if strings.HasPrefix(line, "addx") {
			var num int
			fmt.Sscanf(line, "addx %d", &num)
			tick++

			if (tick+20)%40 == 0 {
				total += (register * tick)
			}
			rowID := (tick - 1) / 40
			colID := (tick - 1) % 40

			if colID == register-1 || colID == register || colID == register+1 {
				screen[rowID][colID] = "#"
			}
			register += num
		}
		tick++
	}

	fmt.Println("part 1:", total)

	for _, row := range screen {
		fmt.Println(strings.Join(row, ""))
	}
}
