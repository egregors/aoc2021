/*
	https://adventofcode.com/2021/day/5
*/

package days

import (
	"math"
	"strconv"
	"strings"
)

//nolint:revive // meh
type Line struct {
	x1, y1, x2, y2 int
}

func isHorOrVert(l Line) bool {
	return l.x1 == l.x2 || l.y1 == l.y2
}

func isDiagonal(l Line) bool {
	return math.Abs(float64(l.x1-l.x2)) == math.Abs(float64(l.y1-l.y2))
}

func getLine(s string) Line {
	s1 := strings.Split(s, "->")
	l := strings.Split(s1[0], ",")
	r := strings.Split(s1[1], ",")
	x1, _ := strconv.Atoi(strings.TrimSpace(l[0]))
	y1, _ := strconv.Atoi(strings.TrimSpace(l[1]))
	x2, _ := strconv.Atoi(strings.TrimSpace(r[0]))
	y2, _ := strconv.Atoi(strings.TrimSpace(r[1]))
	return Line{x1, y1, x2, y2}
}

func makeFieldAndLines(xs []string, size int) ([][]int, []Line) {
	field := make([][]int, size)
	for i := 0; i < size; i++ {
		field[i] = make([]int, size)
	}

	lines := make([]Line, len(xs))
	for i := 0; i < len(lines); i++ {
		lines[i] = getLine(xs[i])
	}
	return field, lines
}

func countDangerPoints(size int, field [][]int) int {
	danger := 0
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if field[y][x] >= 2 {
				danger++
			}
		}
	}
	return danger
}

func markHorAndVert(lines []Line, field [][]int) {
	for _, l := range lines {
		if isHorOrVert(l) { // hor
			if l.y1 == l.y2 {
				if l.x1 < l.x2 {
					for x := l.x1; x <= l.x2; x++ {
						field[l.y1][x]++
					}
				} else {
					for x := l.x1; x >= l.x2; x-- {
						field[l.y1][x]++
					}
				}
			} else { // vert
				if l.y1 < l.y2 {
					for y := l.y1; y <= l.y2; y++ {
						field[y][l.x1]++
					}
				} else {
					for y := l.y1; y >= l.y2; y-- {
						field[y][l.x1]++
					}
				}
			}
		}
	}
}

func markDiagonal(lines []Line, field [][]int) {
	for _, l := range lines {
		if isDiagonal(l) {
			// 0 0 - 8 8
			if l.x1 < l.x2 { // ->
				if l.y1 < l.y2 { // V
					for y, x := l.y1, l.x1; y <= l.y2 && x <= l.x2; y, x = y+1, x+1 {
						field[y][x]++
					}
				} else { // ^
					for y, x := l.y1, l.x1; y >= l.y2 && x <= l.x2; y, x = y-1, x+1 {
						field[y][x]++
					}
				}
			} else { //  <-
				if l.y1 < l.y2 { // V down
					for y, x := l.y1, l.x1; y <= l.y2 && x >= l.x2; y, x = y+1, x-1 {
						field[y][x]++
					}
				} else { // ^ up
					for y, x := l.y1, l.x1; y >= l.y2 && x >= l.x2; y, x = y-1, x-1 {
						field[y][x]++
					}
				}
			}
		}
	}
}

func hydrothermalVentureP1(xs []string, size int) int {
	field, lines := makeFieldAndLines(xs, size)
	markHorAndVert(lines, field)

	return countDangerPoints(size, field)
}

func hydrothermalVentureP2(xs []string, size int) int {
	field, lines := makeFieldAndLines(xs, size)
	markHorAndVert(lines, field)
	markDiagonal(lines, field)

	return countDangerPoints(size, field)
}
