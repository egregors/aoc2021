package days

import (
	"strconv"
	"strings"
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_giantSquid(t *testing.T) {
	// tests
	nums := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	// 													^ (idx: 11)
	boards := [][][]int{
		{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		},
		{
			{3, 15, 0, 2, 22},
			{9, 18, 13, 17, 5},
			{19, 8, 7, 25, 23},
			{20, 11, 10, 24, 4},
			{14, 21, 16, 12, 6},
		},
		{
			{14, 21, 17, 24, 4},
			{10, 16, 15, 9, 19},
			{18, 8, 23, 26, 20},
			{22, 11, 13, 6, 5},
			{2, 0, 12, 3, 7},
		},
	}
	assert.Equal(t, 4512, giantSquidP1(nums, boards))
	assert.Equal(t, 1924, giantSquidP2(nums, boards))

	// solve
	xs, _ := utils.ReadLines("./d4_input.txt")
	nums = nil
	for _, s := range strings.Split(xs[0], ",") {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}

	boards = [][][]int{}
	var b [][]int //nolint:prealloc // i need nil here
	b = nil
	for _, line := range xs[1:] {
		if line == "" {
			if b != nil {
				boards = append(boards, b)
			}
			b = [][]int{}
			continue
		}
		var l []int
		for _, ch := range strings.Split(line, " ") {
			n, err := strconv.Atoi(ch)
			if err == nil {
				l = append(l, n)
			}
		}
		b = append(b, l)
	}
	boards = append(boards, b)

	assert.Equal(t, 35670, giantSquidP1(nums, boards))
	assert.Equal(t, 22704, giantSquidP2(nums, boards))
}
