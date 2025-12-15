package day04

import (
	"os"
)

const newline byte = 10
const dot byte = 46
const at byte = 64

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

func countNeighbors(data []byte, width int, y int, x int) int {
	count := 0

	for _, y2 := range []int{y - 1, y, y + 1} {
		for _, x2 := range []int{x - 1, x, x + 1} {
			if x2 == x && y2 == y {
				continue
			}

			value, found := getValue(data, width, y2, x2)

			if found && value == at {
				count++
			}
		}
	}

	return count
}

func getAccessible(data []byte) []int {
	width, _ := findFirstIndex(data, newline)

	result := make([]int, 0)

	for i, d := range data {
		if d != at {
			continue
		}

		y := i / (width + 1)
		x := i % (width + 1)

		n := countNeighbors(data, width, y, x)

		if n < 4 {
			result = append(result, i)
		}
	}

	return result
}

func Part1(inputPath string) int {
	data, _ := os.ReadFile(inputPath)

	return len(getAccessible(data))
}

func Part2(inputPath string) int {
	data, _ := os.ReadFile(inputPath)

	removed := 0

	for {
		accessible := getAccessible(data)

		if len(accessible) == 0 {
			break
		}

		removed += len(accessible)

		for _, i := range accessible {
			data[i] = dot
		}
	}

	return removed
}
