package days

import (
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_lanternfish(t *testing.T) {
	// tests
	assert.Equal(t, 26, lanternfishP1([]int{3, 4, 3, 1, 2}, 18))
	assert.Equal(t, 5934, lanternfishP1([]int{3, 4, 3, 1, 2}, 80))

	assert.Equal(t, 26, lanternfishP2([]int{3, 4, 3, 1, 2}, 18))
	assert.Equal(t, 5934, lanternfishP2([]int{3, 4, 3, 1, 2}, 80))
	assert.Equal(t, 26984457539, lanternfishP2([]int{3, 4, 3, 1, 2}, 256))

	// solve
	xs, _ := utils.ReadIntsLine("./d6_input.txt")
	p2xs := make([]int, len(xs))
	copy(p2xs, xs)

	assert.Equal(t, 352872, lanternfishP1(xs, 80))
	assert.Equal(t, 1604361182149, lanternfishP2(p2xs, 256))
}
