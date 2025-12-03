package day03

import (
	"aoc2025/utils"
	"math"
)

func findLargest(str string, start int, end int) (uint8, int) {
	var largest uint8 = 0
	index := -1

	for i := start; i < end; i++ {
		// Account for ASCII code offset
		value := str[i] - 48

		if value > largest {
			largest = value
			index = i
		}

		if value == 9 {
			break
		}
	}

	return largest, index
}

func findLargestNumberOfSize(str string, size int) int {
	result := 0

	index := -1

	for i := range size {
		posFromEnd := size - i - 1

		value, newIndex := findLargest(str, index+1, len(str)-posFromEnd)

		index = newIndex

		result += int(value) * int(math.Pow10(posFromEnd))
	}

	return result
}

func Part1() int {
	lr, _ := utils.NewLineReader("day03.txt")

	result := 0

	for {
		line, ok := lr.Next()

		if !ok {
			break
		}

		result += findLargestNumberOfSize(line, 2)
	}

	return result
}

func Part2() int {
	lr, _ := utils.NewLineReader("day03.txt")

	result := 0

	for {
		line, ok := lr.Next()

		if !ok {
			break
		}

		result += findLargestNumberOfSize(line, 12)
	}

	return result
}
