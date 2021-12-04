/*
	https://adventofcode.com/2021/day/4
*/

//nolint:unused,revive // it's ok
package days

type Board struct {
	marked map[int]bool
	xs     [][]int
}

func (b Board) isWon() bool {
	rowLen := len(b.xs[0])
	colLen := len(b.xs[0])

	cols := make([][]int, len(b.xs))
	for i := range cols {
		cols[i] = make([]int, len(b.xs[0]))
	}

	for r, row := range b.xs {
		marked := 0
		for c, n := range row {
			cols[c][r] = n
			if val, ok := b.marked[n]; ok && val {
				marked++
			}
		}
		if marked == rowLen {
			return true
		}
	}

	for _, col := range cols {
		marked := 0
		for _, n := range col {
			if val, ok := b.marked[n]; ok && val {
				marked++
			}
		}
		if marked == colLen {
			return true
		}
	}

	return false
}

func (b Board) getUnmarkedNumsSum() int {
	sum := 0
	for i := 0; i < len(b.xs); i++ {
		for j := 0; j < len(b.xs[0]); j++ {
			if val, ok := b.marked[b.xs[i][j]]; !ok || !val {
				sum += b.xs[i][j]
			}
		}
	}
	return sum
}

func (b *Board) mark(val bool, ns ...int) {
	for _, n := range ns {
		b.marked[n] = val
	}
}

func giantSquidP1(nums []int, boards [][][]int) int {

	bs := make([]Board, len(boards))
	for i, board := range boards {
		bs[i] = Board{
			marked: map[int]bool{
				nums[0]: true,
				nums[1]: true,
				nums[2]: true,
				nums[3]: true,
				nums[4]: true,
			},
			xs: board,
		}
	}
	numsWin := len(nums) + 1
	boardWin := -1

	for bi, b := range bs {
		ni := 4
		for !b.isWon() && ni < len(nums)-1 {
			ni++
			b.mark(true, nums[ni])
		}
		if ni < numsWin {
			numsWin = ni
			boardWin = bi
		}
	}

	return bs[boardWin].getUnmarkedNumsSum() * nums[numsWin]
}

func giantSquidP2(nums []int, boards [][][]int) int {
	bs := make([]Board, len(boards))
	for i, board := range boards {
		bs[i] = Board{
			marked: make(map[int]bool, len(nums)),
			xs:     board,
		}
		bs[i].mark(true, nums...)
	}
	numsWin := 0
	boardWin := -1

	for bi, b := range bs {
		ni := len(nums) - 1
		for b.isWon() && ni > 0 {
			b.mark(false, nums[ni])
			ni--
		}
		ni++
		b.mark(true, nums[ni])
		if ni > numsWin {
			numsWin = ni
			boardWin = bi
		}
	}

	return bs[boardWin].getUnmarkedNumsSum() * nums[numsWin]
}
