package days

import (
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_smokeBasinP1(t *testing.T) {
	// tests
	xs := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}
	assert.Equal(t, 15, smokeBasinP1(xs))
	assert.Equal(t, 1134, smokeBasinP2(xs))

	// solve
	ys, _ := utils.ReadIntsLines("./d9_input.txt", "")
	assert.Equal(t, 631, smokeBasinP1(ys))
	assert.Equal(t, 821560, smokeBasinP2(ys))
}
