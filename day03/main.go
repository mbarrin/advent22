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

	dupe, badge := complete(lines)
	fmt.Println("part 1:", dupe)
	fmt.Println("part 2:", badge)
}

func complete(items []string) (rune, rune) {
	var dupe, badge rune

	for i := 0; i < len(items); i += 3 {
		dupe += dupes(items[i : i+3])
		badge += badges(items[i : i+3])
	}
	return dupe, badge
}

func dupes(items []string) rune {
	var num rune

	for _, item := range items {
		stuff := strings.Split(item, "")
		first := stuff[:len(stuff)/2]
		second := stuff[len(stuff)/2:]

		tmp := doubleIntersection(first, second)
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

		tmp := tripleIntersection(one, two, three)
		num += numberRep([]rune(tmp[0])[0])
	}
	return num
}

func doubleIntersection(first, second []string) []string {
	for _, x := range first {
		for _, y := range second {
			if x == y {
				return []string{x}
			}
		}
	}
	return []string{}
}

func tripleIntersection(first, second, third []string) []string {
	for _, x := range first {
		for _, y := range second {
			for _, z := range third {
				if x == y && y == z {
					return []string{x}
				}
			}
		}
	}
	return []string{}
}

func numberRep(r rune) rune {
	if unicode.IsUpper(r) {
		return r - 38
	} else {
		return r - 96
	}
}
