/*
	https://adventofcode.com/2021/day/17
*/

package days

import (
	"errors"
	"strconv"
	"strings"

	"github.com/egregors/aoc2021/utils"
)

// nolint:gocritic //meh
func parseShot(s string) (int, int, int, int) {
	// target area: x=20..30, y=-10..-5
	//              ^        |
	ps := strings.Split(strings.TrimPrefix(s, "target area: "), ", ")
	rawX := strings.Split(strings.TrimPrefix(ps[0], "x="), "..")
	rawY := strings.Split(strings.TrimPrefix(ps[1], "y="), "..")
	x1, _ := strconv.Atoi(rawX[0])
	x2, _ := strconv.Atoi(rawX[1])
	y1, _ := strconv.Atoi(rawY[0])
	y2, _ := strconv.Atoi(rawY[1])
	xMin := utils.Min(x1, x2)
	xMax := utils.Max(x1, x2)
	yMin := utils.Min(y1, y2)
	yMax := utils.Max(y1, y2)

	return xMin, xMax, yMin, yMax
}

func trickShotP1(s string) int {
	xMin, xMax, yMin, yMax := parseShot(s)
	h := 0
	for vx := 0; vx < 300; vx++ {
		for vy := yMin; vy < 300; vy++ {
			n, err := step(xMin, xMax, yMin, yMax, vx, vy)
			if err == nil {
				if n > h {
					h = n
				}
			}
		}
	}

	return h
}

func trickShotP2(s string) int {
	xMin, xMax, yMin, yMax := parseShot(s)
	c := 0
	for vx := 0; vx < 300; vx++ {
		for vy := yMin; vy < 300; vy++ {
			_, err := step(xMin, xMax, yMin, yMax, vx, vy)
			if err == nil {
				c++
			}
		}
	}

	return c
}

func step(xMin, xMax, yMin, yMax, vx, vy int) (int, error) {
	var x, y, h int
	for {
		x += vx
		y += vy

		if y > h {
			h = y
		}

		if vx < 0 {
			vx++
		} else if vx > 0 {
			vx--
		}

		vy--

		if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
			return h, nil
		}

		if y < yMin || (vx == 0 && (x > xMax || x < xMin)) {
			return -1, errors.New("miss")
		}
	}
}
