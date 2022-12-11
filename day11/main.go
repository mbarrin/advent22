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
	lcd           int
	trueIndex     int
	falseIndex    int
	inspections   int
	worryModifier func()
}

type monkeys []*monkey

var itemsRegexp = regexp.MustCompile(`^\s+Starting items: ((\d+(, )?)+)`)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	ms := newMonkeys(lines, false)
	msBig := newMonkeys(lines, true)

	for i := 0; i < 10000; i++ {
		if i < 20 {
			for i, m := range ms {
				for len(m.items) > 0 {
					ms.throw(i)
				}
			}
		}
		for i, m := range msBig {
			for len(m.items) > 0 {
				msBig.throw(i)
			}
		}
	}

	fmt.Println("part 1:", ms.business())
	fmt.Println("part 2:", msBig.business())

}

func (ms monkeys) business() int {
	activity := []int{}
	for _, m := range ms {
		activity = append(activity, m.inspections)
	}

	sort.IntSlice.Sort(activity)

	return activity[len(activity)-1] * activity[len(activity)-2]
}

func (ms monkeys) throw(source int) {
	ms[source].worryModifier()
	ms[source].inspections++

	item := ms[source].items[0]
	ms[source].items = ms[source].items[1:]

	if item%ms[source].test == 0 {
		dest := ms[ms[source].trueIndex]
		dest.items = append(dest.items, item)
	} else {
		dest := ms[ms[source].falseIndex]
		dest.items = append(dest.items, item)
	}
}

func newMonkeys(lines []string, big bool) monkeys {
	ms := monkeys{}
	var lcd = 1

	for i := 0; i < len(lines); i += 7 {
		m := newMonkey(lines[i:i+7], big)
		ms = append(ms, m)
		lcd *= m.test
	}

	for _, m := range ms {
		m.lcd = lcd
	}

	return ms
}

func newMonkey(info []string, big bool) *monkey {
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
		m.worryModifier = func() {
			m.items[0] += num
			if !big {
				m.items[0] /= 3
			}
			m.items[0] %= m.lcd
		}
	} else if operation[len(operation)-2] == "*" {
		if strings.HasSuffix(info[2], "* old") {
			m.worryModifier = func() {
				m.items[0] *= m.items[0]
				if !big {
					m.items[0] /= 3
				}
				m.items[0] %= m.lcd
			}
		} else {
			num, _ := strconv.Atoi(operation[len(operation)-1])
			m.worryModifier = func() {
				m.items[0] *= num
				if !big {
					m.items[0] /= 3
				}
				m.items[0] %= m.lcd
			}
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
