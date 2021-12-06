/*
	https://adventofcode.com/2021/day/6
*/

package days

func lanternfishP1(xs []int, days int) int {
	day := 1
	for day <= days {
		var nextGen []int

		for i := 0; i < len(xs); i++ {
			xs[i]--
			if xs[i] == -1 {
				xs[i] = 6
				nextGen = append(nextGen, 8)
			}
		}
		xs = append(xs, nextGen...)
		day++
	}

	return len(xs)
}

func nextDay(flock map[int]int, days int) map[int]int {
	if days == 0 {
		return flock
	}
	nextFlock := map[int]int{8: flock[0]}
	for i := 1; i < 9; i++ {
		nextFlock[i-1] = flock[i]
	}
	nextFlock[6] += flock[0]
	return nextDay(nextFlock, days-1)
}

func lanternfishP2(xs []int, days int) int {
	initFlock := make(map[int]int, 9)
	for _, x := range xs {
		initFlock[x]++
	}
	lastDayFlock := nextDay(initFlock, days)
	sum := 0
	for _, v := range lastDayFlock {
		sum += v
	}
	return sum
}
