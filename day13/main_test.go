package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	tests := map[string]struct {
		left   complete
		right  complete
		output bool
	}{
		//"one": {left: buildComplete("1,1,3,1,1"), right: buildComplete("1,1,5,1,1"), output: true},
		"two": {left: buildComplete("[1],[2,3,4]"), right: buildComplete("[1],4"), output: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := tc.left.compare(tc.right)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestBuildComplete(t *testing.T) {
	tests := map[string]struct {
		input  string
		output complete
	}{
		"one":   {input: "", output: complete{}},
		"two":   {input: "[]", output: complete{{values: []rune{}, child: nil}}},
		"three": {input: "4", output: complete{{values: []rune{'4'}, child: nil}}},
		"four": {
			input: "[1],4",
			output: complete{
				{
					values: []rune{'1'}, child: nil,
				},
				{
					values: []rune{'4'}, child: nil,
				},
			},
		},
		"five": {
			input: "1,[2,[3,[4,[5,6,7]]]],8,9",
			output: complete{
				{
					values: []rune{'1'}, child: nil,
				},
				{
					values: []rune{'2'}, child: &info{
						values: []rune{'3'}, child: &info{
							values: []rune{'4'}, child: &info{
								values: []rune{'5', '6', '7'}, child: nil,
							},
						},
					},
				},
				{
					values: []rune{'8'}, child: nil,
				},
				{
					values: []rune{'9'}, child: nil,
				},
			},
		},
		"six": {
			input: "[[]]",
			output: complete{
				{
					values: []rune{}, child: &info{
						values: []rune{}, child: nil,
					},
				},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := buildComplete(tc.input)
			assert.Equal(t, tc.output, actual)
		})
	}
}
