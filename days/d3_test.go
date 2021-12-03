package days

import (
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_binaryDiagnosticP1(t *testing.T) {
	// tests
	assert.Equal(t, 198, binaryDiagnosticP1(
		[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}))
	assert.Equal(t, 230, binaryDiagnosticP2(
		[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}))

	// solve
	xs, _ := utils.ReadLines("./d3_input.txt")
	assert.Equal(t, 1997414, binaryDiagnosticP1(xs))
	assert.Equal(t, 1032597, binaryDiagnosticP2(xs))
}
