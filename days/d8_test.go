package days

import (
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_sevenSegmentSearch(t *testing.T) {
	// tests
	xs := []string{
		"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
	}
	assert.Equal(t, 26, sevenSegmentSearchP1(xs))

	assert.Equal(t, 5353, sevenSegmentSearchP2(
		[]string{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"}),
	)
	assert.Equal(t, 8394, sevenSegmentSearchP2(
		[]string{"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe"}),
	)
	assert.Equal(t, 61229, sevenSegmentSearchP2(xs))

	// solve
	xs, _ = utils.ReadLines("./d8_input.txt")
	assert.Equal(t, 349, sevenSegmentSearchP1(xs))
	assert.Equal(t, 1070957, sevenSegmentSearchP2(xs))
}

func Test_segmentToInt(t *testing.T) {
	ds := map[rune]rune{
		'a': 'a',
		'b': 'b',
		'c': 'c',
		'd': 'd',
		'e': 'e',
		'f': 'f',
		'g': 'g',
	}
	assert.Equal(t, 1, segmentToInt(ds, "fc"))
	assert.Equal(t, 2, segmentToInt(ds, "gedca"))
	assert.Equal(t, 3, segmentToInt(ds, "acdgf"))
	assert.Equal(t, 4, segmentToInt(ds, "bcdf"))
	assert.Equal(t, 5, segmentToInt(ds, "adgbf"))
	assert.Equal(t, 6, segmentToInt(ds, "abdfeg"))
	assert.Equal(t, 7, segmentToInt(ds, "acf"))
	assert.Equal(t, 8, segmentToInt(ds, "acbdefg"))
	assert.Equal(t, 9, segmentToInt(ds, "abdcfg"))

	ds = map[rune]rune{
		'a': 'd',
		'b': 'e',
		'c': 'a',
		'd': 'f',
		'e': 'g',
		'f': 'b',
		'g': 'c',
	}
	assert.Equal(t, 8, segmentToInt(ds, "acedgfb"))
	assert.Equal(t, 5, segmentToInt(ds, "cdfbe"))
	assert.Equal(t, 2, segmentToInt(ds, "gcdfa"))
	assert.Equal(t, 3, segmentToInt(ds, "fbcad"))
	assert.Equal(t, 7, segmentToInt(ds, "dab"))
	assert.Equal(t, 9, segmentToInt(ds, "cefabd"))
	assert.Equal(t, 6, segmentToInt(ds, "cdfgeb"))
	assert.Equal(t, 4, segmentToInt(ds, "eafb"))
	assert.Equal(t, 0, segmentToInt(ds, "cagedb"))
	assert.Equal(t, 1, segmentToInt(ds, "ab"))
}

func Test_contain(t *testing.T) {
	_1 := map[rune]bool{
		'c': true,
		'f': true,
	}

	_7 := map[rune]bool{
		'a': true,
		'c': true,
		'f': true,
	}

	_4 := map[rune]bool{
		'b': true,
		'c': true,
		'd': true,
		'f': true,
	}

	_0 := map[rune]bool{
		'a': true,
		'b': true,
		'c': true,
		'e': true,
		'f': true,
		'g': true,
	}

	assert.True(t, contain(_7, _1))
	assert.True(t, contain(_4, _1))
	assert.True(t, contain(_0, _1))
	assert.True(t, contain(_0, _7))
	assert.False(t, contain(_0, _4))

}
