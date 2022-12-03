package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	lookup := map[string]int{
		"B X": 1, "C Y": 2, "A Z": 3,
		"A X": 4, "B Y": 5, "C Z": 6,
		"C X": 7, "A Y": 8, "B Z": 9,
	}

	lookupTwo := map[string]int{
		"B X": 1, "C X": 2, "A X": 3,
		"A Y": 4, "B Y": 5, "C Y": 6,
		"C Z": 7, "A Z": 8, "B Z": 9,
	}
	var score, scoreTwo int

	for scanner.Scan() {
		line := scanner.Text()

		score += lookup[line]
		scoreTwo += lookupTwo[line]
	}

	fmt.Println("part 1:", score)
	fmt.Println("part 2:", scoreTwo)
}
