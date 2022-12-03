package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("part 1:", dupes(lines))
	fmt.Println("part 2:", badges(lines))
}

func dupes(items []string) rune {
	var num rune

	for _, item := range items {
		stuff := strings.Split(item, "")
		first := stuff[:len(stuff)/2]
		second := stuff[len(stuff)/2:]

		tmp, _ := doubleIntersection(first, second)
		num += numberRep([]rune(tmp[0])[0])
	}
	return num
}

func badges(items []string) rune {
	var num rune
	for i := 0; i < len(items); i += 3 {
		one := strings.Split(items[i], "")
		two := strings.Split(items[i+1], "")
		three := strings.Split(items[i+2], "")

		tmp, _ := tripleIntersection(one, two, three)

		num += numberRep([]rune(tmp[0])[0])
	}

	return num
}

func doubleIntersection(first, second []string) ([]string, bool) {
	for _, x := range first {
		for _, y := range second {
			if x == y {
				return []string{x}, true
			}
		}
	}
	return []string{}, false
}

func tripleIntersection(first, second, third []string) ([]string, bool) {
	for _, x := range first {
		for _, y := range second {
			for _, z := range third {
				if x == y && y == z {
					return []string{x}, true
				}
			}
		}
	}
	return []string{}, false
}

func numberRep(r rune) rune {
	if unicode.IsUpper(r) {
		return r - 38
	} else {
		return r - 96
	}
}
