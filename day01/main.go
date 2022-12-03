package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var cur int
	count := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			count = append(count, cur)
			cur = 0
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(1)
			}
			cur += num
		}
	}
	count = append(count, cur)
	sort.Sort(sort.Reverse(sort.IntSlice(count)))

	fmt.Printf("part 1: %d\n", count[0])

	var partTwo int
	for _, x := range count[:3] {
		partTwo += x
	}

	fmt.Printf("part 2: %d\n", partTwo)
}
