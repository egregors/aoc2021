/*
	https://adventofcode.com/2021/day/8
*/

package days

import (
	"sort"
	"strings"
)

func sevenSegmentSearchP1(xs []string) int {
	res := 0
	for _, x := range xs {
		r := strings.Split(x, " | ")[1]
		for _, w := range strings.Split(r, " ") {
			switch len(w) {
			case 2, 3, 4, 7:
				res++
			}
		}
	}
	return res
}

func sevenSegmentSearchP2(xs []string) int {
	sum := 0
	for _, x := range xs {
		startAndEnd := strings.Split(x, " | ")
		start, end := strings.Split(startAndEnd[0], " "), strings.Split(startAndEnd[1], " ")

		var input, output []string

		for _, w := range start {
			input = append(input, SortString(w))
		}
		for _, w := range end {
			output = append(output, SortString(w))
		}

		// set of all digits
		digits := make(map[string]bool)
		for _, w := range input {
			digits[w] = true
		}

		// nums
		nums := make(map[int]string)

		digitsLen6 := make(map[string]bool)
		digitsLen5 := make(map[string]bool)

		// 1, 7, 4, 8
		for k, v := range digits {
			switch len(k) {
			case 2:
				nums[1] = k
			case 3:
				nums[7] = k
			case 4:
				nums[4] = k
			case 7:
				nums[8] = k

			case 5:
				digitsLen5[k] = v
			case 6:
				digitsLen6[k] = v
			}
		}

		// 6, 9, 0
		for k := range digitsLen6 {
			if len(intersection(k, nums[1])) == 1 {
				nums[6] = k
				continue
			}
			if len(intersection(k, nums[4])) == 4 {
				nums[9] = k
				continue
			}
		}

		for k := range digitsLen6 {
			if k != nums[6] && k != nums[9] {
				nums[0] = k
			}
		}

		// 5, 3, 2
		for k := range digitsLen5 {
			if len(intersection(k, nums[6])) == 5 {
				nums[5] = k
				continue
			}
			if len(intersection(k, nums[1])) == 2 {
				nums[3] = k
			}
		}

		for k := range digitsLen5 {
			if k != nums[3] && k != nums[5] {
				nums[2] = k
			}
		}

		// invert
		segments := make(map[string]int, 9)
		for k, v := range nums {
			segments[v] = k
		}

		rank := 1000
		localSum := 0
		for _, d := range output {
			localSum += segments[d] * rank
			rank /= 10
		}
		sum += localSum
	}

	return sum
}

func intersection(a, b string) string {
	set := make([]rune, 0)
	hash := make(map[rune]bool)

	for _, ch := range a {
		hash[ch] = true
	}

	for _, ch := range b {
		if _, ok := hash[ch]; ok {
			set = append(set, ch)
		}
	}

	return string(set)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

// SortString performs []rune sorting
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
