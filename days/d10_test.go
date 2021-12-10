package days

import (
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_syntaxScoring(t *testing.T) {
	// tests
	xs := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}
	assert.Equal(t, 26397, syntaxScoringP1(xs))
	assert.Equal(t, 288957, syntaxScoringP2(xs))

	// solve
	xs, _ = utils.ReadLines("d10_input.txt")
	assert.Equal(t, 392367, syntaxScoringP1(xs))
	assert.Equal(t, 2192104158, syntaxScoringP2(xs))
}
