package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarker(t *testing.T) {
	tests := map[string]struct {
		input  []rune
		count  int
		output int
	}{
		"one":   {input: []rune("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), count: 4, output: 7},
		"two":   {input: []rune("bvwbjplbgvbhsrlpgdmjqwftvncz"), count: 4, output: 5},
		"three": {input: []rune("nppdvjthqldpwncqszvftbrmjlhg"), count: 4, output: 6},
		"four":  {input: []rune("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), count: 4, output: 10},
		"five":  {input: []rune("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), count: 4, output: 11},
		"six":   {input: []rune("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), count: 14, output: 19},
		"seven": {input: []rune("bvwbjplbgvbhsrlpgdmjqwftvncz"), count: 14, output: 23},
		"eight": {input: []rune("nppdvjthqldpwncqszvftbrmjlhg"), count: 14, output: 23},
		"nine":  {input: []rune("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), count: 14, output: 29},
		"ten":   {input: []rune("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), count: 14, output: 26},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := marker(tc.input, tc.count)
			assert.Equal(t, tc.output, actual)
		})
	}
}
