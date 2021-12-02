package days

import (
	"strconv"
	"strings"
)

func diveP1(xs []string) int {
	horPos, depth := 0, 0
	for _, s := range xs {
		cmd := strings.Split(s, " ")
		dir := cmd[0]
		val, _ := strconv.Atoi(cmd[1])
		switch dir {
		case "forward":
			horPos += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}
	return horPos * depth
}

func diveP2(xs []string) int {
	horPos, depth, aim := 0, 0, 0
	for _, s := range xs {
		cmd := strings.Split(s, " ")
		dir := cmd[0]
		val, _ := strconv.Atoi(cmd[1])
		switch dir {
		case "forward":
			horPos += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	return horPos * depth
}
