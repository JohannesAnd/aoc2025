package day01

import (
	"aoc2025/utils"
	"strconv"
)

func Part1() int {
	lr, _ := utils.NewLineReader("day01.txt")

	position := 50
	count := 0

	for {
		line, ok := lr.Next()

		if !ok {
			break
		}

		distance, _ := strconv.Atoi(line[1:])

		if line[0] == 'L' {
			distance *= -1
		}

		position = utils.Mod(position+distance, 100)

		if position == 0 {
			count++
		}
	}

	return count
}

func Part2() int {
	lr, _ := utils.NewLineReader("day01.txt")

	position := 50
	count := 0

	for {
		line, ok := lr.Next()

		if !ok {
			break
		}

		distance, _ := strconv.Atoi(line[1:])

		for range distance {
			if line[0] == 'L' {
				if position == 0 {
					position = 99
				} else {
					position -= 1
				}
			} else {
				if position == 99 {
					position = 0
				} else {
					position += 1
				}
			}

			if position == 0 {
				count++
			}
		}
	}

	return count
}
