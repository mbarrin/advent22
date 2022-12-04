package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	total := 0
	totalTwo := 0

	for scanner.Scan() {
		line := scanner.Text()

		if contains(line) {
			total++
		}

		if overlap(line) {
			totalTwo++
		}
	}

	fmt.Println("part 1:", total)
	// 640 is too low
	fmt.Println("part 2:", totalTwo)
}

func contains(line string) bool {
	var leftMin, leftMax, rightMin, rightMax int
	fmt.Sscanf(line, "%d-%d,%d-%d", &leftMin, &leftMax, &rightMin, &rightMax)

	if (leftMin <= rightMin && rightMax <= leftMax) || (rightMin <= leftMin && leftMax <= rightMax) {
		return true
	}

	return false
}

func overlap(line string) bool {
	var leftMin, leftMax, rightMin, rightMax int
	fmt.Sscanf(line, "%d-%d,%d-%d", &leftMin, &leftMax, &rightMin, &rightMax)
	if (leftMin <= rightMin && rightMax <= leftMax) ||
		(rightMin <= leftMin && leftMax <= rightMax) ||
		(leftMin <= rightMin && leftMax >= rightMin && leftMax <= rightMax) ||
		(rightMin <= leftMin && rightMax >= leftMin && rightMax <= leftMax) {
		return true
	}

	return false
}
