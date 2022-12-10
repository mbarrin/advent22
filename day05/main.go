package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type stack []string

var transform = map[int]int{
	1: 0, 5: 1, 9: 2, 13: 3, 17: 4,
	21: 5, 25: 6, 29: 7, 33: 8,
}

var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	lines := []string{}
	floor := [9]stack{}
	newFloor := [9]stack{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines[0:8] {
		boxes := strings.Split(line, "")

		for i := 1; i < len(boxes); i += 4 {
			char := boxes[i]

			if isLetter(char) {
				floor[transform[i]].pushBottom(char)
				newFloor[transform[i]].pushBottom(char)
			}
		}
	}

	for _, line := range lines[10:] {
		var count, from, to int

		fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)

		for i := 0; i < count; i++ {
			top := floor[from-1].popTop()
			floor[to-1].pushTop(top)
		}

		top := newFloor[from-1].popMultiTop(count)
		newFloor[to-1].pushMultiTop(top)
	}

	partOne := []string{}
	partTwo := []string{}

	for i := range floor {
		partOne = append(partOne, floor[i][len(floor[i])-1])
		partTwo = append(partTwo, newFloor[i][len(floor[i])-1])
	}

	fmt.Println("part 1: ", strings.Join(partOne, ""))
	fmt.Println("part 2: ", strings.Join(partTwo, ""))
}

// remove from end
func (s *stack) popTop() string {
	e := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return e
}

// add to end
func (s *stack) pushTop(e string) {
	*s = append(*s, e)
}

// add to start
func (s *stack) pushBottom(e string) {
	*s = append([]string{e}, *s...)
}

// remove many from end
func (s *stack) popMultiTop(count int) []string {
	e := (*s)[len(*s)-count : len(*s)]
	*s = (*s)[0 : len(*s)-count]
	return e
}

// add many to end
func (s *stack) pushMultiTop(e []string) {
	*s = append(*s, e...)
}
