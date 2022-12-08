package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = [][]int{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func TestVisibleTrees(t *testing.T) {
	tests := map[string]struct {
		input  [][]int
		output int
	}{
		"1": {input: input, output: 21},
	}

	for name, tc := range tests {
		var seen = &[][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}
		t.Run(name, func(t *testing.T) {
			actual := visibleTrees(tc.input, seen)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestSuitableTree(t *testing.T) {
	tests := map[string]struct {
		input  [][]int
		output int
	}{
		"1": {input: input, output: 8},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := suitableTree(tc.input)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestLookRight(t *testing.T) {
	tests := map[string]struct {
		input  []int
		output int
	}{
		"1": {input: input[1], output: 1},
		"2": {input: input[2], output: 0},
		"3": {input: input[3], output: 1},
	}

	for name, tc := range tests {
		var seen = &[][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}
		t.Run(name, func(t *testing.T) {
			actual := lookRight(tc.input, 0, seen)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestLookLeft(t *testing.T) {
	tests := map[string]struct {
		input  []int
		output int
	}{
		"1": {input: input[1], output: 1},
		"2": {input: input[2], output: 2},
		"3": {input: input[3], output: 0},
	}

	for name, tc := range tests {
		var seen = &[][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}
		t.Run(name, func(t *testing.T) {
			actual := lookLeft(tc.input, 4, seen)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestLookDown(t *testing.T) {
	tests := map[string]struct {
		input  []int
		output int
	}{
		"1": {input: []int{0, 5, 5, 3, 5}, output: 1},
		"2": {input: []int{3, 5, 3, 5, 3}, output: 1},
		"3": {input: []int{7, 1, 3, 4, 9}, output: 0},
	}

	for name, tc := range tests {
		var seen = &[][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}
		t.Run(name, func(t *testing.T) {
			actual := lookDown(tc.input, 0, seen)
			assert.Equal(t, tc.output, actual)
		})
	}
}

func TestLookUp(t *testing.T) {
	tests := map[string]struct {
		input  []int
		output int
	}{
		"1": {input: []int{0, 5, 5, 3, 5}, output: 0},
		"2": {input: []int{3, 5, 3, 5, 3}, output: 1},
		"3": {input: []int{7, 1, 3, 4, 9}, output: 0},
	}

	for name, tc := range tests {
		var seen = &[][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}
		t.Run(name, func(t *testing.T) {
			actual := lookUp(tc.input, 4, seen)
			assert.Equal(t, tc.output, actual)
		})
	}
}
