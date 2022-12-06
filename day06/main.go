package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	chars := []rune{}

	for scanner.Scan() {
		chars = append(chars, []rune(scanner.Text())...)
	}
	fmt.Println("part 1:", marker(chars, 4))
	fmt.Println("part 2:", marker(chars, 14))
}

func marker(line []rune, count int) int {
	for i := 0; i < len(line)-count; i++ {
		set := map[rune]int{}

		for j := i; j < i+count; j++ {
			set[line[j]]++
			if keyLength(set) == count {
				return (i + count)
			}
		}
	}
	return 0
}

func keyLength(set map[rune]int) int {
	var keys []rune
	for k := range set {
		keys = append(keys, k)
	}
	return len(keys)
}
