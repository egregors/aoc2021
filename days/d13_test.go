package days

import (
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_transparentOrigamiP1(t *testing.T) {
	// tests
	sample1 := []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
		"",
		"fold along y=7",
		"fold along x=5",
	}
	assert.Equal(t, 17, transparentOrigamiP1(sample1, 1))

	// solve
	xs, err := utils.ReadLines("./d13_input.txt")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 720, transparentOrigamiP1(xs, 1))
	transparentOrigamiP2(xs)
}
