package days

import (
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_dive(t *testing.T) {
	// tests
	assert.Equal(t, 150, diveP1(
		[]string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}))
	assert.Equal(t, 900, diveP2(
		[]string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}))

	// solve
	xs, _ := utils.ReadLines("./d2_input.txt")
	assert.Equal(t, 2027977, diveP1(xs))
	assert.Equal(t, 1903644897, diveP2(xs))
}
