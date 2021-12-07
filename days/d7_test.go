package days

import (
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_theTreacheryOfWhales(t *testing.T) {
	// tests
	xs := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	assert.Equal(t, 37, theTreacheryOfWhalesP1(xs))
	assert.Equal(t, 168, theTreacheryOfWhalesP2(xs))

	// solve
	xs, _ = utils.ReadIntsLine("./d7_input.txt")
	assert.Equal(t, 335330, theTreacheryOfWhalesP1(xs))
	assert.Equal(t, 92439766, theTreacheryOfWhalesP2(xs))
}
