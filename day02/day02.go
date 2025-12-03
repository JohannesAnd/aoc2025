package day02

import (
	"aoc2025/utils"
	"strconv"
	"strings"
)

type IdRange struct {
	from int
	to   int
}

func getData() []IdRange {
	lr, _ := utils.NewLineReader("day02.txt")

	dataString, _ := lr.Next()

	segments := strings.Split(dataString, ",")

	data := make([]IdRange, len(segments))

	for i, s := range segments {
		numbers := strings.Split(s, "-")

		from, _ := strconv.Atoi(numbers[0])
		to, _ := strconv.Atoi(numbers[1])

		data[i] = IdRange{from, to}
	}

	return data
}

func isInvalidPart1(number int) bool {
	str := strconv.Itoa(number)

	if utils.IsOdd(len(str)) {
		return false
	}

	return str[0:len(str)/2] == str[len(str)/2:]
}

func Part1() int {
	data := getData()

	sum := 0

	for _, d := range data {
		for i := d.from; i <= d.to; i++ {
			if isInvalidPart1(i) {
				sum += i
			}
		}
	}

	return sum
}

func checkIsRepeating(pattern string, value string) bool {
	patternSize := len(pattern)
	valueSize := len(value)

	if valueSize%patternSize != 0 {
		return false
	}

	for t := range valueSize / patternSize {
		if value[t*patternSize:(t+1)*patternSize] != pattern {
			return false
		}
	}

	return true
}

func isInvalidPart2(number int) bool {
	str := strconv.Itoa(number)

	for size := 1; size < len(str); size++ {
		if checkIsRepeating(str[0:size], str) {
			return true
		}
	}

	return false
}

func Part2() int {
	data := getData()

	sum := 0

	for _, d := range data {
		for i := d.from; i <= d.to; i++ {
			if isInvalidPart2(i) {
				sum += i
			}
		}
	}

	return sum
}
