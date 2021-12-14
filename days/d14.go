/*
	https://adventofcode.com/2021/day/14
*/

package days

import (
	"math"
	"sort"
	"strings"
)

func getTemplateAndPairs(xs []string) (tmp []rune, m map[string]rune) {
	tmp = []rune(xs[0])
	m = make(map[string]rune, len(xs)-1)
	for i := 2; i < len(xs); i++ {
		line := strings.Split(xs[i], " -> ")
		m[line[0]] = []rune(line[1])[0] //nolint:gocritic //meh
	}
	return tmp, m
}

func extendedPolymerizationP1(xs []string, steps int) int {
	t, ps := getTemplateAndPairs(xs)

	for step := 0; step < steps; step++ {
		stepRes := []rune{t[0]}

		for i := 1; i < len(t); i++ {
			var tmp []rune
			pair := string(stepRes[len(stepRes)-1]) + string(t[i])
			if p, ok := ps[pair]; ok {
				tmp = []rune{p}
			}
			tmp = append(tmp, t[i])
			stepRes = append(stepRes, tmp...)
		}
		t = stepRes
	}

	counts := make(map[rune]int, 26)
	max, min := 0, math.MaxInt
	for _, r := range t {
		counts[r]++
	}
	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return max - min
}

func extendedPolymerizationP2(xs []string, steps int) int {
	t, ps := getTemplateAndPairs(xs)
	pairs := make(map[string]int)
	for i := 0; i < len(t)-1; i++ {
		pairs[string(t[i:i+2])]++
	}

	for i := 0; i < steps; i++ {
		step := make(map[string]int)
		for k, v := range pairs {
			step[k] = v
		}

		for k := range pairs {
			if _, ok := ps[k]; !ok {
				continue
			}
			if step[k] == 0 {
				continue
			}

			step[k] -= pairs[k]

			// injection
			mid := string(ps[k])
			p1 := string(k[0]) + mid
			p2 := mid + string(k[1])

			step[p1] += pairs[k]
			step[p2] += pairs[k]
		}

		// next
		newPairs := make(map[string]int)
		for k, v := range step {
			if v > 0 {
				newPairs[k] = v
			}
		}
		pairs = newPairs
	}

	counts := make(map[uint8]int)
	for k := range pairs {
		counts[k[0]] += pairs[k] / 2
		counts[k[1]] += pairs[k] / 2
	}

	ys, i := make([]int, len(counts)), 0
	for _, v := range counts {
		ys[i] = v
		i++
	}
	sort.Ints(ys)
	var minCh uint8
	for k, v := range counts {
		if v == ys[0] {
			minCh = k
		}
	}

	delta := 0
	if rune(minCh) == t[0] {
		delta++
	}

	res := ys[len(ys)-1] - ys[0]

	return res + delta
}
