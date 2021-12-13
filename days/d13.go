/*
	https://adventofcode.com/2021/day/13
*/

package days

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(xs []string) (field [][]rune, folds []map[string]int) {
	var dots [][]int
	dotsEndLine := 0
	maxX, maxY := 0, 0
	for i, x := range xs {
		if x != "" {
			ds := strings.Split(x, ",")
			x, _ := strconv.Atoi(ds[0])
			y, _ := strconv.Atoi(ds[1])
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
			dots = append(dots, []int{x, y})
		} else {
			dotsEndLine = i
			break
		}
	}

	for i := dotsEndLine + 1; i < len(xs); i++ {
		words := strings.Split(xs[i], " ")
		fold := strings.Split(words[2], "=")
		n, _ := strconv.Atoi(fold[1])
		folds = append(folds, map[string]int{fold[0]: n})
	}

	field = make([][]rune, maxY+1)
	for y := 0; y < len(field); y++ {
		field[y] = make([]rune, maxX+1)
		for x := 0; x < len(field[0]); x++ {
			field[y][x] = '.'
		}
	}

	for _, dot := range dots {
		field[dot[1]][dot[0]] = '#'
	}

	return field, folds
}

func show(page [][]rune) {
	for r := 0; r < len(page); r++ {
		for c := 0; c < len(page[0]); c++ {
			fmt.Printf("%c ", page[r][c])
		}
		fmt.Println()
	}
	fmt.Println()
}

func fold(xs [][]rune, way string, pos int) [][]rune {
	var a, b [][]rune
	if way == "y" {
		a = xs[:pos]
		b = xs[pos+1:]
		for r := 0; r < len(a); r++ {
			for c := 0; c < len(a[0]); c++ {
				if a[r][c] == '#' {
					continue
				}
				a[r][c] = b[len(b)-1-r][c]
			}
		}
		return a

	} else if way == "x" {
		a = make([][]rune, len(xs))
		b = make([][]rune, len(xs))
		for r := 0; r < len(a); r++ {
			a[r] = make([]rune, pos)
			b[r] = make([]rune, pos)
			for c := 0; c < len(xs[0]); c++ {
				if c < pos {
					a[r][c] = xs[r][c]
				}
				if c > pos {
					b[r][c-pos-1] = xs[r][c]
				}
			}
		}
		for r := 0; r < len(a); r++ {
			for c := 0; c < len(a[0]); c++ {
				if a[r][c] == '#' {
					continue
				}
				a[r][c] = b[r][len(b[0])-1-c]
			}
		}
		return a
	} else {
		panic(fmt.Sprintf("wrong way: %s", way))
	}
}

func transparentOrigamiP1(xs []string, n int) (count int) {
	page, folds := parse(xs)
	for step := 0; step < n; step++ {
		f := folds[step]
		for k, v := range f {
			page = fold(page, k, v)
		}
	}

	for r := 0; r < len(page); r++ {
		for c := 0; c < len(page[0]); c++ {
			if page[r][c] == '#' {
				count++
			}
		}
	}
	return count
}

func transparentOrigamiP2(xs []string) {
	page, folds := parse(xs)
	for _, f := range folds {
		for k, v := range f {
			page = fold(page, k, v)
		}
	}
	show(page)
}
