package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items         []int
	test          int
	trueIndex     int
	falseIndex    int
	inspections   int
	worryModifier func()
}

var itemsRegexp = regexp.MustCompile(`^\s+Starting items: ((\d+(, )?)+)`)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	monkeys := []*monkey{}

	for i := 0; i < len(lines); i += 7 {
		m := newMonkey(lines[i : i+7])
		monkeys = append(monkeys, m)
	}

	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				new, item := m.throw()
				monkeys[new].recieve(item)
			}
		}
	}

	activity := []int{}

	for _, m := range monkeys {
		activity = append(activity, m.inspections)
	}

	sort.IntSlice.Sort(activity)

	fmt.Println("part 1:", activity[len(activity)-1]*activity[len(activity)-2])
}

func (m *monkey) recieve(item int) {
	m.items = append(m.items, item)
}

func (m *monkey) throw() (int, int) {
	m.worryModifier()
	m.inspections++

	item := m.items[0]
	m.items = m.items[1:]

	if item%m.test == 0 {
		return m.trueIndex, item
	} else {
		return m.falseIndex, item
	}
}

func newMonkey(info []string) *monkey {
	m := monkey{items: []int{}}

	// Create the initial list of items
	itemMatches := itemsRegexp.FindStringSubmatch(info[1])
	for _, x := range strings.Split(itemMatches[1], ", ") {
		num, _ := strconv.Atoi(x)
		m.items = append(m.items, num)
	}

	// Create the worry modifier
	operation := strings.Split(info[2], " ")

	if operation[len(operation)-2] == "+" {
		num, _ := strconv.Atoi(operation[len(operation)-1])
		m.worryModifier = func() { m.items[0] += num; m.items[0] /= 3 }
	} else if operation[len(operation)-2] == "*" {
		if strings.HasSuffix(info[2], "* old") {
			m.worryModifier = func() { m.items[0] *= m.items[0]; m.items[0] /= 3 }
		} else {
			num, _ := strconv.Atoi(operation[len(operation)-1])
			m.worryModifier = func() { m.items[0] *= num; m.items[0] /= 3 }
		}
	} else {
		panic("oh no")
	}

	// Create the test value
	fmt.Sscanf(info[3], "  Test: divisible by %d", &m.test)

	// Create the values where to throw items
	fmt.Sscanf(info[4], "    If true: throw to monkey %d", &m.trueIndex)
	fmt.Sscanf(info[5], "    If false: throw to monkey %d", &m.falseIndex)

	return &m
}
