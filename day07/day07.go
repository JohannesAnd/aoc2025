package day07

import (
	"os"
)

const newline = '\n'
const start = 'S'

func findFirstIndex(data []byte, value byte) (int, bool) {
	for i := 0; i < len(data); i++ {
		if data[i] == value {
			return i, true
		}
	}

	return -1, false
}

func getValue(data []byte, width int, y int, x int) (byte, bool) {
	index := y*(width+1) + x

	if y >= 0 && x >= 0 && x < width && index < len(data) {
		return data[index], true
	}

	return 0, false
}

func setValue(data []byte, width int, y int, x int, val byte) bool {
	index := y*(width+1) + x

	if y >= 0 && x >= 0 && x < width && index < len(data) {
		data[index] = val

		return true
	}

	return false
}

func spawn(data []byte, width int, y int, x int) int {
	val, found := getValue(data, width, y, x)

	if !found {
		return 0
	}

	if val == '.' {
		setValue(data, width, y, x, '|')

		return spawn(data, width, y+1, x)
	}

	if val == '|' {
		return 0
	}

	if val == '^' {
		return 1 + spawn(data, width, y, x+1) + spawn(data, width, y, x-1)
	}

	panic("Unknown value")
}

var memo = make(map[int]int)

func spawn2(data []byte, width int, y int, x int) int {
	key := y*width + x
	value := memo[key]

	if value != 0 {
		return value
	}

	val, found := getValue(data, width, y, x)

	res := 0

	if !found {
		res = 1
	} else if val == '.' {
		res = spawn2(data, width, y+1, x)
	} else if val == '^' {
		res = spawn2(data, width, y, x+1) + spawn2(data, width, y, x-1)
	}

	memo[key] = res

	return res
}

func Part1() int {
	data, _ := os.ReadFile("day07.txt")

	width, _ := findFirstIndex(data, newline)
	x, _ := findFirstIndex(data, start)
	y := 1

	return spawn(data, width, y, x)
}

func Part2() int {

	data, _ := os.ReadFile("day07.txt")

	width, _ := findFirstIndex(data, newline)
	x, _ := findFirstIndex(data, start)
	y := 1

	return spawn2(data, width, y, x)
}
