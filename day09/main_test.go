package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFollow(t *testing.T) {
	tests := map[string]struct {
		inputHead  *point
		inputTail  *point
		outputHead *point
		outputTail *point
	}{
		"above": {inputHead: &point{0, 2}, inputTail: &point{0, 0}, outputHead: &point{0, 2}, outputTail: &point{0, 1}},
		"below": {inputHead: &point{0, -2}, inputTail: &point{0, 0}, outputHead: &point{0, -2}, outputTail: &point{0, -1}},
		"right": {inputHead: &point{2, 0}, inputTail: &point{0, 0}, outputHead: &point{2, 0}, outputTail: &point{1, 0}},
		"left":  {inputHead: &point{-2, 0}, inputTail: &point{0, 0}, outputHead: &point{-2, 0}, outputTail: &point{-1, 0}},

		"diag 1": {inputHead: &point{2, -1}, inputTail: &point{0, 0}, outputHead: &point{2, -1}, outputTail: &point{1, -1}},
		"diag 2": {inputHead: &point{1, -2}, inputTail: &point{0, 0}, outputHead: &point{1, -2}, outputTail: &point{1, -1}},

		"diag 3": {inputHead: &point{2, 1}, inputTail: &point{0, 0}, outputHead: &point{2, 1}, outputTail: &point{1, 1}},
		"diag 4": {inputHead: &point{1, 2}, inputTail: &point{0, 0}, outputHead: &point{1, 2}, outputTail: &point{1, 1}},

		"diag 5": {inputHead: &point{-1, 2}, inputTail: &point{0, 0}, outputHead: &point{-1, 2}, outputTail: &point{-1, 1}},
		"diag 6": {inputHead: &point{-2, 1}, inputTail: &point{0, 0}, outputHead: &point{-2, 1}, outputTail: &point{-1, 1}},

		"diag 7": {inputHead: &point{-2, -1}, inputTail: &point{0, 0}, outputHead: &point{-2, -1}, outputTail: &point{-1, -1}},
		"diag 8": {inputHead: &point{-1, -2}, inputTail: &point{0, 0}, outputHead: &point{-1, -2}, outputTail: &point{-1, -1}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc.inputTail.follow(tc.inputHead)
			assert.Equal(t, tc.outputHead, tc.inputHead)
			assert.Equal(t, tc.outputTail, tc.inputTail)
		})
	}
}
