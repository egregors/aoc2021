package days

import (
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_sonarSweepP1(t *testing.T) {
	// tests
	assert.Equal(t, 7, sonarSweepP1(
		[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
	))
	assert.Equal(t, 5, sonarSweepP2(
		[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
	))

	// solve
	xs, _ := utils.ReadInts("./d1_input.txt")
	assert.Equal(t, 1713, sonarSweepP1(xs))
	assert.Equal(t, 1734, sonarSweepP2(xs))
}
