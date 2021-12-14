package days

import (
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_extendedPolymerizationP1(t *testing.T) {
	// tests
	sample1 := []string{
		"NNCB",
		"\n",
		"CH -> B",
		"HH -> N",
		"CB -> H",
		"NH -> C",
		"HB -> C",
		"HC -> B",
		"HN -> C",
		"NN -> C",
		"BH -> H",
		"NC -> B",
		"NB -> B",
		"BN -> B",
		"BB -> N",
		"BC -> B",
		"CC -> N",
		"CN -> C",
	}
	assert.Equal(t, 1588, extendedPolymerizationP1(sample1, 10))
	assert.Equal(t, 1588, extendedPolymerizationP2(sample1, 10))
	assert.Equal(t, 2188189693529, extendedPolymerizationP2(sample1, 40))

	// solve
	xs, _ := utils.ReadLines("./d14_input.txt")
	assert.Equal(t, 2027, extendedPolymerizationP1(xs, 10))
	assert.Equal(t, 2265039461737, extendedPolymerizationP2(xs, 40))
}
