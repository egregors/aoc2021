/*
	https://adventofcode.com/2021/day/10
*/

package days

import (
	"sort"
)

func syntaxScoringP1(xs []string) int {
	var corrupted []rune
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	for _, x := range xs {
		var l []rune
	loop:
		for _, ch := range x {
			switch ch {
			case '(', '[', '{', '<':
				l = append(l, ch)
			case ')', ']', '}', '>':
				if l[len(l)-1] != pairs[ch] {
					corrupted = append(corrupted, ch)
					break loop
				} else {
					l = l[:len(l)-1]
				}
			}
		}
	}

	sum := 0
	for _, ch := range corrupted {
		sum += points[ch]
	}

	return sum
}

func syntaxScoringP2(xs []string) int {
	open2close := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	close2open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	points := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	var score []int

loop:
	for _, x := range xs {
		var l []rune
		for _, ch := range x {
			switch ch {
			case '(', '[', '{', '<':
				l = append(l, ch)
			case ')', ']', '}', '>':
				if l[len(l)-1] != open2close[ch] {
					continue loop
				} else {
					l = l[:len(l)-1]
				}
			}
		}
		res := 0
		for i := len(l) - 1; i >= 0; i-- {
			res *= 5
			res += points[close2open[l[i]]]
		}
		score = append(score, res)
	}
	sort.Ints(score)
	return score[len(score)>>1]
}
