/*
	https://adventofcode.com/2021/day/15
*/

package days

import (
	"log"

	"github.com/RyanCarrier/dijkstra"
)

func getIdx(i, j, w int) int {
	return i*w + j
}

func findPath(xs [][]int) int {
	graph := dijkstra.NewGraph()
	w := len(xs[0])

	for i, r := range xs {
		for j := range r {
			curIdx := getIdx(i, j, w)
			graph.AddVertex(curIdx)
		}
	}

	for i, r := range xs {
		for j := range r {
			curIdx := getIdx(i, j, w)
			// ^
			if i > 0 {
				_ = graph.AddArc(getIdx(i-1, j, w), curIdx, int64(xs[i][j]))
			}
			// v
			if i < len(xs)-1 {
				_ = graph.AddArc(getIdx(i+1, j, w), curIdx, int64(xs[i][j]))
			}
			// <
			if j > 0 {
				_ = graph.AddArc(getIdx(i, j-1, w), curIdx, int64(xs[i][j]))
			}
			// >
			if j < w-1 {
				_ = graph.AddArc(getIdx(i, j+1, w), curIdx, int64(xs[i][j]))
			}
		}
	}

	best, err := graph.Shortest(0, getIdx(len(xs)-1, w-1, w))
	if err != nil {
		log.Fatal(err)
	}
	return int(best.Distance)
}

func conv(row []int, n int) []int {
	newRow := make([]int, len(row))
	for i, num := range row {
		newNum := num + n
		if newNum > 9 {
			newNum -= 9
		}
		newRow[i] = newNum
	}
	return newRow
}

func increaseCave(xs [][]int) [][]int {
	res := make([][]int, len(xs))

	for ir, r := range xs {
		res[ir] = make([]int, len(xs[0]))
		for ic, c := range r {
			res[ir][ic] = c
		}
	}

	for i := 1; i < 5; i++ {
		for ri, r := range xs {
			res[ri] = append(res[ri], conv(r, i)...)
		}
	}

	for i := 1; i < 5; i++ {
		for j := 0; j < len(xs); j++ {
			res = append(res, conv(res[j], i))
		}
	}

	return res
}

func chitonP1(xs [][]int) int {
	return findPath(xs)
}

func chitonP2(xs [][]int) int {
	xs = increaseCave(xs)
	return findPath(xs)
}
