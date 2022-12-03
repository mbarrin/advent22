package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestDupes(t *testing.T) {
	tests := map[string]struct {
		input  []string
		output rune
	}{
		"default": {input: input, output: 157},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := dupes(tc.input)
			assert.Equal(t, actual, tc.output)
		})
	}
}

func TestBadges(t *testing.T) {
	tests := map[string]struct {
		input  []string
		output rune
	}{
		"default": {input: input, output: 70},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := badges(tc.input)
			assert.Equal(t, actual, tc.output)
		})
	}
}
