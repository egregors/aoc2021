/*
	https://adventofcode.com/2021/day/1
*/

package days

func sonarSweepP1(xs []int) int {
	count := 1
	for i := 1; i < len(xs)-1; i++ {
		if xs[i+1] > xs[i] {
			count++
		}
	}
	return count
}

func sonarSweepP2(xs []int) int {
	count := 0
	sum := xs[0] + xs[1] + xs[2]
	for i := 1; i < len(xs)-2; i++ {
		if sum < xs[i]+xs[i+1]+xs[i+2] {
			count++
		}
		sum = xs[i] + xs[i+1] + xs[i+2]
	}
	return count
}
