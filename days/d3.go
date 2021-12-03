package days

import (
	"strconv"
)

//nolint:revive //it's ok
const (
	O = uint8('0')
	I = uint8('1')
)

func binaryDiagnosticP1(xs []string) int {
	var gamma, epsilon string

	for i := 0; i < len(xs[0]); i++ {
		m := make(map[uint8]int, 2)
		for _, x := range xs {
			m[x[i]]++
		}

		if m[O] > m[I] {
			gamma += string(O)
			epsilon += string(I)
		} else {
			gamma += string(I)
			epsilon += string(O)
		}
	}

	gamma10, _ := strconv.ParseInt(gamma, 2, 64)
	epsilon10, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(gamma10 * epsilon10)
}

func getRating(xs []string, bit int, pred func(int, int) bool) int {
	if len(xs) == 1 {
		res, _ := strconv.ParseInt(xs[0], 2, 64)
		return int(res)
	}

	m := make(map[uint8]int)
	for _, x := range xs {
		m[x[bit]]++
	}

	var newXs []string
	var goodBit uint8

	if pred(m[I], m[O]) {
		goodBit = I
	} else {
		goodBit = O
	}

	for _, x := range xs {
		if x[bit] == goodBit {
			newXs = append(newXs, x)
		}
	}

	return getRating(newXs, bit+1, pred)
}

func binaryDiagnosticP2(xs []string) int {
	return getRating(
		xs,
		0,
		func(a, b int) bool {
			return a >= b
		}) * getRating(
		xs,
		0,
		func(a, b int) bool {
			return a < b
		})
}
