/*
	https://adventofcode.com/2021/day/7
*/

package days

import (
	"math"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func fac(n int) int {
	s := 0
	for n > 0 {
		s += n
		n--
	}
	return s
}

func theTreacheryOfWhalesP1(xs []int) int {
	sort.Ints(xs)
	sum := math.MaxInt

	for i := xs[0]; i <= xs[len(xs)-1]; i++ {
		s := 0
		for _, x := range xs {
			s += abs(x - i)
		}
		if s < sum {
			sum = s
		}
	}

	return sum
}

func theTreacheryOfWhalesP2(xs []int) int {
	sort.Ints(xs)
	sum := math.MaxInt

	for i := xs[0]; i <= xs[len(xs)-1]; i++ {
		s := 0
		for _, x := range xs {
			s += fac(abs(x - i))
		}
		if s < sum {
			sum = s
		}
	}

	return sum
}
