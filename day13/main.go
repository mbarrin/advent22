package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type info struct {
	values []rune
	child  *info
}

type complete []*info

func main() {
	file, _ := os.Open("sample.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
	}
}

func buildComplete(line string) complete {
	complete := complete{}

	values := []rune(line)

	for i := 0; i < len(line); i++ {
		r := rune(line[i])
		if r == '[' {
			tmp, count := buildInfo(values[i+1:])
			complete = append(complete, tmp)
			i += count
		} else if unicode.IsNumber(r) {
			new := newInfo()
			new.values = append(new.values, r)
			complete = append(complete, new)
		} else {
			continue
		}
	}

	return complete
}

func buildInfo(runes []rune) (*info, int) {
	info := newInfo()

	i := 0
	for i = 0; i < len(runes); i++ {
		r := rune(runes[i])
		if r == '[' {
			tmp, count := buildInfo(runes[i+1:])
			info.child = tmp
			i += count + 1
		} else if r == ']' {
			return info, i
		} else if unicode.IsNumber(r) {
			info.values = append(info.values, r)
		} else {
			continue
		}
	}
	return info, i
}

func (left complete) compare(right complete) bool {
	if len(right) > len(left) {
		return false
	}

	for _, x := range left {
		fmt.Println(x)
	}
	for _, x := range right {
		fmt.Println(x)
	}

	for i := 0; i < len(left); i++ {
		for j := 0; j < len(left[i].values); j++ {
			if left[i].values[j] > right[i].values[j] {
				return false
			}
		}
	}

	return true
}

func newInfo() *info {
	return &info{values: []rune{}, child: nil}
}
