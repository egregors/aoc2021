/*
	https://adventofcode.com/2021/day/11
*/

package days

func flash(r, c int, xs [][]int) int {
	count := 1
	xs[r][c] = 0

	// ! ! .
	// ! x .
	// . . .
	if r > 0 {
		if xs[r-1][c] != 0 {
			xs[r-1][c]++
			if xs[r-1][c] > 9 {
				count += flash(r-1, c, xs)
			}
		}
	}
	if c > 0 {
		if xs[r][c-1] != 0 {
			xs[r][c-1]++
			if xs[r][c-1] > 9 {
				count += flash(r, c-1, xs)
			}
		}
	}
	if r > 0 && c > 0 {
		if xs[r-1][c-1] != 0 {
			xs[r-1][c-1]++
			if xs[r-1][c-1] > 9 {
				count += flash(r-1, c-1, xs)
			}
		}
	}

	// . . .
	// . x !
	// . ! !
	if r < len(xs)-1 {
		if xs[r+1][c] != 0 {
			xs[r+1][c]++
			if xs[r+1][c] > 9 {
				count += flash(r+1, c, xs)
			}
		}
	}
	if c < len(xs[0])-1 {
		if xs[r][c+1] != 0 {
			xs[r][c+1]++
			if xs[r][c+1] > 9 {
				count += flash(r, c+1, xs)
			}
		}

	}
	if r < len(xs)-1 && c < len(xs[0])-1 {
		if xs[r+1][c+1] != 0 {
			xs[r+1][c+1]++
			if xs[r+1][c+1] > 9 {
				count += flash(r+1, c+1, xs)
			}
		}

	}

	// . . .
	// . x .
	// ! . .
	if r < len(xs)-1 && c > 0 {
		if xs[r+1][c-1] != 0 {
			xs[r+1][c-1]++
			if xs[r+1][c-1] > 9 {
				count += flash(r+1, c-1, xs)
			}
		}
	}

	// . . !
	// . x .
	// . . .
	if r > 0 && c < len(xs[0])-1 {
		if xs[r-1][c+1] != 0 {
			xs[r-1][c+1]++
			if xs[r-1][c+1] > 9 {
				count += flash(r-1, c+1, xs)
			}
		}
	}

	return count
}

func dumboOctopusP1(xs [][]int, steps int) int {
	flashes := 0
	for i := 0; i < steps; i++ {
		for ri := 0; ri < len(xs); ri++ {
			for ci := 0; ci < len(xs[0]); ci++ {
				xs[ri][ci]++
			}
		}

		for ri := 0; ri < len(xs); ri++ {
			for ci := 0; ci < len(xs[0]); ci++ {
				if xs[ri][ci] > 9 {
					flashes += flash(ri, ci, xs)
				}
			}
		}

	}
	return flashes
}

func dumboOctopusP2(xs [][]int) int {
	flashes, step := 0, 0
	for flashes != 100 {
		flashes = 0
		for ri := 0; ri < len(xs); ri++ {
			for ci := 0; ci < len(xs[0]); ci++ {
				xs[ri][ci]++
			}
		}
		for ri := 0; ri < len(xs); ri++ {
			for ci := 0; ci < len(xs[0]); ci++ {
				if xs[ri][ci] > 9 {
					flashes += flash(ri, ci, xs)
				}
			}
		}
		step++
	}
	return step
}
