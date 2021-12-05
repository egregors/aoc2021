package days

import (
	"testing"

	"github.com/egregors/aoc2021/utils"
	"github.com/stretchr/testify/assert"
)

func Test_hydrothermalVentureP1(t *testing.T) {
	// tests
	xs := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}
	assert.Equal(t, 5, hydrothermalVentureP1(xs, 10))
	assert.Equal(t, 12, hydrothermalVentureP2(xs, 10))

	// solve
	xs, _ = utils.ReadLines("./d5_input.txt")
	assert.Equal(t, 5145, hydrothermalVentureP1(xs, 1000))
	assert.Equal(t, 16518, hydrothermalVentureP2(xs, 1000))

}

func Test_isDiagonal(t *testing.T) {
	type args struct {
		l Line
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 1",
			args{l: Line{1, 1, 3, 3}},
			true,
		},
		{
			"Test 2",
			args{l: Line{9, 7, 7, 9}},
			true,
		},
		{
			"Test 3",
			args{l: Line{8, 0, 0, 8}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDiagonal(tt.args.l); got != tt.want {
				t.Errorf("isDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}
